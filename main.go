package main

import (
	"log"
	"os"
	"product-api/controller"
	"product-api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	db.SetupDatabaseConnection()

	r.POST("/product", controller.PostProduct)
	r.GET("/product", controller.GetProduct)
	r.GET("/product/:id", controller.GetProductByID)
	r.PUT("/product/:id", controller.UpdateProductByID)
	r.DELETE("/product/:id", controller.DeleteProductByID)

	log.Fatal(r.Run(":" + port))
}
