package routes

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/db"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http/controllers"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/gin-gonic/gin"
)

func RegisterBillRoutes(router *gin.RouterGroup) {
	billRepo := db.NewBillRepository()
	productRepo := db.NewProductRepository()
	billService := usecases.NewBillService(billRepo, productRepo)
	billController := controllers.NewBillController(billService)

	api := router.Group("/bills")
	{
		api.GET("/:id", billController.GetBillById)
		api.GET("", billController.GetBills)
		api.POST("", billController.CreateBill)
	}
}
