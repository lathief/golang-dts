package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GORMModel
	UserID      uint   `json:"user_id"`
	Title       string `json:"title" validate:"required-Title is required"`
	Description string `json:"description" validate:"required-Description is required"`
	User        *User  `json:"-"`
}

type ProductTest struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
