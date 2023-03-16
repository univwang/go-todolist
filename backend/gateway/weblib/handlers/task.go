package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTaskList(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.Key中取出服务实例
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization")) //拿到的是当前访问用户的Id
	taskReq.Uid = uint64(claims.Id)
	// 调用服务端的函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"data":  taskResp.TaskList,
		"count": taskResp.Count,
	})
}

// CreateTask
// 创建备忘录
func CreateTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.Key中取出服务实例
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization")) //拿到的是当前访问用户的Id
	taskReq.Uid = uint64(claims.Id)
	// 调用服务端的函数
	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})

}

// GetTaskDetail
// 获取备忘录详情

func GetTaskDetail(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.Key中取出服务实例
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization")) //拿到的是当前访问用户的Id
	taskReq.Uid = uint64(claims.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) //获取传入的id
	taskReq.Id = uint64(id)
	// 调用服务端的函数
	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})

}

// UpdateTask
// 更新备忘录

func UpdateTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.Key中取出服务实例
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization")) //拿到的是当前访问用户的Id
	taskReq.Uid = uint64(claims.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) //获取传入的id
	taskReq.Id = uint64(id)
	// 调用服务端的函数
	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})

}

// DeleteTask
// 删除备忘录

func DeleteTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 从gin.Key中取出服务实例
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization")) //拿到的是当前访问用户的Id
	taskReq.Uid = uint64(claims.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) //获取传入的id
	taskReq.Id = uint64(id)
	// 调用服务端的函数
	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"data": taskResp.TaskDetail,
	})

}
