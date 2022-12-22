package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
}

func (Product) TableName() string {
	return "products"
}
