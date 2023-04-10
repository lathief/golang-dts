package repository

import (
	"assignment-9/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) FindByID(db *gorm.DB, id string) *models.Product {
	var product models.Product

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil
	} else {
		return &product
	}
}

func (r *ProductRepositoryMock) FindAll(db *gorm.DB) *[]models.Product {
	var product []models.Product

	if err := db.Find(&product).Error; err != nil {
		return nil
	} else {
		return &product
	}
}
