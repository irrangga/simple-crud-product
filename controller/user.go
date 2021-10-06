package controller

import (
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
	}
	db.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
