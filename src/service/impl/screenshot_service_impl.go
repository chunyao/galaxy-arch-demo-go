package impl

import (
	"app/src/common/utils/screenshot_util"
	textwatermark "app/src/common/utils/watermark"
	"app/src/dao"
	"app/src/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type ScreenshotServiceImpl struct {
}

func (ScreenshotServiceImpl) Convert(user *screenshot_util.UserSrc) string {
	var obj screenshot_util.ScreenShotUtil
	obj.Init(user)
	SavePath := "./data"
	str := textwatermark.FontInfo{30, user.UserName, textwatermark.Center, 20, 20, 255, 255, 0, 255}
	arr := make([]textwatermark.FontInfo, 0)
	arr = append(arr, str)
	//加水印图片路径
	// fileName := "123123.jpg"
	fileName := obj.Path
	w := new(textwatermark.Water)
	w.Pattern = "2006/01/02"
	textwatermark.Ttf = "./resources/font/MSYH.TTC" //字体路径
	url, err := w.New(SavePath, fileName, arr)
	if err != nil {
		fmt.Println(err)
	}
	return url
}

func (ScreenshotServiceImpl) GetAccountInfo(account string) *dao.ReportSubscribeDao {
	reportSubscribe := &dao.ReportSubscribeDao{}
	reportSubscribe.SelectByAccount(account)
	return reportSubscribe
}

func (ScreenshotServiceImpl) AddAccountInfo(account *model.ReportSubscribe) *dao.ReportSubscribeDao {
	reportSubscribe := &dao.ReportSubscribeDao{}
	reportSubscribe.AddReportSubscribe(account)
	return reportSubscribe
}

func (ScreenshotServiceImpl) CreateConvertTask() *[]model.ReportSubscribe {
	reportSubscribe := &dao.ReportSubscribeDao{}
	reportSubscribe.SelectByJobDate()

	return reportSubscribe.ReportSubscribeModelList
}

func (ScreenshotServiceImpl) AddTask(task *model.ReportSubscribe) {
	reportSubscribeTask := &dao.ReportSubscribeTaskDao{}
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	findRecord := model.ReportSubscribeTask{
		MabangAccount: task.MabangAccount,
		JobDate:       t.Format("2006-01-02"),
	}
	reportSubscribeTask.FindRecord(findRecord)
	log.Info(reportSubscribeTask.ReportSubscribeTaskModel)
	if reportSubscribeTask.ReportSubscribeTaskModel.MabangAccount == "" {
		record := model.ReportSubscribeTask{
			UnionId:        task.UnionId,
			MabangAccount:  task.MabangAccount,
			MabangPassword: task.MabangPassword,
			CosUrl:         task.CosUrl,
			JobDate:        t.Format("2006-01-02"),
			Status:         0,
		}
		reportSubscribeTask.AddRecord(record)
	}
}
func (ScreenshotServiceImpl) DoTask() *[]model.ReportSubscribeTask {
	reportSubscribeTask := &dao.ReportSubscribeTaskDao{}
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	findRecord := model.ReportSubscribeTask{
		JobDate: t.Format("2006-01-02"),
		Status:  0,
	}
	reportSubscribeTask.FindRecordList(findRecord)
	return reportSubscribeTask.ReportSubscribeTaskModelList

}
func (ScreenshotServiceImpl) DoTaskFinish(record model.ReportSubscribeTask) {
	reportSubscribeTask := &dao.ReportSubscribeTaskDao{}
	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))
	updateRecord := model.ReportSubscribeTask{
		JobDate:       t.Format("2006-01-02"),
		MabangAccount: record.MabangAccount,
	}
	reportSubscribeTask.UpdateRecord(updateRecord, record.CosUrl)
}
