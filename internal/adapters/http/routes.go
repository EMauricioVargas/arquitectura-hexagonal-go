package http

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes configura las rutas HTTP
func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.HelloHandler)
}
