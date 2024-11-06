package main

import (
	"ToDoList/internal/engine"
	"ToDoList/internal/logs"
	"ToDoList/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	var setLog logs.AddLog
	setLog = logs.NewDefaultLog()
	setLog.SetUpLogs()
	eh := engine.NewEngineHandler()
	e := gin.New()
	var sh service.GinService
	sh.SetUpRoutes(e, eh)
	// 这里应该使用go route	启动验证码清理程序
	if err := e.Run(":8080"); err != nil {
		logrus.Fatal("建立路由器失败")
	}
}
