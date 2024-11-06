package engine

import (
	user "ToDoList/internal/User"
	"ToDoList/internal/email"
	"ToDoList/internal/utils"
	"net/http"
	"strings"

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
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Email 用于接收验证时返回的邮箱
type Email struct {
	Email string `json:"email"`
}

// EmailVerification 用于使用邮箱验证码登录
type EmailVerification struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
	Password         string `json:"password"`
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

// newEmail 返回一个新的Email
func newEmail() Email {
	return Email{}
}

// newEmailVerification 返回一个新的EmailVerification
func newEmailVerification() EmailVerification {
	return EmailVerification{}
}

// SendVerificationCode 发送验证码
func (eh EngineHandler) SendVerificationCode(ctx *gin.Context) {
	em := newEmail()
	err := ctx.ShouldBind(&em)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法连接结构体，", err)
		return
	}
	var uh user.UserHandle
	uh = user.NewUserManager()
	_, err = uh.CheckEmail(em.Email)
	if err == nil { // 找到了用户
		ctx.JSON(http.StatusUnauthorized, Response{
			Message: "err",
			Content: "已有的邮箱",
		})
		return
	} else if err.Error() != "没有此用户" { // 判断是否为服务器出错
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("查找用户时出错，", err)
		return
	}
	var vcm utils.VerificationCodeManager
	vcm = utils.NewVerificationCodeHandler()
	err = vcm.CheckTheSendingFrequency(em.Email)
	if err != nil {
		if err.Error() == "创建验证码时间间隔小于一分钟" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err.Error())
		return
	}
	emailManager := email.NewEmailManager()
	err = emailManager.ConfigureEmail(em.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("发送邮件时失败，", err)
		return
	}
	ctx.JSON(http.StatusOK, Response{
		Message: "ok",
		Content: "发送验证码成功",
	})
	// 先将验证码储存在mysql中，以后可以迭代为储存在redis中
}

// SignUp 注册
func (eh EngineHandler) SignUp(ctx *gin.Context) {
	ev := newEmailVerification()
	err := ctx.ShouldBind(&ev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法连接结构体user，", err)
		return
	}
	var vcm utils.VerificationCodeManager
	vcm = utils.NewVerificationCodeHandler()
	err = vcm.CheckTheVerificationCode(ev.Email, ev.VerificationCode)
	if err != nil {
		if err.Error() != "错误的验证码" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		return
	}
	var uh user.UserHandle
	uh = user.NewUserManager()
	u, err := user.NewUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法创建新用户，", err)
		return
	}
	u.Email = ev.Email
	u.Password = ev.Password
	err = uh.AddUser(u)
	if err != nil {
		if err.Error() == "没有此用户" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: "没有此用户",
			})
			return
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
	var uh user.UserHandle
	uh = user.NewUserManager()
	u, err := uh.CheckEmail(uq.Email)
	if err != nil { // 如果有错误返回，判断返回类型
		if err.Error() == "没有此用户" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: "没有此用户",
			})
			return
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
		th := utils.NewTokenHandler()
		tokenString, err := th.GenerateToken(u.Uuid, u.IsAdmin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "ok",
			Content: gin.H{
				"token": tokenString,
			},
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

// AuthMiddleware Token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: "缺少请求头",
			})
			ctx.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "err",
				Content: "错误的请求格式",
			})
		}
		// 未完成————————————————————————————————————————————————————————————————————————
	}
}
