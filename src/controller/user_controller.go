package controller

import (
	"app/src/common/config/cache"
	"app/src/dto"
	"app/src/service"
	"app/src/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
}

func UserApi(router *gin.Engine) {

	userHandler := UserHandler{
		userService: &impl.UserServiceImpl{},
	}

	userGroup := router.Group("user/")
	{
		userGroup.GET("/:id", userHandler.user)
	}
}

// 根据ID查询用户
func (userHandler UserHandler) user(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	var o interface{}
	_, _ = cache.LocalCache.Get("User"+userIdStr, &o)
	if o == nil {
		user := userHandler.userService.User(userId)
		cache.LocalCache.Set("User"+userIdStr, user.UserModel)
		o = user.UserModel
	}
	c.JSON(http.StatusOK, dto.Ok(o))
}
