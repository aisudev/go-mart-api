package delivery

import (
	"net/http"
	"strconv"
	"viper/domain"
	"viper/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.MartUsecase
}

func NewMartHandler(e *echo.Group, u domain.MartUsecase) *Handler {
	h := Handler{usecase: u}

	e.POST("/mart", h.CreateMartHandler)
	e.GET("/mart/:id", h.GetMarthandler)
	e.GET("/mart/:id/products", h.GetMartProductHandler)
	e.GET("/mart", h.GetAllMartHandler)
	e.PUT("/mart/:id", h.UpdateMartHandler)
	e.DELETE("/mart/:id", h.DeleteMartHandler)

	return &h
}

func (h *Handler) CreateMartHandler(c echo.Context) error {

	reqMap := domain.Mart{}

	if err := c.Bind(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	if err := h.usecase.CreateMart(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))
}

func (h *Handler) GetMarthandler(c echo.Context) error {

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	var mart *domain.Mart
	var err error

	if mart, err = h.usecase.GetMart(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, mart))
}

func (h *Handler) GetAllMartHandler(c echo.Context) error {

	var mart []domain.Mart
	var err error

	if mart, err = h.usecase.GetAllMart(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, mart))

}

func (h *Handler) UpdateMartHandler(c echo.Context) error {

	reqMap := domain.Mart{}

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	if err := c.Bind(&reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	if err := h.usecase.UpdateMart(uint(id), &reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))

}

func (h *Handler) DeleteMartHandler(c echo.Context) error {

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	if err := h.usecase.DeleteMart(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, nil))

}

func (h *Handler) GetMartProductHandler(c echo.Context) error {

	param := c.Param("id")
	id, _ := strconv.ParseUint(param, 10, 32)

	reqMap := []domain.MartProduct{}
	var err error

	if reqMap, err = h.usecase.GetMartProducts(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, err, nil))
	}

	return c.JSON(http.StatusOK, utils.Response(true, nil, reqMap))

}
