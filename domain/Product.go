package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID       uint           `json:"-" gorm:"primaryKey"`
	MartID   uint           `json:"mart_id"`
	Name     string         `json:"name" gorm:"varchar(50)"`
	Amount   int            `json:"amount"`
	Price    int            `json:"price"`
	CreateAt time.Time      `json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`
}

type ProductUsecase interface {
	CreateProduct(*Product) error
	GetProduct(uint) (*Product, error)
	GetAllProduct() ([]Product, error)
	UpdateProduct(uint, *Product) error
	DeleteProduct(uint) error
}

type ProductRepository interface {
	CreateProduct(*Product) error
	GetProduct(uint) (*Product, error)
	GetAllProduct() ([]Product, error)
	UpdateProduct(uint, *Product) error
	DeleteProduct(uint) error
}
