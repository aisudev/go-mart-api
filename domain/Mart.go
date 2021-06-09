package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mart struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name" gorm:"varchar(50)"`
	CreateAt time.Time      `json:"-" gorm:"autoCreateTime"`
	DeleteAt gorm.DeletedAt `json:"-"`

	Products []Product `json:"-" gorm:"foreignKey:MartID"`
}

type MartProduct struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

type MartUsecase interface {
	CreateMart(*Mart) error
	GetMart(uint) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(uint, *Mart) error
	DeleteMart(uint) error
	GetMartProducts(uint) ([]MartProduct, error)
}

type MartRepository interface {
	CreateMart(*Mart) error
	GetMart(uint) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(uint, *Mart) error
	DeleteMart(uint) error
	GetMartProducts(uint) ([]MartProduct, error)
}
