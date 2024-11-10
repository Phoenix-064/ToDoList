package data

import (
	"ToDoList/internal/models"

	"gorm.io/gorm"
)

type HandleCommunityWishes interface {
	GetWishes() (*[]models.CommunityWish, error)
	AddView(id string) error
	AddToWish(uuid string, id string) error
}

type CommunityWishesHandler struct {
	db *gorm.DB
}

func NewCommunityWishesHandler(db *gorm.DB) *CommunityWishesHandler {
	return &CommunityWishesHandler{db: db}
}

// GetWishes 返回 100 条 Wish
func (h *CommunityWishesHandler) GetWishes() (*[]models.CommunityWish, error) {
	wishes := []models.CommunityWish{}
	if result := h.db.Find(&wishes).Limit(100); result.Error != nil {
		return &[]models.CommunityWish{}, result.Error
	}
	return &wishes, nil
}

// AddView 添加浏览量
func (h *CommunityWishesHandler) AddView(id string) error {
	wish := &models.CommunityWish{}
	if result := h.db.Where("id = ?", id).First(wish); result.Error != nil {
		return result.Error
	}
	wish.Viewed += 1
	return nil
}

// AddToWish 添加至心愿
func (h *CommunityWishesHandler) AddToWish(uuid string, id string) error {
	tempWish := models.CommunityWish{}
	if result := h.db.Where("id = ?", id).First(&tempWish); result.Error != nil {
		return result.Error
	}
	wish := NewWish(id, tempWish.Event, false, tempWish.Description, false)
	WishManager := NewWishManager(h.db)
	if err := WishManager.AddWishes(uuid, wish); err != nil {
		return err
	}
	return nil
}
