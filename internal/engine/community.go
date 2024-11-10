package engine

import (
	"ToDoList/internal/middleware"
	"ToDoList/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (eh *EngineHandler) GetCommunityWishes(ctx *gin.Context) {
	wishes, err := eh.CommunityWishManager.GetWishes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: "",
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: wishes,
	})
}

func (eh *EngineHandler) AddView(ctx *gin.Context) {
	id := ctx.Query("id")
	if err := eh.CommunityWishManager.AddView(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err,
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}

func (eh *EngineHandler) AddToWish(ctx *gin.Context) {
	id := ctx.Query("id")
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err,
		})
		logrus.Error(err)
		return
	}
	if err := eh.CommunityWishManager.AddToWish(uuid, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err,
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}
