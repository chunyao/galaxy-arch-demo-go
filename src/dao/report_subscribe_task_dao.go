package dao

import (
	"app/src/common/config/db"
	"app/src/model"
	log "github.com/sirupsen/logrus"
)

type ReportSubscribeTaskDao struct {
	ReportSubscribeTaskModel     *model.ReportSubscribeTask
	ReportSubscribeTaskModelList *[]model.ReportSubscribeTask
}

func (reportSubscribeTaskDao *ReportSubscribeTaskDao) AddRecord(record model.ReportSubscribeTask) {
	//orm 查询用法
	if err := db.DBs["db"].Db.Create(&record).Error; err != nil {
		log.Error("任务插入失败", err)
	}
}
func (reportSubscribeTaskDao *ReportSubscribeTaskDao) FindRecord(record model.ReportSubscribeTask) {
	//orm 查询用法
	db.DBs["db"].Db.Where(record).First(&reportSubscribeTaskDao.ReportSubscribeTaskModel)
}
func (reportSubscribeTaskDao *ReportSubscribeTaskDao) FindRecordList(record model.ReportSubscribeTask) {
	//orm 查询用法
	db.DBs["db"].Db.Where(record).Find(&reportSubscribeTaskDao.ReportSubscribeTaskModelList)
}

func (reportSubscribeTaskDao *ReportSubscribeTaskDao) UpdateRecord(record model.ReportSubscribeTask, cosUrl string) {
	//orm 查询用法
	db.DBs["db"].Db.Where(record).Updates(model.ReportSubscribeTask{Status: 1, CosUrl: cosUrl})
}
