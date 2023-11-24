package rabbitMq

import (
	"app/src/common/utils"
	"app/src/listener"
	"app/src/service/impl"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
) //导入mq包

// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost (默认是5672端口)
// 端口可在 /etc/rabbitMq/rabbitMq-env.conf 配置文件设置，也可以启动后通过netstat -tlnp查看
const MQURL = ""

type Rabbit struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// routing Key
	RoutingKey string
	//MQ链接字符串
	Mqurl string
}

var RabbitMq Rabbit

func InitRabbitMqDeadConsumer() {

	connct := "amqp://" + viper.GetString("rabbitmq.username") + ":" +
		viper.GetString("rabbitmq.password") + "@" + viper.GetString("rabbitmq.host") + ":" +
		viper.GetString("rabbitmq.port") + "/" + viper.GetString("rabbitmq.vhost")
	log.Info("初始化RabbitMq 消费者")

	for _, queueName := range viper.GetStringMapString("rabbitmq.queue.dead") {
		go func(queueName string) {
			var i = 0

			conn, err := amqp.Dial(connct)
			checkErr(err, "创建连接失败")
			for i = 0; i <= 10; i++ {
				go func() {
					deadListenerConsumer := listener.DeadListenerConsumer{
						DeadQueueService: &impl.DeadQueueServiceImpl{},
					}
					ch, err := conn.Channel()
					checkErr(err, "创建Channel失败")
					ch.Qos(20, 0, false)
					go subscribe(ch, utils.NewUUID(), queueName, func(msgs <-chan amqp.Delivery, s string) {
						for msg := range msgs {
							if deadListenerConsumer.Do(msg.Body) == true {
								msg.Ack(false)
							} else {
								msg.Nack(false, true)
							}
						}
					})
				}()
			}
			log.Info("创建消费者" + queueName)
		}(queueName)

	}

}
func subscribe(ch *amqp.Channel, key string, queue string, callback func(<-chan amqp.Delivery, string)) {
	msgs, err := ch.Consume(queue, key, false, false, false, false, nil)

	if err != nil {
		checkErr(err, "获取消息失败")
	}
	callback(msgs, key)
}

// 创建结构体实例
func RabbitMQProduce(queueName, exchange, routingKey string) *Rabbit {
	connect := "amqp://" + viper.GetString("rabbitmq.username") + ":" + viper.GetString("rabbitmq.password") + "@" + viper.GetString("rabbitmq.host") + ":" + viper.GetString("rabbitmq.port") + "/" + viper.GetString("rabbitmq.vhost")

	rabbitMQ := Rabbit{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      connect,
	}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)

	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return &rabbitMQ

}

// 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *Rabbit) ReleaseRes() {
	mq.Conn.Close()
	mq.Channel.Close()
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
