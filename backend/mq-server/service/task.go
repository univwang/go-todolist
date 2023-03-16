package service

import (
	"encoding/json"
	"log"
	"mq-server/model"
)

// 从RabbitMQ中获取任务，写入数据库

func CreateTask() {
	channel, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	err = channel.Qos(1, 0, false)
	msg, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 消费者处于监听状态，我们要阻塞主进程
	go func() {
		for d := range msg {
			// 将消息反序列化
			// 将消息写入数据库
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			d.Ack(false) //确认已被消费
		}
	}()
}
