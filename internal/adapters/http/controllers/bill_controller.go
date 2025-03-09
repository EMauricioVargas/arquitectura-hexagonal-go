package controllers

import (
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/application/usecases"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BillController struct {
	BillService usecases.BillService
}

func NewBillController(service usecases.BillService) *BillController {
	return &BillController{BillService: service}
}

func (bc *BillController) GetBillById(c *gin.Context) {
	bill, err := bc.BillService.GetBillByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bill)
}

func (bc *BillController) GetBills(c *gin.Context) {
	bills, err := bc.BillService.GetBills()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bills)
}

func (bc *BillController) CreateBill(c *gin.Context) {
	var bill entities.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bc.BillService.CreateBill(&bill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, bill)
}
