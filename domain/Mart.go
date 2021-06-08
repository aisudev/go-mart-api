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

type MartUsecase interface {
	CreateMart(*Mart) error
	GetMart(uint) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(uint, *Mart) error
	DeleteMart(uint) error
}

type MartRepository interface {
	CreateMart(*Mart) error
	GetMart(uint) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(uint, *Mart) error
	DeleteMart(uint) error
}
