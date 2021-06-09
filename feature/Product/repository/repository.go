package repository

import (
	"viper/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product *domain.Product) error {

	if err := r.db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetProduct(id uint) (*domain.Product, error) {

	reqMap := domain.Product{}

	if err := r.db.Where("id = ?", id).Find(&reqMap).Error; err != nil {
		return nil, err
	}

	return &reqMap, nil
}

func (r *productRepository) GetAllProduct() ([]domain.Product, error) {
	var products []domain.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) UpdateProduct(id uint, product *domain.Product) error {

	if err := r.db.Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(id uint) error {

	var product domain.Product

	if err := r.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
