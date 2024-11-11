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
	Description     string `json:"description"`
	Event           string `json:"event"`
	ID              string `json:"id"`
	ImportanceLevel int64  `json:"importance_level"`
	IsCycle         bool   `json:"is_cycle"`
}

type TodosWrapper struct {
	Todos []*models.Todo `json:"todos"`
}

type RequestChangeComplete struct {
	Completed string `json:"completed"`
	ID        string `json:"id"`
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
	todo := data.NewTodo(tempTodo.ID, tempTodo.Event, tempTodo.IsCycle, tempTodo.Description, int(tempTodo.ImportanceLevel))
	if err = eh.TodoManager.AddTodo(uuid, todo); err != nil {
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
	TodosWrapper := TodosWrapper{}
	if err := ctx.ShouldBind(&TodosWrapper); err != nil {
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
	if err = eh.TodoManager.SaveTheUserTodos(uuid, TodosWrapper.Todos); err != nil {
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
	todoID := ctx.Query("id")
	if err = eh.TodoManager.DeleteTodo(uuid, todoID); err != nil {
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

// // GetATodo 获得一个随机 todo
// func (eh *EngineHandler) GetATodo(ctx *gin.Context) {
// 	uuid, _, err := middleware.GetHeader(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, models.Response{
// 			Message: "err",
// 			Content: err.Error(),
// 		})
// 		logrus.Error(err)
// 		return
// 	}
// 	todo, err := eh.TodoManager.RandomlySelectTodo(uuid)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, models.Response{
// 			Message: "err",
// 			Content: err.Error(),
// 		})
// 		logrus.Error(err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, models.Response{
// 		Message: "ok",
// 		Content: todo,
// 	})
// }

// 更新一个 todo
func (eh *EngineHandler) UpdateTodo(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	todo := &models.Todo{}
	if err := ctx.ShouldBind(todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err = eh.TodoManager.UpdateTodo(uuid, todo.ID, todo); err != nil {
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

// RecordCompletionTime 记录完成时间
func (eh *EngineHandler) RecordCompletionTime(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}

	RequestTodo := RequestChangeComplete{}
	if err := ctx.ShouldBind(&RequestTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	todo := &models.Todo{
		CompletedDate: RequestTodo.Completed,
	}
	if err := eh.TodoManager.UpdateTodo(uuid, RequestTodo.ID, todo); err != nil {
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
