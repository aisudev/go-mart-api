package delivery

import (
	"net/http"
	"strconv"
	"viper/domain"
	"viper/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(e *echo.Group, u domain.ProductUsecase) *Handler {
	h := Handler{usecase: u}

	e.POST("/product", h.CreateProductHandler)
	e.GET("/product/:id", h.GetProducthandler)
	e.GET("/product", h.GetAllProductHandler)
	e.PUT("/product/:id", h.UpdateProductHandler)
	e.DELETE("/product/:id", h.DeleteProductHandler)

	return &h
}

func (h *Handler) CreateProductHandler(c echo.Context) error {

	reqMap := domain.Product{}

	if err := c.Bind(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	if err := h.usecase.CreateProduct(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))
}

func (h *Handler) GetProducthandler(c echo.Context) error {

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	var product *domain.Product
	var err error

	if product, err = h.usecase.GetProduct(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, product))
}

func (h *Handler) GetAllProductHandler(c echo.Context) error {

	var product []domain.Product
	var err error

	if product, err = h.usecase.GetAllProduct(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, product))

}

func (h *Handler) UpdateProductHandler(c echo.Context) error {

	reqMap := domain.Product{}

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	if err := c.Bind(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	if err := h.usecase.UpdateProduct(uint(id), &reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))

}

func (h *Handler) DeleteProductHandler(c echo.Context) error {

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	if err := h.usecase.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))

}
