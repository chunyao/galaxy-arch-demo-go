package controller

import (
	"app/src/common/utils/image_util"
	"app/src/dto"
	"app/src/service"
	"app/src/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type ImageController struct {
	imageService service.ImageService
}

func ImageApi(router *gin.Engine) {

	imageController := ImageController{
		imageService: &impl.ImageServiceImpl{},
	}

	imageGroup := router.Group(viper.GetString("server.appName") + "/api/")
	{
		imageGroup.POST("/image", imageController.convert)
	}

}
func (imageController ImageController) convert(ctx *gin.Context) {
	var img []image_util.ImageSrc
	ctx.BindJSON(&img)
	newImg := imageController.imageService.Convert(&img)
	ctx.JSON(http.StatusOK, dto.Ok(newImg))
}
