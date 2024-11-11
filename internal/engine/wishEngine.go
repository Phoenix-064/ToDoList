package engine

import (
	data "ToDoList/internal/Data"
	"ToDoList/internal/middleware"
	"ToDoList/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WishRequest struct {
	Description string `json:"description"`
	Event       string `json:"event"`
	ID          string `json:"id"`
	IsCycle     bool   `json:"is_cycle"`
	IsShared    bool   `json:"is_shared"`
}

type AddToTodoRequest struct {
	ID string `json:"id"`
}

// RandomlySelectWish 获取随机的 wish
func (eh *EngineHandler) RandomlySelectWish(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	wish, err := eh.WishManager.RandomlySelectWish(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: wish,
	})
}

// GetWishes 获取用户所有 wish
func (eh *EngineHandler) GetWishes(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	wishes, err := eh.WishManager.ReadUserWishes(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: gin.H{
			"wishes": wishes,
		},
	})
}

// DeleteWish 删除一条 wish
func (eh *EngineHandler) DeleteWish(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	WishID := ctx.Query("id")
	if err = eh.WishManager.DeleteWish(uuid, WishID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
	return
}

// UpdateWish 修改一条 wish
func (eh *EngineHandler) UpdateWish(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	wish := &models.Wish{}
	if err = ctx.ShouldBind(wish); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err = eh.WishManager.UpdateWish(uuid, wish.ID, wish); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}

// AddWish 添加一条 wish
func (eh *EngineHandler) AddWish(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	wish := &WishRequest{}
	if err = ctx.ShouldBind(wish); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	tempWish := data.NewWish(wish.ID, wish.Event, wish.IsCycle, wish.Description, wish.IsShared)
	if err = eh.WishManager.AddWishes(uuid, tempWish); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}

// AddToTodo 将一个 wish 添加至待办
func (eh *EngineHandler) AddToTodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	WishID := &AddToTodoRequest{}
	if err = ctx.ShouldBind(&WishID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err = eh.WishManager.AddWishToTodo(uuid, WishID.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}
