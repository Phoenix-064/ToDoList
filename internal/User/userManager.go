package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	Email    string `json:"email" gorm:"column:email;not null;type:varchar(30)"`
	Uuid     string `json:"uuid" gorm:"column:uuid;primaryKey"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(30)"`
	IsAdmin  bool   `json:"isAdmin" gorm:"column:isAdmin;not null;default:false"`
}

// UserList 用户列表(使用json,测试时使用)
type UserList struct {
	list map[string]User //[uuid]user
}

// UserManager 基于数据库的用户管理
type UserManager struct {
	db *gorm.DB
}

// UserHandle 用户管理的方法
type UserHandle interface {
	AddUser(User) error
	DeleteUser(string) error
	CheckUser(string) (User, error) //可以输入uuid,email
	CheckUuid(uuid string) (User, error)
	CheckEmail(email string) (User, error)
	UpdateUser(former User, later User) error
}

// newUser 返回一个新的用户
func NewUser() (User, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}
	return User{Uuid: u.String(), IsAdmin: false}, nil
}

// newUserList 返回一个新的用户列表
func NewUserList() *UserList {
	return &UserList{}
}

// NewUserManager 返回一个新的用户管理系统
func NewUserManager() *UserManager {
	dsn := "root:123@tcp(127.0.0.1:3306)/todoList?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("无法连接数据库")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		logrus.Error("无法连接数据库")
	}
	return &UserManager{db: db}
}

// AddUser 添加用户
func (ul *UserList) AddUser(u User) error {
	ul.list[u.Uuid] = u
	return nil
}

// CheckUser 查找用户
func (ul *UserList) CheckUser(c string) (User, error) {
	u, ok := ul.list[c]
	if ok {
		return u, nil
	} else {
		for _, j := range ul.list {
			if j.Email == c {
				return j, nil
			}
		}
	}
	return u, errors.New("没有此用户")
}

// DeleteUser 删除用户
func (ul *UserList) DeleteUser(c string) error {
	u, err := ul.CheckUser(c)
	if err != nil {
		return err
	}
	delete(ul.list, u.Uuid)
	return nil
}

// UpdateUser 更新用户信息
func (ul *UserList) UpdateUser(former User, later User) error {
	_, err := ul.CheckUser(former.Uuid)
	if err != nil {
		return err
	}
	ul.list[former.Uuid] = later
	return nil
}

// AddUser 添加用户
func (um *UserManager) AddUser(u User) error {
	_, err := um.CheckUser(u.Email)
	if err == nil {
		return errors.New("已有的邮箱")
	}
	result := um.db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckUser 查找用户，支持查找uuid和email
func (um *UserManager) CheckUser(c string) (User, error) {
	var user User
	result := um.db.Where("email = ?  OR uuid = ?", c, c).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("没有此用户")
		}
		return user, result.Error
	}
	return user, nil
}

// CheckUuid 查找Uuid
func (um *UserManager) CheckUuid(uuid string) (User, error) {
	var user User
	result := um.db.Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("没有此用户")
		}
		return user, result.Error
	}
	return user, nil
}

// CheckEmail 查找email
func (um *UserManager) CheckEmail(email string) (User, error) {
	var user User
	result := um.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("没有此用户")
		}
		return user, result.Error
	}
	return user, nil
}

// DeleteUser 删除用户
func (um *UserManager) DeleteUser(c string) error {
	temp, err := um.CheckUser(c)
	// 如果要删除多个数据
	// 应该进行提醒
	//
	if err != nil {
		return err
	}
	um.db.Delete(&temp)
	return nil
}

// UpdateUser 更新用户数据
func (um *UserManager) UpdateUser(former User, later User) error {
	var u User
	result := um.db.Model(&u).Where("uuid = ?", former.Uuid).Updates(later)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
