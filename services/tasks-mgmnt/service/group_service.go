package service

import (
	"context"
	"github.com/google/uuid"
	"tasks-mgmnt/model/request"
)

type GroupService interface {
	CreateGroup(context.Context, *ctx.RequestContext, *request.CreateGroupRequest) (*response.GroupDto, error)
	GetGroups(context.Context, *ctx.RequestContext) ([]*response.GroupDto, error)
	GetGroup(context.Context, *ctx.RequestContext, uuid.UUID) (*response.GroupDto, error)
	DeleteGroup(context.Context, *ctx.RequestContext, uuid.UUID) (*response.GroupDto, error)
	UpdateGroup(context.Context, *ctx.RequestContext, *request.UpdateGroupRequest) (*response.GroupDto, error)
}
