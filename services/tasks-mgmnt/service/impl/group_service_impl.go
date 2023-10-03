package impl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"microservice-kit/ctx"
	"tasks-mgmnt/model/domain"
	"tasks-mgmnt/model/request"
	"tasks-mgmnt/repo"
	"tasks-mgmnt/repo/impl"
)

type AccountServiceImpl struct {
	log         *zerolog.Logger
	accountRepo repo.AccountRepo
	userRepo    repo.UserRepo
}

func NewAccountServiceImpl(
	log *zerolog.Logger,
	db bun.IDB,
	accountRepo repo.AccountRepo,
	userRepo repo.UserRepo,
	provisioningClient client.ProvisioningClient,
) *AccountServiceImpl {
	return &AccountServiceImpl{
		log:                log,
		db:                 db,
		accountRepo:        accountRepo,
		userRepo:           userRepo,
		provisioningClient: provisioningClient,
	}
}

func CreateGroup(name, description string, users []domain.User) (domain.Group, error) {
	group := domain.Group{
		Name:        name,
		Description: description,
		Users:       users,
	}

	rows := impl.DB.Create(&group).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return group, nil
	}

	return domain.Group{}, fmt.Errorf("failed to create group")
}

func (a *GroupServiceImpl) CreateGroup(reqCtx *ctx.RequestContext, req *request.CreateGroupRequest) (*domain.Group, error) {
	groupID := uuid.New()

	group := &domain.Group{
		ID:          groupID,
		Name:        req.AdditionalProperties["Name"],
		Description: reqCtx.Description,
	}

	rows := impl.DB.Create(&group).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return group, nil
	}

	return &domain.Group{}, fmt.Errorf("failed to create group")
}
