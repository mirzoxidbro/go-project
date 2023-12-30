package router

import (
	"go-project/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	{
		usersRouter := baseRouter.Group("/users")
		usersRouter.GET("", controller.UserControllerExecution().FindAll)
		usersRouter.GET("/:userId", controller.UserControllerExecution().FindById)
		usersRouter.POST("", controller.UserControllerExecution().Create)
		usersRouter.PUT("/:userId", controller.UserControllerExecution().Update)
		usersRouter.DELETE("/:userId", controller.UserControllerExecution().Delete)
	}

	return router
}
