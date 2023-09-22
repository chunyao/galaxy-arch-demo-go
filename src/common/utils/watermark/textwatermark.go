package textwatermark

import (
	"app/src/common/config/cos"
	"app/src/common/utils"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/anthonynsimon/bild/transform"
	"github.com/golang/freetype"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"time"
)

// 水印的位置
const (
	TopLeft int = iota
	TopRight
	BottomLeft
	BottomRight
	Center
)

// 字体路径
var Ttf string

type Water struct {
	Pattern string //增加按时间划分的子目录：默认没有时间划分的子目录
}

func (w *Water) New(SavePath, fileName string, typeface []FontInfo) (url string, err error) {
	var subPath string
	subPath = w.Pattern
	dirs, err := createDir(SavePath, subPath)
	if err != nil {
		return "", err
	}
	imgfile, _ := os.Open(fileName)
	defer imgfile.Close()
	_, str, err := image.DecodeConfig(imgfile)
	if err != nil {
		return "", err
	}

	newName := fmt.Sprintf("%s%s.%s", dirs, getRandomString(10), str)
	if str == "gif" {
		err = gifFontWater(fileName, newName, typeface)
	} else {
		url, _ = staticFontWater(fileName, newName, str, typeface)
	}
	return url, err
}

// gif图片水印
func gifFontWater(file, name string, typeface []FontInfo) (err error) {
	imgfile, _ := os.Open(file)
	defer imgfile.Close()
	var err2 error
	gifimg2, _ := gif.DecodeAll(imgfile)
	gifs := make([]*image.Paletted, 0)
	x0 := 0
	y0 := 0
	yuan := 0
	for k, gifimg := range gifimg2.Image {
		img := image.NewNRGBA(gifimg.Bounds())
		if k == 0 {
			x0 = img.Bounds().Dx()
			y0 = img.Bounds().Dy()
		}
		fmt.Printf("%v, %v\n", img.Bounds().Dx(), img.Bounds().Dy())
		if k == 0 && gifimg2.Image[k+1].Bounds().Dx() > x0 && gifimg2.Image[k+1].Bounds().Dy() > y0 {
			yuan = 1
			break
		}
		if x0 == img.Bounds().Dx() && y0 == img.Bounds().Dy() {
			for y := 0; y < img.Bounds().Dy(); y++ {
				for x := 0; x < img.Bounds().Dx(); x++ {
					img.Set(x, y, gifimg.At(x, y))
				}
			}
			waterBytes, _ := common(typeface) //添加文字水印
			waterfile, _ := os.Create("water.png")
			waterPng, _ := png.Decode(bytes.NewReader(waterBytes))
			err = png.Encode(waterfile, waterPng)
			if err2 != nil {
				break
			}
			//定义一个新的图片调色板img.Bounds()：使用原图的颜色域，gifimg.Palette：使用原图的调色板
			p1 := image.NewPaletted(gifimg.Bounds(), gifimg.Palette)
			//把绘制过文字的图片添加到新的图片调色板上
			draw.Draw(p1, gifimg.Bounds(), img, image.ZP, draw.Src)
			//把添加过文字的新调色板放入调色板slice
			gifs = append(gifs, p1)
		} else {
			gifs = append(gifs, gifimg)
		}
	}
	if yuan == 1 {
		return errors.New("gif: image block is out of bounds")
	} else {
		if err2 != nil {
			return err2
		}
		//保存到新文件中
		newfile, err := os.Create(name)
		if err != nil {
			return err
		}
		defer newfile.Close()
		g1 := &gif.GIF{
			Image:     gifs,
			Delay:     gifimg2.Delay,
			LoopCount: gifimg2.LoopCount,
		}
		err = gif.EncodeAll(newfile, g1)
		return err
	}
}

