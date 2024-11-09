package user

import (
	"ToDoList/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserList 用户列表(使用json,测试时使用)
type UserList struct {
	list map[string]models.User //[uuid]models.user
}

// UserManager 基于数据库的用户管理
type UserManager struct {
	db *gorm.DB
}

// HandleUser 用户管理的方法
type HandleUser interface {
	AddUser(models.User) error
	DeleteUser(string) error
	CheckUser(string) (models.User, error) //可以输入uuid,email
	CheckUuid(uuid string) (models.User, error)
	CheckEmail(email string) (models.User, error)
	UpdateUser(former models.User, later models.User) error
}

// newUser 返回一个新的用户
func NewUser() (models.User, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return models.User{}, err
	}
	return models.User{Uuid: u.String(), IsAdmin: false}, nil
}

// newUserList 返回一个新的用户列表
func NewUserList() *UserList {
	return &UserList{}
}

// NewUserManager 返回一个新的用户管理系统
func NewUserManager(db *gorm.DB) *UserManager {
	return &UserManager{db: db}
}

// AddUser 添加用户
func (ul *UserList) AddUser(u models.User) error {
	ul.list[u.Uuid] = u
	return nil
}

// CheckUser 查找用户
func (ul *UserList) CheckUser(c string) (models.User, error) {
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
func (ul *UserList) UpdateUser(former models.User, later models.User) error {
	_, err := ul.CheckUser(former.Uuid)
	if err != nil {
		return err
	}
	ul.list[former.Uuid] = later
	return nil
}

// AddUser 添加用户
func (um *UserManager) AddUser(u models.User) error {
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
func (um *UserManager) CheckUser(c string) (models.User, error) {
	var user models.User
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
func (um *UserManager) CheckUuid(uuid string) (models.User, error) {
	var user models.User
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
func (um *UserManager) CheckEmail(email string) (models.User, error) {
	var user models.User
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
	if err != nil {
		return err
	}
	um.db.Delete(&temp)
	return nil
}

// UpdateUser 更新用户数据
func (um *UserManager) UpdateUser(former models.User, later models.User) error {
	var u models.User
	result := um.db.Model(&u).Where("uuid = ?", former.Uuid).Updates(later)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
