package utils

import (
	"ToDoList/internal/db"
	"ToDoList/internal/models"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VerificationCodeHandler struct {
	db *gorm.DB
}

type VerificationCodeManager interface {
	NewVerificationCode(email string) (models.VerificationCode, error)
	CheckTheSendingFrequency(email string) error
	CleanUpExpiredVerificationCodes() error
	CheckTheVerificationCode(email string, code string) error
}

// NewVerificationCodeHandler 初始化一个VerificationCode
func NewVerificationCodeHandler() VerificationCodeHandler {
	db, err := db.NewDataBase()
	if err != nil {
		logrus.Fatal(err)
	}
	return VerificationCodeHandler{db: db.DB}
}

// NewVerificationCode 初始化并储存一个VerificationCode
func (vch VerificationCodeHandler) NewVerificationCode(email string) (models.VerificationCode, error) {
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	verificationCode := models.VerificationCode{
		Email:      email,
		CreateTime: time.Now(),
		Code:       code,
		Used:       false,
	}
	result := vch.db.Create(&verificationCode)
	if result.Error != nil {
		return verificationCode, result.Error
	}
	return verificationCode, nil
}

// CheckTheSendingFrequency 检查验证码发送频率
func (vch VerificationCodeHandler) CheckTheSendingFrequency(email string) error {
	lastCode := models.VerificationCode{}
	result := vch.db.Where("email = ?", email).Order("create_time DESC").First(&lastCode)
	logrus.Println(time.Now().Sub(lastCode.CreateTime) < time.Minute)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		// 这里实际上可以改进，在查找是要求要是未使用的
		// 但我想让用户体验一下一次发出多条验证码的感觉
	}
	nextAllowedTime := lastCode.CreateTime.Add(time.Minute)
	if time.Now().Before(nextAllowedTime) {
		return errors.New("创建验证码时间间隔小于一分钟")
	}
	return nil
}

// CleanUpExpiredVerificationCodes 清理过期验证码，似乎没有用
func (vch VerificationCodeHandler) CleanUpExpiredVerificationCodes() error {
	ticker := time.NewTicker(6 * time.Hour)
	var code models.VerificationCode
	for range ticker.C { // 每6小时触发一次
		vch.db.Where("create_time < ? OR used = ?",
			time.Now().Add(6*time.Hour),
			true).Delete(&code)
	}
	return nil
}

// CheckTheVerificationCode 检查验证码是否正确与是否过期与是否使用
func (vch VerificationCodeHandler) CheckTheVerificationCode(email string, code string) error {
	var verificationCode models.VerificationCode
	result := vch.db.Where("email = ? AND code = ? AND create_time > ? AND used = ?",
		email, code, time.Now().Add(-5*time.Minute), false).First(&verificationCode)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("错误的验证码")
		}
		return result.Error
	}
	if err := vch.db.Model(&verificationCode).Where("code = ? AND email = ?", code, email).Update("used", true).Error; err != nil {
		return err
	}
	return nil
}
