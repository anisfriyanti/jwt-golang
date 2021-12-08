package controllers

import (
	"jwtapi-product/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
	Qty string `json:"qty" binding:"required"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Qty string `json:"qty"`
}

// GET /products
// Find all products
func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GET /products/:id
// Find a product
func FindProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// POST /products
// Create new product
func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	product := models.Product{Name: input.Name, Price: input.Price, Qty: input.Qty}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /products/:id
// Update a product
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//models.DB.Model(&product).Updates(input)
	models.DB.Model(&product).Updates(models.Product{Name: input.Name, Price: input.Price, Qty: input.Qty})

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /products/:id
// Delete a product
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
