package model

type DeadQueueConsumer struct {
	Id           int    `json:"id" gorm:"primary_key"`
	BusinessCode string `json:"business_code" gorm:"size:255"`
	QueueName    string `json:"queue_name" gorm:"size:255"`
	Body         string `json:"body"`
	/**
	 * 0 初始化 1 已发送  2发送失败  3 抛弃
	 */
	Status     int    `json:"status" gorm:"size:1"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
