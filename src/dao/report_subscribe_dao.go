package dao

import (
	"app/src/common/config/db"
	"app/src/model"
	log "github.com/sirupsen/logrus"
	"time"
)

type ReportSubscribeDao struct {
	ReportSubscribeModel     model.ReportSubscribe
	ReportSubscribeModelList *[]model.ReportSubscribe
}

func (reportSubscribe *ReportSubscribeDao) SelectByAccount(account string) {
	//orm 查询用法
	db.DBs["db"].Db.Where(model.ReportSubscribe{
		MabangAccount: account,
	}).First(&reportSubscribe.ReportSubscribeModel)
}
func (reportSubscribe *ReportSubscribeDao) SelectByJobDate() {
	//orm 原生
	date := time.Now().Format("2006-01-02")
	db.DBs["db"].Db.Raw("select * from report_subscribe where job_date <> ?", date).Find(&reportSubscribe.ReportSubscribeModelList)
}

func (reportSubscribe *ReportSubscribeDao) AddReportSubscribe(account *model.ReportSubscribe) {
	//orm 原生
	date := time.Now().Format("2006-01-02")
	account.JobDate = date
	if err := db.DBs["db"].Db.Create(&account).Error; err != nil {
		log.Error("任务插入失败", err)
	}
}
