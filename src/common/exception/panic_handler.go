package exception

import (
	"app/src/common/config/token"
	"app/src/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

// ErrHandle 统一异常处理
func ErrHandle(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			apiErr, isApiErr := r.(*ApiError)
			tokenErr, isTokenErr := r.(*token.TokenError)
			if isApiErr {
				// 打印错误堆栈信息
				log.WithField("ErrMsg", apiErr.Error()).Error("PanicHandler handled apiError: ")
				// 封装通用json返回
				c.JSON(http.StatusInternalServerError, apiErr)
			} else if isTokenErr {
				// 打印错误堆栈信息
				log.WithField("ErrMsg", tokenErr.Error()).Error("PanicHandler handled tokenError: ")
				// 封装通用json返回
				c.JSON(http.StatusUnauthorized, tokenErr)
			} else {
				// 打印错误堆栈信息
				err := r.(error)
				log.WithField("ErrMsg", err.Error()).Error("PanicHandler handled ordinaryError: ")
				debug.PrintStack()
				// 封装通用json返回
				c.JSON(http.StatusInternalServerError, NewApiError(dto.InternalServerError, dto.GetResultMsg(dto.InternalServerError)))
			}
			c.Abort()
		}
	}()
	c.Next()
}
