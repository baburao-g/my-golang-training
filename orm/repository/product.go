package repository

import (
	"orm/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	GetProduct(int) (model.Product, error)
	GetAllProduct() ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}

}

func (db *productRepository) GetProduct(id int) (product model.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) GetAllProduct() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) AddProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) DeleteProduct(product model.Product) (model.Product, error) {
	// if err := db.connection.First(&product, product.ID).Error; err != nil {
	// 	return product, err
	// }
	return product, db.connection.Delete(&product).Error
}

func (db *productRepository) UpdateProduct(product model.Product) (model.Product, error) {

	// if err := db.connection.First(&product, product.ID).Error; err != nil {
	// 	return product, err
	// }
	return product, db.connection.Model(&product).Updates(&product).Error
}
