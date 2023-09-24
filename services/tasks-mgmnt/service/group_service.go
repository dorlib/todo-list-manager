package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/lightbits-intel/neon_insights/services/account-mgmt/model/request"
	"github.com/lightbits-intel/neon_insights/services/account-mgmt/model/response"
	"github.com/lightbitslabs/neon_insight/libs/microservice_kit/ctx"
)

type GroupService interface {
	CreateGroup(context.Context, *ctx.RequestContext, *request.CreateGroupRequest) (*response.GroupDto, error)
	GetGroups(context.Context, *ctx.RequestContext) ([]*response.GroupDto, error)
	GetGroup(context.Context, *ctx.RequestContext, uuid.UUID) (*response.GroupDto, error)
	DeleteGroup(context.Context, *ctx.RequestContext, uuid.UUID) (*response.GroupDto, error)
	UpdateGroup(context.Context, *ctx.RequestContext, *request.UpdateGroupRequest) (*response.GroupDto, error)
}
