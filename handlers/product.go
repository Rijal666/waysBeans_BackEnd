package handlers

import (
	"backEnd/dto/result"
	"backEnd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *productHandler{
	return &productHandler{ProductRepository}
}

func (h *productHandler) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: products})
}
