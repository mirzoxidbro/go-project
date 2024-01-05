package router

import (
	"go-project/controller"
	middleware "go-project/middlewares"
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
		usersRouter.GET("/:userId", middleware.RoleMiddleware("admin"), controller.UserControllerExecution().FindById)
		usersRouter.GET("", controller.UserControllerExecution().FindAll)
		usersRouter.POST("", controller.UserControllerExecution().Create)
		usersRouter.PUT("/:userId", middleware.RoleMiddleware("admin"), controller.UserControllerExecution().Update)
		usersRouter.DELETE("/:userId", middleware.RoleMiddleware("admin"), controller.UserControllerExecution().Delete)
	}

	{
		authRouter := baseRouter.Group("/auth")
		authRouter.POST("/login", controller.Login)
		authRouter.Use(middleware.AuthMiddleware()).POST("/me", controller.CurrentUser)
	}

	{
		companyRouter := baseRouter.Group("/company")
		companyRouter.GET("/", middleware.RoleMiddleware("admin"), controller.CompanyControllerExecution().FindAll)
		companyRouter.GET("/:companyId", middleware.RoleMiddleware("admin,company"), controller.CompanyControllerExecution().FindById)
		companyRouter.POST("/", middleware.RoleMiddleware("admin,company"), controller.CompanyControllerExecution().Create)
		companyRouter.PUT("/:companyId", middleware.RoleMiddleware("admin,company"), controller.CompanyControllerExecution().Update)
		companyRouter.DELETE("/:companyId", middleware.RoleMiddleware("admin,company"), controller.CompanyControllerExecution().Delete)
	}
	return router
}
