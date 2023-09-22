package controller

import (
	"app/src/common/utils/screenshot_util"
	"app/src/dto"
	"app/src/service"
	"app/src/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type ScreenshotController struct {
	screenshotService service.ScreenshotService
}

func ScreenshotApi(router *gin.Engine) {

	screenshotController := ScreenshotController{
		screenshotService: &impl.ScreenshotServiceImpl{},
	}

	screenGroup := router.Group(viper.GetString("server.appName") + "/api/")
	{
		screenGroup.POST("/screen", screenshotController.convert)
		screenGroup.GET("/account/:id", screenshotController.getAccountInfo)
	}
}

// 根据ID查询用户 Redis 使用Demo
func (screenshotController ScreenshotController) convert(ctx *gin.Context) {
	var user screenshot_util.UserSrc
	ctx.BindJSON(&user)
	data := screenshotController.screenshotService.Convert(&user)
	ctx.JSON(http.StatusOK, dto.Ok(data))
}
func (screenshotController ScreenshotController) getAccountInfo(ctx *gin.Context) {
	accountStr := ctx.Param("id")
	data := screenshotController.screenshotService.GetAccountInfo(accountStr)
	ctx.JSON(http.StatusOK, dto.Ok(data))
}
