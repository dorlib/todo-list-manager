package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
)

type GroupController struct {
	log         *zerolog.Logger
	validator   *validator.RequestValidator
	userService service.TaskService
}

func NewGroupController(log *zerolog.Logger, validator *validator.RequestValidator, GroupService service.TaskService) *GroupController {
	return &GroupController{
		log:         log,
		validator:   validator,
		userService: taskService,
	}
}

func (g *GroupController) CreateGroup(c echo.Context) error {
	req, err := request.BindAndValidateCreateGroupRequest(c, g.validator)
	if err != nil {
		return err
	}

	res, err := g.groupService.CreateOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (g *GroupController) GetGroups(c echo.Context) error {
	res, err := g.groupService.GetMany(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (g *GroupController) GetTask(c echo.Context) error {
	req, err := request.BindAndValidateGetGroupRequest(c)
	if err != nil {
		return err
	}

	res, err := g.groupService.GetOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (g *GroupController) DeleteGroup(c echo.Context) error {
	req, err := request.BindAndValidateDeleteGroupRequest(c)
	if err != nil {
		return err
	}

	err := g.groupService.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (g *GroupController) UpdateGroup(c echo.Context) error {
	req, err := request.BindAndValidateUpdategroupRequest(c)
	if err != nil {
		return err
	}

	res, err := g.groupService.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
