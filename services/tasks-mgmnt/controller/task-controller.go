package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
)

type TaskController struct {
	log         *zerolog.Logger
	validator   *validator.RequestValidator
	userService service.TaskService
}

func NewTaskController(log *zerolog.Logger, validator *validator.RequestValidator, TaskService service.TaskService) *TaskController {
	return &TaskController{
		log:         log,
		validator:   validator,
		userService: taskService,
	}
}

func (t *TaskController) CreateTask(c echo.Context) error {
	req, err := request.BindAndValidateCreateTaskRequest(c, t.validator)
	if err != nil {
		return err
	}

	res, err := t.taskService.CreateOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (t *TaskController) GetTasks(c echo.Context) error {
	res, err := t.taskService.GetMany(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (t *TaskController) GetTask(c echo.Context) error {
	req, err := request.BindAndValidateGetTaskRequest(c)
	if err != nil {
		return err
	}

	res, err := t.taskService.GetOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (t *TaskController) DeleteTask(c echo.Context) error {
	req, err := request.BindAndValidateDeleteTaskRequest(c)
	if err != nil {
		return err
	}

	err := t.taskService.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (t *TaskController) UpdateTask(c echo.Context) error {
	req, err := request.BindAndValidateUpdateTaskRequest(c)
	if err != nil {
		return err
	}

	res, err := t.taskService.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
