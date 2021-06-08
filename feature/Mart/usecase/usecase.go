package usecase

import (
	"viper/domain"
)

type martUsecase struct {
	repo domain.MartRepository
}

func NewMartUsecase(repo domain.MartRepository) domain.MartUsecase {
	return &martUsecase{
		repo: repo,
	}
}

func (u *martUsecase) CreateMart(mart *domain.Mart) error {
	return u.repo.CreateMart(mart)
}

func (u *martUsecase) GetMart(id uint) (*domain.Mart, error) {
	return u.repo.GetMart(id)
}

func (u *martUsecase) GetAllMart() ([]domain.Mart, error) {
	return u.repo.GetAllMart()
}

func (u *martUsecase) UpdateMart(id uint, mart *domain.Mart) error {
	return u.repo.UpdateMart(id, mart)
}

func (u *martUsecase) DeleteMart(id uint) error {
	return u.repo.DeleteMart(id)
}
