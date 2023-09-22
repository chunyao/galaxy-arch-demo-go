package main

import (
	"app/src/common/utils/watermark"
	"fmt"
)

func main() {
	SavePath := "./kaf"
	str := textwatermark.textwatermark{30, "我是水印", textwatermark.Center, 20, 20, 255, 255, 0, 255}
	arr := make([]textwatermark.FontInfo, 0)
	arr = append(arr, str)
	//加水印图片路径
	// fileName := "123123.jpg"
	fileName := "test1.png"
	w := new(textwatermark.textwatermark)
	w.Pattern = "2006/01/02"
	textwatermark.Ttf = "./MSYH.TTC" //字体路径
	err := w.New(SavePath, fileName, arr)
	if err != nil {
		fmt.Println(err)
	}

}
