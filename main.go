package main

import (
	"log"
	"os"
	"product-api/controller"
	"product-api/db"
	"product-api/middleware"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
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

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "product-api",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/register", controller.Register)
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/product", controller.PostProduct)
		auth.GET("/product", controller.GetAllProducts)
		auth.GET("/product/:id", controller.GetProductByID)
		auth.PUT("/product/:id", controller.UpdateProductByID)
		auth.DELETE("/product/:id", controller.DeleteProductByID)
	}

	log.Fatal(r.Run(":" + port))
}
