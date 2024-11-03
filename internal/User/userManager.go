package user

import "github.com/google/uuid"

// User 用户信息
type User struct {
	email    string
	uuid     string
	name     string
	password string
	isAdmin  bool
}

// UserList 用户列表
type UserList struct {
	list map[string]User //[uuid]user
}

// UserManager 用户管理
type UserManager struct {
}

// UserHandle 用户管理的方法
type UserHandle interface {
	AddUser(User) error
	DeleteUser(string) error
	CheckUser(User) (bool, string, error) //string值为uuid
	UpdateUser(User) error
}

// newUser 返回一个新的用户
func newUser() (User, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}

	return User{uuid: u.String(), isAdmin: false}, nil
}

// newUserList 返回一个新的用户列表
func newUserList() UserList {
	return UserList{}
}

// NewUserManager 返回一个新的用户管理系统
func NewUserManager() UserManager {
	return UserManager{}
}

// AddUser 添加用户
func (ul UserList) AddUser(u User) error {
	ul.list[u.uuid] = u
	return nil
}

// CheckUser 查找用户
func (ul UserList) CheckUser() {

}

// DeleteUser 删除用户
func (ul UserList) DeleteUser(string) error {

	return nil
}
