package controllers

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService usecases.UserService
}

func NewUserController(service usecases.UserService) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) GetUser(c *gin.Context) {
	user, err := uc.UserService.GetUserByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}
