package service

import (
	"app/src/common/utils/screenshot_util"
	"app/src/dao"
	"app/src/model"
)

type ScreenshotService interface {
	Convert(user *screenshot_util.UserSrc) string
	GetAccountInfo(account string) *dao.ReportSubscribeDao
	CreateConvertTask() *[]model.ReportSubscribe
	AddTask(task *model.ReportSubscribe)
	DoTask() *[]model.ReportSubscribeTask
	DoTaskFinish(record model.ReportSubscribeTask)
}
