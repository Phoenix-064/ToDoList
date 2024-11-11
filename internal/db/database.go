package db

import (
	"ToDoList/internal/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDataBase() (Database, error) {
	dsn := os.Getenv("DSN")
	// dsn := "root:123@tcp(127.0.0.1:3306)/todoList?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}, &models.Wish{}); err != nil {
		return Database{}, err
	}
	if err := db.AutoMigrate(&models.CommunityWish{}); err != nil {
		return Database{}, err
	}
	return Database{DB: db}, nil
}
