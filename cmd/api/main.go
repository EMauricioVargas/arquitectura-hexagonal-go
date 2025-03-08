package main

import (
	"fmt"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Registrar Handlers
	http.RegisterRoutes(r)

	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
