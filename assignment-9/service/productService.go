package service

import (
	"assignment-9/models"
	"assignment-9/repository"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (s ProductService) GetOneProduct(db *gorm.DB, id string) (*models.Product, error) {
	fmt.Fprintln(os.Stdout, db)
	product := s.Repository.FindByID(db, id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s ProductService) GetAllProduct(db *gorm.DB) (*[]models.Product, error) {
	fmt.Fprintln(os.Stdout, db)
	product := s.Repository.FindAll(db)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
