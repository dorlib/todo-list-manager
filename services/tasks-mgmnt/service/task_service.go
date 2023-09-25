package service

import (
	"context"
	"github.com/google/uuid"
	"tasks-mgmnt/model/request"
)

type TaskService interface {
	CreateTask(context.Context, *ctx.RequestContext, *request.CreateTaskRequest) (*response.TaskDto, error)
	GetTasks(context.Context, *ctx.RequestContext) ([]*response.TaskDto, error)
	GetTask(context.Context, *ctx.RequestContext, uuid.UUID) (*response.TaskDto, error)
	DeleteTask(context.Context, *ctx.RequestContext, uuid.UUID) (*response.TaskDto, error)
	UpdateTask(context.Context, *ctx.RequestContext, *request.UpdateTaskRequest) (*response.TaskDto, error)
}
