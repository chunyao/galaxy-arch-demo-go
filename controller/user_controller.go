package controller

import (
	"github.com/gin-gonic/gin"
	"mabang-arch-demo-go/common/config/redis"
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
	redis.Obj().Set(c, "User"+userIdStr, "123123", 60)

	var o interface{}
	redis.Obj().Get(c, "User"+userIdStr, &o)

	c.JSON(http.StatusOK, dto.Ok(o))
}
