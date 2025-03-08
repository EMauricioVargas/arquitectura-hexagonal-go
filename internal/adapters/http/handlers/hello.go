package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloHandler responde con un mensaje JSON
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, Hexagonal Architecture!"})
}
