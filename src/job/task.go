package job

import (
	"app/src/job/task"
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
)

type Task struct {
	CreateScreenshotTaskHandler task.CreateScreenshotTask
	DoScreenshotTaskHandler     task.DoScreenshotTask
}

func (task Task) DoTask() (dispatch map[string]func(cxt context.Context, param *xxl.RunReq) string) {
	dispatch = make(map[string]func(cxt context.Context, param *xxl.RunReq) string)
	//注册job
	dispatch["CreateScreenshotTaskHandler"] = task.CreateScreenshotTaskHandler.Run
	dispatch["DoScreenshotTaskHandler"] = task.DoScreenshotTaskHandler.Run

	return dispatch
}
