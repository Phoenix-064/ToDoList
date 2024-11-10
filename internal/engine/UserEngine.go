package engine

import (
	data "ToDoList/internal/Data"
	user "ToDoList/internal/User"
	"ToDoList/internal/email"
	"ToDoList/internal/middleware"
	"ToDoList/internal/models"
	"ToDoList/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// EnginHandler gin框架下的控制器
type EngineHandler struct {
	TodoManager data.HandleTodo
	UserManager user.HandleUser
	WishManager data.HandleWish
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

// NewEngineHandler 返回一个EngineHandler
func NewEngineHandler(db *gorm.DB) EngineHandler {
	return EngineHandler{
		TodoManager: data.NewTodoGormManager(db),
		UserManager: user.NewUserManager(db),
		WishManager: data.NewWishManager(db),
	}
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
func (eh *EngineHandler) SendVerificationCode(ctx *gin.Context) {
	em := newEmail()
	err := ctx.ShouldBind(&em)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法连接结构体，", err)
		return
	}
	_, err = eh.UserManager.CheckEmail(em.Email)
	if err == nil { // 找到了用户
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Message: "err",
			Content: "已有的邮箱",
		})
		return
	} else if err.Error() != "没有此用户" { // 判断是否为服务器出错
		ctx.JSON(http.StatusInternalServerError, models.Response{
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
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err.Error())
		return
	}
	emailManager := email.NewEmailManager()
	err = emailManager.ConfigureEmail(em.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("发送邮件时失败，", err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "发送验证码成功",
	})
	// 先将验证码储存在mysql中，以后可以迭代为储存在redis中
}

// SignUp 注册
func (eh *EngineHandler) SignUp(ctx *gin.Context) {
	ev := newEmailVerification()
	err := ctx.ShouldBind(&ev)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
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
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		return
	}
	u, err := user.NewUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("无法创建新用户，", err)
		return
	}
	u.Email = ev.Email
	u.Password = ev.Password
	err = eh.UserManager.AddUser(u)
	if err != nil {
		if err.Error() == "没有此用户" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Message: "err",
				Content: "没有此用户",
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "err",
				Content: err.Error(),
			})
			logrus.Error("添加用户出错，", err)
			return
		}
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "",
	})
}

// SignIn 登录
func (eh *EngineHandler) SignIn(ctx *gin.Context) {
	uq := newUserRequest()
	err := ctx.ShouldBind(&uq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error("连接结构体User失败，", err)
		return
	}
	u, err := eh.UserManager.CheckEmail(uq.Email)
	if err != nil { // 如果有错误返回，判断返回类型
		if err.Error() == "没有此用户" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Message: "err",
				Content: "没有此用户",
			})
			return
		} else { // 错误类型不为用户不存在，则应该是出错了
			ctx.JSON(http.StatusInternalServerError, models.Response{
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
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "err",
				Content: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Message: "ok",
			Content: gin.H{
				"token": tokenString,
			},
		})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Message: "err",
			Content: "密码错误",
		})
		return
	}
}

// AdminDeleteUser 删除用户
func (eh *EngineHandler) AdminDeleteUser(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	if err := eh.UserManager.DeleteUser(uuid); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Message: "ok",
		Content: "删除成功",
	})
}

// DeleteUser 删除用户
func (eh *EngineHandler) DeleteUser(ctx *gin.Context) {
	uuid, _, err := middleware.GetHeader(ctx)
	logrus.Info(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "err",
			Content: err.Error(),
		})
		logrus.Error(err)
		return
	}
	err = eh.UserManager.DeleteUser(uuid)
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
		Content: "删除成功",
	})
}
