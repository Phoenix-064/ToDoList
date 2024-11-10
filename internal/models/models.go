package models

// User 用户信息
type User struct {
	Email    string `json:"email" gorm:"column:email;not null;type:varchar(30)"`
	Uuid     string `json:"uuid" gorm:"column:uuid;primaryKey;type:varchar(36)"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(30)"`
	IsAdmin  bool   `json:"is_admin" gorm:"column:is_admin;not null;default:false"`
	Todos    []Todo `json:"todos" gorm:"foreignKey:UserUuid;references:Uuid"`
	Wishes   []Wish `json:"wishes" gorm:"foreignKey:UserUuid;references:Uuid"`
}

// Todo 待办
type Todo struct {
	ID              string `json:"id" gorm:"column:id;primaryKey"`
	UserUuid        string `json:"user_uuid" gorm:"column:user_uuid;index;type:varchar(36)"`
	Event           string `json:"event" gorm:"column:event;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	Completed       bool   `json:"completed" gorm:"column:completed"`
	IsCycle         bool   `json:"is_cycle" gorm:"column:is_cycle"`
	Description     string `json:"description" gorm:"column:description;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	ImportanceLevel int    `json:"importance_level" gorm:"column:importance_level"`
}

// Response 标准回应结构体
type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// Wish 希望不要加班了
type Wish struct {
	ID          string `json:"id" gorm:"column:id;primaryKey"`
	UserUuid    string `json:"user_uuid" gorm:"column:user_uuid;index;type:varchar(36)"`
	Event       string `json:"event" gorm:"column:event;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	IsCycle     bool   `json:"is_cycle" gorm:"column:is_cycle"`
	Description string `json:"description" gorm:"column:description;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	IsShared    bool   `json:"is_wish" gorm:"column:is_shared"`
}

// CommunityWish 社区心愿
type CommunityWish struct {
	Description string `json:"description"`
	Event       string `json:"event"`
	ID          string `json:"id"`
	Viewed      int64  `json:"viewed"`
}
