package internal

import (
	"github.com/gin-gonic/gin"
	jwt_service "sweete/core/pkg/jwt"
	password_service "sweete/core/pkg/password"
	"sweete/user_service/internal/delivery/http/v1/auth_controller"
	"sweete/user_service/internal/middlewares"
	user_repository "sweete/user_service/internal/repositories/user"
	auth_usecase "sweete/user_service/internal/usecases/auth"
)

func Routes(engine *gin.Engine) {
	userRepository := user_repository.New()

	jwtService := jwt_service.NewJWTService()
	passwordService := password_service.NewPasswordService()

	authUseCase := auth_usecase.New(userRepository, jwtService, passwordService)

	authController := auth_controller.New(authUseCase)

	routesWeb := engine.Group("")
	{
		routesWeb.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Service user",
			})
		})
	}

	// api
	routesApi := engine.Group("/api/v1").Use(middlewares.Recover)
	{
		routesApi.POST("/login", authController.Login)
		routesApi.POST("/register", authController.Register)
	}
}
