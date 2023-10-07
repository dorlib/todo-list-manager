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

type TaskController struct {
	log         *zerolog.Logger
	validator   *validator.RequestValidator
	taskService service.TaskService
}

func NewTaskController(log *zerolog.Logger, validator *validator.RequestValidator, taskService service.TaskService) *TaskController {
	return &TaskController{
		log:         log,
		validator:   validator,
		taskService: taskService,
	}
}

func (t *TaskController) CreateTask(c echo.Context) error {
	req, err := request.BindAndValidateCreateRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.CreateTask(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (t *TaskController) GetTasks(c echo.Context) error {
	_, err := request.BindAndValidateGetRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.GetTasks(c.Request().Context(), Ctx.FromEchoContext(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (t *TaskController) GetTask(c echo.Context) error {
	req, err := request.BindAndValidateGetRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.GetTask(c.Request().Context(), Ctx.FromEchoContext(c), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (t *TaskController) DeleteTask(c echo.Context) error {
	req, err := request.BindAndValidateDeleteRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.DeleteTask(c.Request().Context(), Ctx.FromEchoContext(c), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (t *TaskController) UpdateTask(c echo.Context) error {
	req, err := request.BindAndValidateUpdateRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.UpdateTask(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
