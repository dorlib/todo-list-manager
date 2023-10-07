package controller

import (
	Ctx "github.com/dorlib/todo-list-manager/libs/microservice-development-kit/ctx"
	"github.com/dorlib/todo-list-manager/libs/microservice-development-kit/validator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"tasks-mgmnt/model/request"
	"tasks-mgmnt/service"
)

type GroupController struct {
	log          *zerolog.Logger
	validator    *validator.RequestValidator
	groupService service.GroupService
}

func NewGroupController(log *zerolog.Logger, validator *validator.RequestValidator, GroupService service.GroupService) *GroupController {
	return &GroupController{
		log:          log,
		validator:    validator,
		groupService: GroupService,
	}
}

func (g *GroupController) CreateGroup(c echo.Context) error {
	req, err := request.BindAndValidateCreateRequest(c, g.validator)
	if err != nil {
		return err
	}

	res, err := g.groupService.CreateGroup(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (g *GroupController) GetGroups(c echo.Context) error {
	_, err := request.BindAndValidateGetRequest(c, g.validator)
	if err != nil {
		return err
	}

	res, err := g.groupService.GetGroups(c.Request().Context(), Ctx.FromEchoContext(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (g *GroupController) DeleteGroup(c echo.Context) error {
	req, err := request.BindAndValidateDeleteRequest(c, g.validator)
	if err != nil {
		return err
	}

	res, err := g.groupService.DeleteGroup(c.Request().Context(), Ctx.FromEchoContext(c), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (g *GroupController) UpdateGroup(c echo.Context) error {
	req, err := request.BindAndValidateUpdateRequest(c, g.validator)
	if err != nil {
		return err
	}

	res, err := g.groupService.UpdateGroup(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
