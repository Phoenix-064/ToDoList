package models

// User 用户信息
type User struct {
	Email    string `json:"email" gorm:"column:email;not null;type:varchar(30)"`
	Uuid     string `json:"uuid" gorm:"column:uuid;primaryKey"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(30)"`
	IsAdmin  bool   `json:"isAdmin" gorm:"column:is_admin;not null;default:false"`
	Todo     []Todo `json:"todos" gorm:"foreignKey:user_uuid;references:Uuid"`
}

// Todo 待办
type Todo struct {
	ID              string `json:"id" gorm:"column:id;primaryKey"`
	UserUuid        string `json:"user_uuid" gorm:"column:user_uuid"`
	Event           string `json:"event" gorm:"column:event"`
	Completed       bool   `json:"completed" gorm:"column:completed"`
	IsCycle         bool   `json:"is_cycle" gorm:"column:is_cycle"`
	Description     string `json:"description" gorm:"column:description"`
	IsWish          bool   `json:"is_wish" gorm:"column:is_wish"`
	ImportanceLevel int    `json:"importance_level" gorm:"column:importance_level"`
}

// Response 标准回应结构体
type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
