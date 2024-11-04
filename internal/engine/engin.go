package engine

import (
	user "ToDoList/internal/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EnginHandler gin框架下的控制器
type EngineHandler struct {
}

// Response 标准回应结构体
type Response struct {
	message string
	content interface{}
}

// UserRequest 标准用户信息请求体
type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NetUserHandler interface {
	AddUser()
	CheckUser()
	DeleteUser()
	UpdateUser()
}

// newUserRequest 返回一个新的UserRequest
func newUserRequest() UserRequest {
	return UserRequest{}
}

func (e EngineHandler) AddUser(ctx *gin.Context) error {
	uq := newUserRequest()
	err := ctx.ShouldBind(&uq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			message: err.Error(),
		})
		return err
	}
	user.UserList()
	return nil
}
