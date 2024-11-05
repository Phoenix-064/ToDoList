package engine

import (
	user "ToDoList/internal/User"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// EnginHandler gin框架下的控制器
type EngineHandler struct {
}

// Response 标准回应结构体
type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
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

// NewEngineHandler 返回一个EngineHandler
func NewEngineHandler() EngineHandler {
	return EngineHandler{}
}

// newUserRequest 返回一个新的UserRequest
func newUserRequest() UserRequest {
	return UserRequest{}
}

// SignUp 注册
func (eh EngineHandler) SignUp(ctx *gin.Context) {
	uq := newUserRequest()
	err := ctx.ShouldBind(&uq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法连接结构体user，", err)
		return
	}
	dsn := "root:123@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	var um user.UserHandle
	um = user.NewUserManager(dsn)
	u, err := user.NewUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法创建新用户，", err)
		return
	}
	u.Email = uq.Email
	u.Name = uq.Name
	u.Password = uq.Password
	err = um.AddUser(u)
	if err != nil {
		if err.Error() == "没有此用户" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: "没有此用户",
			})
			logrus.Info("没有此用户")
		} else {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "err",
				Content: err.Error(),
			})
			logrus.Error("添加用户出错，", err)
			return
		}
	}
	ctx.JSON(http.StatusOK, Response{
		Message: "ok",
		Content: "",
	})
}

// SignIn 登录
func (eh EngineHandler) SignIn(ctx *gin.Context) {
	uq := newUserRequest()
	err := ctx.ShouldBind(&uq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("连接结构体User失败，", err)
		return
	}
	dsn := "root:123@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	var um user.UserHandle
	um = user.NewUserManager(dsn)
	u, err := um.CheckUser(uq.Email)
	if err != nil { // 如果有错误返回，判断返回类型
		if err.Error() == "没有此用户" {
			u, err := um.CheckUser(uq.Name)
			if err != nil { // 如果错误返回还不为空，那么要不是出错了，要不是不存在此用户
				ctx.JSON(http.StatusBadRequest, Response{
					Message: "err",
					Content: err.Error(),
				})
				// 直接返回错误就可以了，然后发现要记录错误
				if err.Error() == "没有此用户" {
					logrus.Info("没有此用户")
					return
				} else {
					logrus.Error("查找用户时出错，", err)
					return
				}
			}
			// 查找姓名时未出错，则判断密码是否正确
			if u.Password == uq.Password {
				// 这一部分用于判断邮箱是否有误，但好像没用
				// if u.Email == uq.Email {
				// 	ctx.JSON(http.StatusOK, Response{
				// 		Message: "ok",
				// 		Content: nil,
				// 	})
				// 	return
				// } else {
				// 	ctx.JSON(http.StatusUnauthorized, Response{
				// 		Message: "err",
				// 		Content: "邮箱错误",
				// 	})
				// 	return
				// }
			} else {
				ctx.JSON(http.StatusUnauthorized, Response{
					Message: "err",
					Content: "密码错误",
				})
				return
			}
		} else { // 错误类型不为用户不存在，则应该是出错了
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "err",
				Content: err.Error(),
			})
			logrus.Error("查找用户时出错，", err)
			return
		}
	}
	// 查找邮箱时未出错，则判断密码是否正确
	if u.Password == uq.Password {
		ctx.JSON(http.StatusOK, Response{
			Message: "ok",
			Content: nil,
		})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, Response{
			Message: "err",
			Content: "密码错误",
		})
		logrus.Info("密码错误")
		return
	}
}
