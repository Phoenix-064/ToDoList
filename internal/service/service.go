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
			User.POST("/delete", middleware.AuthMiddleware(), eh.DeleteUser)
			User.POST("/change-password", middleware.AuthMiddleware(), eh.ChangePassword)
		}
		Admin := ToDoList.Group("/admin")
		{
			Admin.POST("/delete", middleware.AuthMiddleware(), middleware.AdministratorVerifiesMiddleware(), eh.DeleteUser)
		}
		Wish := ToDoList.Group("/wish")
		{
			Wish.GET("", middleware.AuthMiddleware(), eh.GetWishes)
			Wish.GET("/random", middleware.AuthMiddleware(), eh.RandomlySelectWish)
			Wish.POST("/delete", middleware.AuthMiddleware(), eh.DeleteWish)
			Wish.POST("/update", middleware.AuthMiddleware(), eh.UpdateWish)
			Wish.POST("/add", middleware.AuthMiddleware(), eh.AddWish)
			Wish.POST("/add-todo", middleware.AuthMiddleware(), eh.AddToTodo)
		}
		ToDoList.GET("", middleware.AuthMiddleware(), eh.GetAllTodo)
		ToDoList.POST("/add", middleware.AuthMiddleware(), eh.CreateTodo)
		ToDoList.POST("/updateImportanceLevel", middleware.AuthMiddleware(), eh.SaveAllTodos)
		ToDoList.POST("delete", middleware.AuthMiddleware(), eh.DeleteTodo)
		ToDoList.POST("/update", middleware.AuthMiddleware(), eh.UpdateTodo)
		Community := ToDoList.Group("/community")
		{
			Community.GET("", eh.GetCommunityWishes)
			Community.POST("/add-viewed", eh.AddView)
			Community.POST("/add-to-wish", middleware.AuthMiddleware(), eh.AddToWish)
		}
	}
}
