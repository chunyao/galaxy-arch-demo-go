package service

import (
	"app/src/common/utils/image_util"
)

type ImageService interface {
	Convert(img *[]image_util.ImageSrc) []*image_util.ImageUtil
}
