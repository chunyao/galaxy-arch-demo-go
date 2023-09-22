package screenshot_util

import (
	_ "bytes"
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
	_ "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "io/ioutil"
	"os"
	_ "os"
	"time"
)

type UserSrc struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FileName string `json:"fileName"`
}
type ScreenShotUtil struct {
	Path string `json:"path"` //path
}

func (screenShotUtil *ScreenShotUtil) Init(user *UserSrc) *ScreenShotUtil {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置浏览器选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	// 创建一个浏览器实例
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// 导航到指定的URL
	var buf []byte
	err := chromedp.Run(ctx, chromedp.Navigate("https://www.mabangerp.com"),
		chromedp.Sleep(2*time.Second),
		chromedp.Click(`document.querySelector("#login-btn")`, chromedp.ByJSPath),
		chromedp.SendKeys(`account`, user.UserName, chromedp.ByID),
		chromedp.SendKeys(`password`, user.Password, chromedp.ByID),
		chromedp.Click(`document.querySelector("#account-btn")`, chromedp.ByJSPath),
		chromedp.Sleep(2*time.Second),
		//chromedp.Click(`document.querySelector("#mb-nav > li:nth-child(15) > a")`, chromedp.ByJSPath),
		//chromedp.Click(`document.querySelector("#M0012200MenuId > div > div > div > div > div.con > div > a")`, chromedp.ByJSPath),
		chromedp.Navigate("https://www.mabangerp.com/index.php?mod=main.cloudbi"),
		chromedp.Sleep(10*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			//定位登录按钮
			//

			// 获取页面截图
			var err error
			buf, err = page.CaptureScreenshot().WithQuality(90).WithClip(&page.Viewport{X: 0, Y: 0, Width: 1920, Height: 1080, Scale: 1}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}))
	if err != nil {
		log.Error("访问失败", err)
	}
	// 将截图保存到文件
	err = os.WriteFile(user.FileName, buf, 0644)
	if err != nil {
		log.Error("保存失败", err)
	}
	screenShotUtil.Path = user.FileName
	return screenShotUtil
}
