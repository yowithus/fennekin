package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yowithus/fennekin/common"
	"github.com/yowithus/fennekin/controllers"
)

func main() {
	common.InitDB()

	router := gin.Default()
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:productId", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:productId", controllers.UpdateProduct)
	router.DELETE("/products/:productId", controllers.DeleteProduct)
	router.Run(":2205")
}
