package listener

import (
	"app/src/service"
	log "github.com/sirupsen/logrus"
)

type DeadListenerConsumer struct {
	DeadQueueService service.DeadQueueService
}

func (deadConsumer DeadListenerConsumer) Do(msg []byte) bool {
	log.Printf("Received a message: %s", msg)
	deadConsumer.DeadQueueService.Save(msg)
	return true
}
