package engine

import (
	data "ToDoList/internal/Data"
	"ToDoList/internal/middleware"
	"ToDoList/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetTodo struct {
	ID    string `json:"ID"`
	Event string `json:"event"`
}

// GetAllTodo 获取所有待办事项
func (eh *EngineHandler) GetAllTodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	todos, err := eh.TodoManager.ReadUserTodos(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("读取数据时出错，", err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: gin.H{
			"todos": todos,
		},
	})
}

// CreateTodo 创建待办事项
func (eh *EngineHandler) CreateTodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	var tempTodo GetTodo
	if err = ctx.ShouldBind(&tempTodo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	todo := data.NewTodo(tempTodo.ID, tempTodo.Event)
	if err = eh.TodoManager.AddTodo(uuid, *todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "添加成功",
	})
}

// SaveAllTodos 添加所有 todo
func (eh *EngineHandler) SaveAllTodos(ctx *gin.Context) {
	var todos []models.Todo
	if err := ctx.ShouldBind(&todos); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("绑定结构体失败，", err)
		return
	}
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err = eh.TodoManager.SaveTheUserTodos(uuid, todos); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "保存成功",
	})
}

// DeleteTodo 删除 todo
func (eh *EngineHandler) DeleteTodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	var tempTodo models.Todo
	if err = ctx.ShouldBind(&tempTodo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err = eh.TodoManager.DeleteTodo(uuid, tempTodo.ID); err != nil {
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

// GetATodo 获得一个随机 todo
func (eh *EngineHandler) GetATodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	todo, err := eh.TodoManager.RandomlySelectTodo(uuid)
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
		Content: todo,
	})
}
