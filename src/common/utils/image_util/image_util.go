package image_util

import (
	"app/src/common/config/http"
	"bytes"
	_ "bytes"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	_ "golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	_ "golang.org/x/image/webp"
	"image"
	_ "image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	_ "io/ioutil"
	"os"
	_ "os"
	"strings"
)

type ImageSrc struct {
	Url     string `json:"url"`
	Path    string `json:"path"`
	Convert bool   `json:"convert"`
}
type ImageUtil struct {
	Width  int    `json:"width"`  //宽
	Height int    `json:"height"` //高
	Size   int    `json:"size"`   //size
	Type   string `json:"type"`   //type
	Path   string `json:"path"`   //path
}

func (imageUtil *ImageUtil) Init(fileName string, newPath string) *ImageUtil {
	var option jpeg.Options
	fmt.Println(" downloading", fileName)
	option.Quality = 100
	response, err := http2.HttpClient.Get(fileName)
	if err != nil {
		fmt.Println("Error while downloading", fileName, ":", err)
		return nil
	}
	data, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		panic(err)
	}
	file, imageType, _ := image.Decode(bytes.NewReader(data))
	switch imageType {
	case `jpeg`:
		imageUtil.Type = `jpeg`
	case `png`:
		imageUtil.Type = `png`
		imageUtil.DecodeImageWidthHeight(data, imageUtil.Type)
		m := resize.Resize(uint(imageUtil.Width), 0, file, resize.Lanczos3)
		newFile, _ := os.Create(newPath)
		png.Encode(newFile, m)
		newFile.Close()
		return imageUtil
	case `gif`:
		imageUtil.Type = `gif`
	case `bmp`:
		imageUtil.Type = `bmp`
	case `webp`:
		imageUtil.Type = `webp`
	default:
		// 尝试以 webp 进行解码
		file, err := webp.Decode(bytes.NewReader(data))
		if err == nil {
			println(`这是 webp 文件`)
			imageUtil.Type = `webp`
		} else {
			println(err)
			imageUtil.Type = `error`
			println(fileName)
			return imageUtil

		}

		imageUtil.DecodeImageWidthHeight(data, imageUtil.Type)
		m := resize.Resize(uint(imageUtil.Width), 0, file, resize.Lanczos3)
		newFile, _ := os.Create(newPath)
		jpeg.Encode(newFile, m, &option)
		newFile.Close()
		return imageUtil
	}

	imageUtil.DecodeImageWidthHeight(data, imageUtil.Type)
	m := resize.Resize(uint(imageUtil.Width), 0, file, resize.Lanczos3)
	newFile, _ := os.Create(newPath)
	jpeg.Encode(newFile, m, &option)
	newFile.Close()
	imageUtil.Path = newPath
	return imageUtil
}

/**
* 入参： JPG 图片文件的二进制数据
* 出参：JPG 图片的宽和高
**/
func (imageUtil *ImageUtil) DecodeImageWidthHeight(imgBytes []byte, fileType string) (*ImageUtil, error) {
	var (
		imgConf image.Config
		err     error
	)
	switch strings.ToLower(fileType) {
	case "jpg", "jpeg":
		imgConf, err = jpeg.DecodeConfig(bytes.NewReader(imgBytes))
	case "webp":
		imgConf, err = webp.DecodeConfig(bytes.NewReader(imgBytes))
	case "png":
		imgConf, err = png.DecodeConfig(bytes.NewReader(imgBytes))
	case "tif", "tiff":
		imgConf, err = tiff.DecodeConfig(bytes.NewReader(imgBytes))
	case "gif":
		imgConf, err = gif.DecodeConfig(bytes.NewReader(imgBytes))
	case "bmp":
		imgConf, err = bmp.DecodeConfig(bytes.NewReader(imgBytes))
	default:
		return nil, errors.New("unknown file type")
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	imageUtil.Height = imgConf.Height
	imageUtil.Width = imgConf.Width
	return imageUtil, nil
}
