package controller

import (
	"app/src/common/config/cache"
	"app/src/dto"
	"app/src/model"
	"app/src/service"
	"app/src/service/impl"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	userService service.UserService
	rds         cache.RDBManager
}

func UserApi(router *gin.Engine) {

	userController := UserController{
		userService: &impl.UserServiceImpl{},
	}

	userGroup := router.Group(viper.GetString("server.appName") + "/api/")
	{
		userGroup.GET("/:id", userController.user)
	}
}

// 根据ID查询用户 Redis 使用Demo
func (userController UserController) user(ctx *gin.Context) {
	userIdStr := ctx.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	var data model.User
	o, _ := cache.RDs["php"].Redis.Get(ctx, "User:"+userIdStr).Result()
	if len(o) == 0 {
		user := userController.userService.User(userId)
		paramJson, _ := json.Marshal(user.UserModel)
		cache.RDs["php"].Redis.Set(ctx, "User:"+userIdStr, string(paramJson), 60*time.Second)
		o = string(paramJson)

	}

	json.Unmarshal([]byte(o), &data)
	//userHandler.userService.SaveUserMongo(data)
	data2 := userController.userService.UserMongo(data.Id).UserDoModel
	fmt.Println(data2)
	ctx.JSON(http.StatusOK, dto.Ok(data2))
}
