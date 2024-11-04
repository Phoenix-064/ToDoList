package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	Email    string `json:"email" gorm:"column:email;not null;type:varchar(30)"`
	Uuid     string `json:"uuid" gorm:"column:uuid;primaryKey"`
	Name     string `json:"name" gorm:"column:name;not null;type:varchar(30)"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(30)"`
	IsAdmin  bool   `json:"isAdmin" gorm:"column:isAdmin;not null;default:false"`
}

// UserList 用户列表(测试时使用)
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
	CheckUser(string) ([]User, error) //可以输入,email,name
	UpdateUser(former User, later User) error
}

// newUser 返回一个新的用户
func newUser() (User, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}
	return User{Uuid: u.String(), IsAdmin: false}, nil
}

// newUserList 返回一个新的用户列表
func NewUserList() UserList {
	return UserList{}
}

// NewUserManager 返回一个新的用户管理系统
func NewUserManager(dsn string) UserManager {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {

	}
	return UserManager{db: db}
}

// AddUser 添加用户
func (ul *UserList) AddUser(u User) error {
	ul.list[u.Uuid] = u
	return nil
}

// CheckUser 查找用户
func (ul *UserList) CheckUser(c string) ([]User, error) {
	u, ok := ul.list[c]
	var temp []User
	if ok {
		temp = append(temp, u)
		return temp, nil
	} else {
		for _, j := range ul.list {
			if j.Name == c && j.Email == c {
				temp = append(temp, j)
			}
		}
	}
	if len(temp) != 0 {
		return temp, nil
	}
	return temp, errors.New("没有此用户")
}

// DeleteUser 删除用户
func (ul *UserList) DeleteUser(c string) error {
	u, err := ul.CheckUser(c)
	if err != nil {
		return err
	}
	//这里未来可以加一个警告功能，对于删除了多个用户时
	// if len(u) > 1{

	// }
	for _, j := range u {
		delete(ul.list, j.Uuid)
	}
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
	result := um.db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckUser 查找用户
func (um *UserManager) CheckUser(c string) ([]User, error) {
	var temp []User
	err := um.db.Where("name = ?", c).Find(&temp).Error
	if err != nil {
		return temp, err
	}
	err = um.db.Where("uuid = ?", c).Find(&temp).Error
	if err != nil {
		return temp, err
	}
	err = um.db.Where("email = ?", c).Find(&temp).Error
	if err != nil {
		return temp, err
	}
	if len(temp) != 0 {
		return temp, nil
	}
	return temp, errors.New("没有此用户")
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
