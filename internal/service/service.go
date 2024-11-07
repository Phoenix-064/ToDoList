package service

import (
	"ToDoList/internal/engine"

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
		ToDoList.GET("")
	}
}
