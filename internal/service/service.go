package service

import (
	"ToDoList/internal/engine"
	"ToDoList/internal/middleware"

	"github.com/gin-gonic/gin"
)

type GinService struct {
}

func NewGinService() GinService {
	return GinService{}
}

func (gs GinService) SetUpRoutes(e *gin.Engine, eh engine.EngineHandler) {
	ToDoList := e.Group("/todolist")
	//建立相关路由组
	{
		User := ToDoList.Group("/user")
		{
			User.POST("/signin", eh.SignIn)
			User.POST("/signup/send-code", eh.SendVerificationCode)
			User.POST("/signup", eh.SignUp)
		}
		ToDoList.GET("", middleware.AuthMiddleware(), eh.GetAllTodo)
		ToDoList.POST("/add", middleware.AuthMiddleware(), eh.CreateTodo)
		ToDoList.POST("/update", middleware.AuthMiddleware(), eh.SaveAllTodos)
		ToDoList.POST("delete", middleware.AuthMiddleware(), eh.DeleteTodo)
		ToDoList.GET("/random", middleware.AuthMiddleware(), eh.GetATodo)
	}
}