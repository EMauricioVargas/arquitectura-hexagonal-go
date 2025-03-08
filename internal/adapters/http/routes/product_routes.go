package routes

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/db"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http/controllers"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup) {
	productRepo := db.NewProductRepository()
	productService := usecases.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	api := router.Group("/products")
	{
		api.GET("/:id", productController.GetProduct)
		api.GET("", productController.GetProducts)
	}
}
