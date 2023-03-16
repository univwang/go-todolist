package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"task/model"
	"task/service"
)

// 创建备忘录，将备忘录信息放到消息队列中

func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	channel, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("channel error + " + err.Error())
	}
	q, _ := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req) // 将请求序列化
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("publish error + " + err.Error())
	}
	return err
}

// 实现备忘录接口 获取备忘录列表

func (*TaskService) GetTasksList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var tasks []model.Task
	var count uint32

	// 查找备忘录
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid = ?", req.Uid).Find(&tasks).Error
	if err != nil {
		err = errors.New("get tasks list error + " + err.Error())
		return err
	}
	// 统计数量
	model.DB.Model(&model.Task{}).Where("uid = ?", req.Uid).Count(&count)

	var taskResp []*service.TaskModel
	for _, task := range tasks {
		taskResp = append(taskResp, BuildTask(task))
	}
	resp.TaskList = taskResp
	resp.Count = count
	return nil
}

// 获取备忘录详细信息

func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	task := model.Task{}
	model.DB.First(&task, req.Id)
	taskResp := BuildTask(task)
	resp.TaskDetail = taskResp
	return nil
}

// 修改备忘录

func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	task := model.Task{}
	model.DB.Model(&model.Task{}).Where("id = ? AND uid = ?", req.Id, req.Uid).First(&task)
	task.Title = req.Title
	task.Content = req.Content
	task.Status = int(req.Status)
	model.DB.Save(&task)
	resp.TaskDetail = BuildTask(task)
	return nil
}

// 删除备忘录

func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	model.DB.Model(&model.Task{}).Where("id = ? AND uid = ?", req.Id, req.Uid).Delete(&model.Task{})
	return nil
}
