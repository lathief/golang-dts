package repository

import (
	"assignment-9/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(db *gorm.DB, id string) *models.Product
	FindAll(db *gorm.DB) *[]models.Product
}
