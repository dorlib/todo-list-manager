package router

import "github.com/labstack/echo/v4"

func routeReady(e *echo.Echo, apiPrefix string, readyController *controller.readyController) {
	subRouter := e.Group(apiPrefix + "/ready")
	subRouter.GET("", readyController.Ready)
}

func routerGroups(router echo.Router, groupController *controller.groupController) {
	subRouter := router.Group("/groups")
	subRouter.POST("", groupController.CreateGroups)
	subRouter.GET("", groupController.GetGroups)
	subRouter.GET("/:"+request.IDParam, groupController.GetGroup)
	subRouter.DELETE("/:"+request.IDParam, groupController.DeleteGroup)
	subRouter.PATCH("/:"+request.IDParam, groupController.UpdateGroup)

}

func routeUsers(router *echo.Group, userController *controller.userController) {
	subRouter := router.Group("/users")
	subRouter.POST("", userController.CreateUser)
	subRouter.GET("", userController.GetUsers)
	subRouter.GET("/:"+request.IDParam, userController.GetUser)
	subRouter.DELETE("/:"+requesrt.IDParam, userController.DeleteUser)
	subRouter.PATCH("/:"+request.IDParamm, userController.UpdateUser)
}

func routeTasks(router *echo.Group, taskController *controller.taskController) {
	subRouter := router.Group("/tasks")
	subRouter.POST("", taskController.CreateTask)
	subRouter.GET("", taskController.GetTasks)
	subRouter.GET("/:"+request.IDParam, taskController.GetTask)
	subRouter.DELETE("/:"+requesrt.IDParam, taskController.DeleteTask)
	subRouter.PATCH("/:"+request.IDParamm, taskController.UpdateTask)
}