// png,jpeg图片水印
func staticFontWater(file, name, status string, typeface []FontInfo) (url string, err error) {
	//需要加水印的图片
	imgfile, _ := os.Open(file)
	imgbinfo, _ := png.Decode(imgfile)

	defer imgfile.Close()
	waterBytes, err := common(typeface) //添加文字水印
	imgwatermark, err := png.Decode(bytes.NewReader(waterBytes))
	b := imgbinfo.Bounds()
	watermarkbg := image.NewNRGBA(image.Rect(0, 0, b.Dx()*4, b.Dy()*4))

	m := image.NewNRGBA(b) //按原图生成新图

	draw.Draw(m, b, imgbinfo, image.ZP, draw.Src) //写入原图

	x, y := 0, 0
	offsetX, offsetY := 3, 25
	maxX := watermarkbg.Bounds().Max.X * 4
	maxY := watermarkbg.Bounds().Max.Y * 4
	i := 0
	for y <= maxY {
		for x <= maxX {
			offset := image.Pt(x, y)
			offset.X = offset.X
			draw.Draw(watermarkbg, imgwatermark.Bounds().Add(offset), imgwatermark, image.ZP, draw.Over)
			x += imgwatermark.Bounds().Dx()
			x += offsetX
		}
		y += imgwatermark.Bounds().Dy()
		y += offsetY
		x = 0
		i++
	}
	watermarkbg2 := image.Image(watermarkbg)
	watermarkbg2 = transform.Rotate(watermarkbg2, -40.0, nil)
	mask := image.NewUniform(color.Alpha{30})
	draw.DrawMask(m, watermarkbg2.Bounds().Add(image.Pt(-watermarkbg2.Bounds().Dx()/2, 0)), watermarkbg2, image.ZP, mask, image.Point{-100, -100}, draw.Over)

	if err != nil {
		return "", err
	}
	//保存到新文件中
	buf := new(bytes.Buffer)
	if err != nil {
		return "", err
	}

	err = png.Encode(buf, m)

	cosName := utils.NewUUID()
	log.Info("/screenshot/" + cosName + ".png")
	reader := bytes.NewReader(buf.Bytes())

	resp, err := cos.Client.Object.Put(context.Background(), "/screenshot/"+cosName+".png", reader, nil)

	log.Info(resp.StatusCode)
	return viper.GetString("tencent.baseUrl") + "screenshot/" + cosName + ".png", err
}

// 添加文字水印函数
func common(typeface []FontInfo) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, 500, 200))

	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)
	var err2 error
	//拷贝一个字体文件到运行目录
	fontBytes, err := os.ReadFile(Ttf)
	if err != nil {
		err2 = err
		return nil, err2
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		err2 = err
		return nil, err2
	}
	errNum := 1
	for _, t := range typeface {
		info := t.Message
		f := freetype.NewContext()
		f.SetDPI(108)
		f.SetFont(font)
		f.SetFontSize(t.Size)
		f.SetClip(img.Bounds())
		f.SetDst(img)
		f.SetSrc(image.NewUniform(color.RGBA{R: t.R, G: t.G, B: t.B, A: t.A}))

		pt := freetype.Pt(10, 10+int(f.PointToFixed(20)>>6))

		_, err = f.DrawString(info, pt)
		if err != nil {
			err2 = err
			break
		}
	}
	if errNum == 0 {
		err2 = errors.New("坐标值不对")
	}
	var buff bytes.Buffer

	png.Encode(&buff, img)
	return buff.Bytes(), err2
}

// 定义添加的文字信息
type FontInfo struct {
	Size     float64 //文字大小
	Message  string  //文字内容
	Position int     //文字存放位置
	Dx       int     //文字x轴留白距离
	Dy       int     //文字y轴留白距离
	R        uint8   //文字颜色值RGBA中的R值
	G        uint8   //文字颜色值RGBA中的G值
	B        uint8   //文字颜色值RGBA中的B值
	A        uint8   //文字颜色值RGBA中的A值
}

// 生成图片名字
func getRandomString(lenght int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	bytesLen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// 检查并生成存放图片的目录
func createDir(SavePath, subPath string) (string, error) {
	var dirs string
	if subPath == "" {
		dirs = fmt.Sprintf("%s/", SavePath)
	} else {
		dirs = fmt.Sprintf("%s/%s/", SavePath, time.Now().Format(subPath))
	}
	_, err := os.Stat(dirs)
	if err != nil {
		err = os.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return dirs, nil
}
