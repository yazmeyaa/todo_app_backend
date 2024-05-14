package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yazmeyaa/todo_app_backend/middlewares"
	"github.com/yazmeyaa/todo_app_backend/services"
)

type RouterControllers struct {
	Tasks TaskController
	Users UserController
	Auth  AuthController
}

func NewRouter(controllers *RouterControllers, validator *validator.Validate, jwtService services.JWTService) *gin.Engine {
	router := gin.Default()

	taskRouter := router.Group("/task")
	taskRouter.Use(middlewares.AuthJWTMiddleware(jwtService))
	taskRouter.GET("/", controllers.Tasks.GetList)
	taskRouter.POST("/", controllers.Tasks.Create)
	taskRouter.GET("/:id", controllers.Tasks.GetById)
	taskRouter.PATCH("/:id", controllers.Tasks.Update)
	taskRouter.DELETE("/:id", controllers.Tasks.Delete)

	// userRouter := router.Group("/user")
	// userRouter.POST("/", controllers.Users.Create)

	authRouter := router.Group("/auth")
	authRouter.POST("/login", controllers.Auth.Login)
	authRouter.POST("/register", controllers.Auth.Register)
	authRouter.POST("/check", middlewares.AuthJWTMiddleware(jwtService), controllers.Auth.CheckToken)

	return router
}
