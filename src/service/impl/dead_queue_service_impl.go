package impl

import (
	"app/src/dao"
	"app/src/model"
)

type DeadQueueServiceImpl struct {
}

func (deadQueueServiceImpl DeadQueueServiceImpl) Save(msg []byte) bool {
	deadQueueConsumerDao := &dao.DeadQueueConsumerDao{}
	record := model.DeadQueueConsumer{
		Body: string(msg),
	}
	deadQueueConsumerDao.AddRecord(record)
	return true
}
