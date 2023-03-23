package controller

import (
	"github.com/gin-gonic/gin"
	cache "mabang-arch-demo-go/common/config/cache"
	"mabang-arch-demo-go/dto"
	"mabang-arch-demo-go/service"
	"mabang-arch-demo-go/service/impl"
	"net/http"
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
	//userId, _ := strconv.Atoi(userIdStr)
	//user := userHandler.userService.User(userId)
	//redis.Instance().Set(c, "User"+userIdStr, "123123", 60)
	cache.Redis.Set(c, "User"+userIdStr, "123123", 60)
	var o interface{}
	cache.Redis.Get(c, "User"+userIdStr, &o)

	c.JSON(http.StatusOK, dto.Ok(o))
}
