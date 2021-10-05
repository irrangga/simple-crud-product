package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	Price       uint
	Stock       uint
	Category    string
}
