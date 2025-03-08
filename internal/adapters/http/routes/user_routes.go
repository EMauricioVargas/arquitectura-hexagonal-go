package routes

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/db"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http/controllers"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userRepo := db.NewUserRepository()
	userService := usecases.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := router.Group("/users")
	{
		api.GET("/:id", userController.GetUser)
		api.GET("", userController.GetUsers)
	}
}
