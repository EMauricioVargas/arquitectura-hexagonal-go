package controllers

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	ProductService usecases.ProductService
}

func NewProductController(service usecases.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	product, err := pc.ProductService.GetProductByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		return
	}
	c.JSON(http.StatusOK, products)
}
