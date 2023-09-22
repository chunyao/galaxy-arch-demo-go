package task

import (
	"app/src/common/utils/screenshot_util"
	"app/src/service/impl"
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
)

type DoScreenshotTask struct {
	screenshotService impl.ScreenshotServiceImpl
}

func (doTask *DoScreenshotTask) Run(cxt context.Context, param *xxl.RunReq) (msg string) {
	accountList := doTask.screenshotService.DoTask()
	//	wg := sync.WaitGroup{}
	//	wg.Add(len(*accountList))
	for _, item := range *accountList {
		//go func(item model.ReportSubscribeTask) {
		//	defer wg.Done()
		user := screenshot_util.UserSrc{
			UserName: item.MabangAccount,
			Password: item.MabangPassword,
			FileName: "./data/" + item.MabangAccount + ".png",
		}
		url := doTask.screenshotService.Convert(&user)
		item.CosUrl = url
		doTask.screenshotService.DoTaskFinish(item)
		//	}(item)
	}
	//	wg.Wait()
	return "success"
}
