package controllers

import (
	"jwtapi-product/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateTransactionInput struct {


	User_Id uint `json:"user_id" binding:"required"`
	Product_Id uint `json:"product_id" binding:"required"`
	Amount uint `json:"amount" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateTransactionInput struct {
	ID  uint `json:"id"`
	User_Id  uint `json:"user_id"`
	Product_Id uint `json:"product_id"`
	Amount uint `json:"amount"`
	Status string `json:"status"`
}

func CreateTransaction(c *gin.Context) {
	// Validate input
	var input CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	transaction := models.Transaction{User_Id: input.User_Id, Product_Id: input.Product_Id, Amount: input.Amount, Status:input.Status}
	models.DB.Create(&transaction)

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// PATCH /TRANSACTIONs/:id
// Update a TRANSACTION
func UpdateTransaction(c *gin.Context) {
	// Get model if exist
	var input UpdateTransactionInput
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", input.ID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	
	models.DB.Model(&transaction).Updates(models.Transaction{Status: input.Status, Amount: input.Amount})

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

