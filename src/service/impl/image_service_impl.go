package impl

import (
	"app/src/common/utils/image_util"
	"sync"
)

type ImageServiceImpl struct {
}

func (ImageServiceImpl) Convert(img *[]image_util.ImageSrc) []*image_util.ImageUtil {
	wg := sync.WaitGroup{}
	var data []*image_util.ImageUtil

	wg.Add(len(*img))
	for _, item := range *img {
		go func(item image_util.ImageSrc) {
			defer wg.Done()
			var obj image_util.ImageUtil
			//
			data = append(data, obj.Init(item.Url, item.Path))

		}(item)
	}
	wg.Wait()

	return data
}
