package controller

import (
	"app/src/common/config/cache"
	"app/src/dto"
	"app/src/service"
	"app/src/service/impl"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

// 根据ID查询用户 Redis 使用Demo
func (userHandler UserHandler) user(ctx *gin.Context) {
	userIdStr := ctx.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	var data interface{}
	//_, _ = cache.LocalCache.Get("User"+userIdStr, &o)
	o, _ := cache.Redis.Get(ctx, "User:"+userIdStr).Result()
	fmt.Println("sdf ", o)
	if len(o) == 0 {
		user := userHandler.userService.User(userId)
		paramJson, _ := json.Marshal(user.UserModel)
		cache.Redis.Set(ctx, "User:"+userIdStr, string(paramJson), 60*time.Second)
		o = string(paramJson)

	}

	json.Unmarshal([]byte(o), &data)
	fmt.Println(data)
	ctx.JSON(http.StatusOK, dto.Ok(data))
}
