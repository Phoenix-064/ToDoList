package user

import "github.com/google/uuid"

type User struct {
	uuid     string
	name     string
	password string
	isAdmin  bool
}

type UserList struct {
	list map[string]User //[uuid]user
}

type UserManager struct {
}

type UserHandle interface {
	addUser(User) error
	deleteUser(string) error
	checkUser(User) (bool, string, error) //string值为uuid
	updateUser(User) error
}

func newUser() (User, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}

	return User{uuid: u.String()}, nil
}

func newUserList() UserList {
	return UserList{}
}

func newUserManager() UserManager {
	return UserManager{}
}

func (ul UserList) addUser(u User) error {
	ul.list[u.uuid] = u
	return nil
}
