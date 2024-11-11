package data

import (
	"ToDoList/internal/models"

	"gorm.io/gorm"
)

type HandleWish interface {
	ReadUserWishes(uuid string) ([]models.Wish, error)
	SaveTheUserWishes(uuid string, wishes []*models.Wish) error
	AddWishes(uuid string, wish *models.Wish) error
	DeleteWish(uuid string, wishID string) error
	RandomlySelectWish(uuid string) (models.Wish, error)
	UpdateWish(userUUID string, wishID string, wish *models.Wish) error
	AddWishToTodo(userUUID string, wishID string) error
}

type WishManager struct {
	db *gorm.DB
}

func NewWishManager(db *gorm.DB) *WishManager {
	return &WishManager{db: db}
}

func NewWish(id string, Event string, isCycle bool, description string, IsShared bool) *models.Wish {
	return &models.Wish{
		ID:          id,
		Event:       Event,
		IsCycle:     isCycle,
		Description: description,
		IsShared:    IsShared,
	}
}

// ReadUserWishes 读取用户所有的 wish
func (m *WishManager) ReadUserWishes(uuid string) ([]models.Wish, error) {
	var wishes []models.Wish
	if result := m.db.Find(&wishes); result.Error != nil {
		return []models.Wish{}, result.Error
	}
	return wishes, nil
}

// SaveTheUserWIshes 保存用户的 Wishes
func (m *WishManager) SaveTheUserWishes(uuid string, wishes []*models.Wish) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		if result := m.db.Where("user_uuid = ?", uuid).Delete(&models.Wish{}); result.Error != nil {
			return result.Error
		}
		if len(wishes) == 0 {
			return nil
		}
		for i := range wishes {
			wishes[i].UserUuid = uuid
		}
		if result := m.db.Create(wishes); result.Error != nil {
			return result.Error
		}
		return nil
	})
}

// AddWishes 添加一个 wish
func (m *WishManager) AddWishes(uuid string, wish *models.Wish) error {
	wish.UserUuid = uuid
	if wish.IsShared {
		h := NewCommunityWishesHandler(m.db)
		if err := h.AddWishToCommunity(wish); err != nil {
			return err
		}
	}
	if result := m.db.Create(wish); result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteWish 删除一个 wish
func (m *WishManager) DeleteWish(uuid string, wishID string) error {
	if result := m.db.Where("user_uuid = ? AND id = ?", uuid, wishID).Delete(&models.Wish{}); result.Error != nil {
		return result.Error
	}
	return nil
}

// RandomlySelectWish 获取一个随机的 wish
func (m *WishManager) RandomlySelectWish(uuid string) (models.Wish, error) {
	var wish models.Wish
	if result := m.db.Where("user_uuid = ?", uuid).Order("RAND()").First(&wish); result.Error != nil {
		return models.Wish{}, result.Error
	}
	return wish, nil
}

// UpdateWish 修改一个 wish
func (m *WishManager) UpdateWish(userUUID string, wishID string, wish *models.Wish) error {
	wish.UserUuid = userUUID
	var formerWish models.Wish
	if result := m.db.Where("user_uuid = ? AND id = ?", userUUID, wishID).First(&formerWish); result.Error != nil {
		return result.Error
	}
	if result := m.db.Where("user_uuid = ? AND id = ?", userUUID, wishID).Updates(*wish); result.Error != nil {
		return result.Error
	}
	if result := m.db.Select("is_shared").Where("user_uuid = ? AND id = ?", userUUID, wishID).Updates(*wish); result.Error != nil {
		return result.Error
	}
	// 检查是否需要添加至社区，或者是否需要从社区删除
	if formerWish.IsShared != wish.IsShared {
		h := NewCommunityWishesHandler(m.db)
		if wish.IsShared {
			if err := h.AddWishToCommunity(wish); err != nil {
				return err
			}
		} else {
			if err := h.DeleteWishFromCommunity(wish); err != nil {
				return err
			}
		}
	}
	return nil
}

// AddWishToTodo 将一个心愿添加至任务
func (m *WishManager) AddWishToTodo(userUUID string, wishID string) error {
	var wish models.Wish
	if result := m.db.Where("user_uuid = ? AND id = ?", userUUID, wishID).First(&wish); result.Error != nil {
		return result.Error
	}
	todo := models.Todo{
		ID:              wish.ID,
		Event:           wish.Event,
		ImportanceLevel: 0,
		UserUuid:        userUUID,
		Completed:       false,
		IsCycle:         wish.IsCycle,
		Description:     wish.Description,
		CompletedDate:   "",
	}
	if result := m.db.Create(&todo); result.Error != nil {
		return result.Error
	}
	if result := m.db.Where("user_uuid = ? AND id = ?", userUUID, wishID).Delete(&models.Wish{}); result.Error != nil {
		return result.Error
	}
	return nil
}
