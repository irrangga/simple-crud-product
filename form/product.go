package form

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Stock       uint   `json:"stock" binding:"required"`
	Category    string `json:"category" binding:"required"`
}
