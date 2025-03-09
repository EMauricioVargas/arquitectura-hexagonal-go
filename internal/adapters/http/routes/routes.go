package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	RegisterProductRoutes(api)
	RegisterUserRoutes(api)
	RegisterBillRoutes(api)
}
