package main

import (
	"fmt"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/adapters/http/routes"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	config.LoadConfig()
	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
