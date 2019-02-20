package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yowithus/fennekin/common"
	"github.com/yowithus/fennekin/models"
)

// GetProducts - get all products
func GetProducts(c *gin.Context) {
	db := common.GetDB()
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// GetProduct - get product by id
func GetProduct(c *gin.Context) {
	db := common.GetDB()
	var product models.Product
	productID := c.Param("productId")

	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct - create product
func CreateProduct(c *gin.Context) {
	db := common.GetDB()
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&product)
	c.JSON(http.StatusOK, product)
}

// UpdateProduct - udate product
func UpdateProduct(c *gin.Context) {
	db := common.GetDB()
	var productRequest models.Product
	productID := c.Param("productId")

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	product.Code = productRequest.Code
	product.Price = productRequest.Price
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct - soft delete product
func DeleteProduct(c *gin.Context) {
	db := common.GetDB()
	var product models.Product
	productID := c.Param("productId")

	if err := db.First(&product, productID).Delete(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"isDeleted": true})
}
