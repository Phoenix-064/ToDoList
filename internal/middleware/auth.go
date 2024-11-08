package middleware

import (
	"ToDoList/internal/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Response 标准回应结构体
type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// AuthMiddleware Token验证中间件
// 使用 bearer token 认证方式
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
			ctx.Abort()
			return
		}
		token := parts[1]
		th := utils.NewTokenHandler()
		claims, err := th.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "err",
				Content: err.Error(),
			})
			logrus.Error("解析token时出错，", err)
			ctx.Abort()
			return
		}
		// logrus.WithFields(logrus.Fields{
		// 	"uuid":    claims.Uuid,
		// 	"isAdmin": claims.IsAdmin,
		// 	"token":   token[:10] + "...",
		// }).Info("Token validated successfully")
		ctx.Set("uuid", claims.Uuid)
		ctx.Set("isAdmin", claims.IsAdmin)
	}
}

func GetHeader(ctx *gin.Context) (string, bool, error) {
	uuid, existent := ctx.Get("uuid")
	if !existent {
		ctx.JSON(http.StatusUnauthorized, Response{
			Message: "err",
			Content: "不存在uuid",
		})
		return "", false, errors.New("不存在的uuid")
	}
	isAdmin, existent := ctx.Get("isAdmin")
	if !existent {
		ctx.JSON(http.StatusUnauthorized, Response{
			Message: "err",
			Content: "错误的信息",
		})
		return "", false, errors.New("不存在管理员信息")
	}
	// logrus.Info(uuid.(string))
	return uuid.(string), isAdmin.(bool), nil
}
