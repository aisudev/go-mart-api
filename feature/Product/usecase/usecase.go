package usecase

import "viper/domain"

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsercase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (u *productUsecase) CreateProduct(product *domain.Product) error {
	return u.repo.CreateProduct(product)
}

func (u *productUsecase) GetProduct(id uint) (*domain.Product, error) {
	return u.repo.GetProduct(id)
}

func (u *productUsecase) GetAllProduct() ([]domain.Product, error) {
	return u.repo.GetAllProduct()
}

func (u *productUsecase) UpdateProduct(id uint, product *domain.Product) error {
	return u.repo.UpdateProduct(id, product)
}

func (u *productUsecase) DeleteProduct(id uint) error {
	return u.repo.DeleteProduct(id)
}
