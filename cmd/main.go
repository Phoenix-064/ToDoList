package main

import (
	"ToDoList/internal/engine"
	"ToDoList/internal/logs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(e *gin.Engine, eh engine.EngineHandler) {
	ToDoList := e.Group("/todolist")
	//建立相关路由组
	{
		User := ToDoList.Group("/user")
		{
			User.POST("/signin", eh.SignIn)
			User.POST("/signup", eh.SignUp)
		}
		ToDoList.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
}

func main() {
	var setLog logs.AddLog
	setLog = logs.NewDefaultLog()
	setLog.SetUpLogs()
	eh := engine.NewEngineHandler()
	e := gin.New()
	SetUpRoutes(e, eh)
	if err := e.Run(":8080"); err != nil {
		logrus.Fatal("建立路由器失败")
	}
}
