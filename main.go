package main

import (
	"fmt"
	"viper/domain"
	"viper/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	martDelivery "viper/feature/Mart/delivery"
	martRepository "viper/feature/Mart/repository"
	martUsecase "viper/feature/Mart/usecase"

	productDelivery "viper/feature/Product/delivery"
	productRepository "viper/feature/Product/repository"
	productUsecase "viper/feature/Product/usecase"
)

var postgresDB *gorm.DB

func init() {

	utils.ViperInit()

}

func main() {
	e := echo.New()

	dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s",
		viper.GetString("postgresql.user"),
		viper.GetString("postgresql.password"),
		viper.GetString("postgresql.dbname"),
		viper.GetString("postgresql.port"),
	)

	var err error
	postgresDB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	group := e.Group("")

	martDelivery.NewMartHandler(group,
		martUsecase.NewMartUsecase(
			martRepository.NewMartRepository(postgresDB),
		))

	productDelivery.NewProductHandler(group,
		productUsecase.NewProductUsercase(
			productRepository.NewProductRepository(postgresDB),
		))

	// AutoMigrate()

	e.Start(":8000")
}

func AutoMigrate() {
	postgresDB.AutoMigrate(
		&domain.Mart{}, &domain.Product{},
	)
}
