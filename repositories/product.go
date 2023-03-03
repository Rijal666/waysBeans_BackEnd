package repositories

import (
	"backEnd/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	CreateProducts(product models.Product)(models.Product, error)
	GetProduct(ID int) (models.Product, error)
	UpdateProduct(product models.Product, ID int) (models.Product, error)


}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var Product []models.Product
	err := r.db.Find(&Product).Error
	return Product,err
}

func (r *repository) CreateProducts(product models.Product) (models.Product,error) {
	err := r.db.Create(&product).Error
	return product, err
}
func (r *repository) GetProduct(ID int) (models.Product, error) {
	var products models.Product
	err := r.db.First(&products, ID).Error

	return products, err
}
func (r *repository) UpdateProduct(product models.Product, ID int) (models.Product,error) {
	err := r.db.Save(&product).Error
	return product,err
}