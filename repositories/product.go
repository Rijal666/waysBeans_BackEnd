package repositories

import (
	"backEnd/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var Product []models.Product
	err := r.db.Find(&Product).Error
	return Product,err
}