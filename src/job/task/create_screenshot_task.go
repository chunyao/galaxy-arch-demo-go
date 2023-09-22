package task

import (
	"app/src/model"
	"app/src/service/impl"
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"sync"
)

type CreateScreenshotTask struct {
	screenshotService impl.ScreenshotServiceImpl
}

func (createTask *CreateScreenshotTask) Run(cxt context.Context, param *xxl.RunReq) (msg string) {
	accountList := createTask.screenshotService.CreateConvertTask()
	wg := sync.WaitGroup{}

	wg.Add(len(*accountList))
	for _, item := range *accountList {
		go func(item model.ReportSubscribe) {
			defer wg.Done()
			createTask.screenshotService.AddTask(&item)

		}(item)
	}
	wg.Wait()
	return "success"
}
