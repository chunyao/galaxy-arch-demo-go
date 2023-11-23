package service

type DeadQueueService interface {
	Save(msg []byte) bool
}
