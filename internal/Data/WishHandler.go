package data

import "ToDoList/internal/models"

type HandleWish interface {
	ReadUserWishes(uuid string) ([]models.Wish, error)
	SaveTheUserWishes(uuid string, wishes []models.Wish) error
	AddWishes(uuid string, wish models.Wish) error
	DeleteWish(uuid string, wishID string) error
	RandomlySelectWish(uuid string) (models.Wish, error)
	UpdateWish(userUUID string, wishID string, wish models.Wish) error
}
