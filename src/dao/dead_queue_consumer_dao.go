package dao

import (
	"app/src/common/config/db"
	"app/src/model"
	log "github.com/sirupsen/logrus"
	"time"
)

type DeadQueueConsumerDao struct {
	DeadQueueConsumerModel model.DeadQueueConsumer
}

func (deadQueueConsumerDao *DeadQueueConsumerDao) AddRecord(record model.DeadQueueConsumer) {
	//orm 原生
	date := time.Now().Format("2006-01-02")
	record.CreateTime = date
	record.UpdateTime = date
	if err := db.DBs["db"].Db.Create(&record).Error; err != nil {
		log.Error("消息插入失败", err)
	}
}
