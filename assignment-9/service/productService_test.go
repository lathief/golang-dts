package service

import (
	// "assignment-9/database"
	"assignment-9/database"
	"assignment-9/repository"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepo}

func TestProductServiceGetOneProductFound(t *testing.T) {
	database.StartDB()
	db := database.GetDB()
	productRepo.Mock.On("FindByID", "2")

	product, err := productService.GetOneProduct(db, "2")

	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.Equal(t, uint(2), product.ID, "product id has to be 2")
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	database.StartDB()
	db := database.GetDB()
	productRepo.Mock.On("FindByID", "1")

	product, err := productService.GetOneProduct(db, "1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	database.StartDB()
	db := database.GetDB()
	productRepo.Mock.On("FindAll")

	product, err := productService.GetAllProduct(db)

	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	database.StartDB()
	db := database.GetDB()
	productRepo.Mock.On("FindAll")

	product, err := productService.GetAllProduct(db)

	assert.NotNil(t, err)
	assert.Nil(t, product)
}
