package repository

import (
	"viper/domain"

	"gorm.io/gorm"
)

type martRepository struct {
	db *gorm.DB
}

func NewMartRepository(db *gorm.DB) domain.MartRepository {
	return &martRepository{
		db: db,
	}
}

func (r *martRepository) CreateMart(mart *domain.Mart) error {

	if err := r.db.Create(&mart).Error; err != nil {
		return err
	}

	return nil
}

func (r *martRepository) GetMart(id uint) (*domain.Mart, error) {

	var reqMap domain.Mart

	if err := r.db.Where("id = ?", id).Find(&reqMap).Error; err != nil {
		return nil, err
	}

	return &reqMap, nil
}

func (r *martRepository) GetAllMart() ([]domain.Mart, error) {

	var reqMap []domain.Mart

	if err := r.db.Find(&reqMap).Error; err != nil {
		return nil, err
	}

	return reqMap, nil
}

func (r *martRepository) UpdateMart(id uint, mart *domain.Mart) error {

	if err := r.db.Where("id = ?", id).Updates(&mart).Error; err != nil {
		return err
	}

	return nil
}

func (r *martRepository) DeleteMart(id uint) error {

	var mart domain.Mart

	if err := r.db.Where("id = ?", id).Delete(&mart).Error; err != nil {
		return err
	}

	return nil
}

func (r *martRepository) GetMartProducts(id uint) ([]domain.MartProduct, error) {

	reqMap := []domain.MartProduct{}

	if err := r.db.Debug().Table("marts").Select("products.name, products.amount, products.price").Where("marts.id = ?", id).Joins("JOIN products on marts.id = products.mart_id").Scan(&reqMap).Error; err != nil {
		return nil, err
	}

	return reqMap, nil
}
