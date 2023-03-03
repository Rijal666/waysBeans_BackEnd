package handlers

import (
	"backEnd/dto"
	"backEnd/dto/result"
	"backEnd/models"
	"backEnd/repositories"
	"net/http"
	"strconv"

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

func (h *productHandler) CreateProducts(c echo.Context) error {
	request := new(dto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	product := models.Product{
		Name: request.Name,
		Price: request.Price,
		Description: request.Desciption,
		Stock: request.Stock,	
		Photo: request.Photo,
	}
	data, err := h.ProductRepository.CreateProducts(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
}

func (h *productHandler) GetProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	products, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(products)})
}

func (h *productHandler) UpdateProducts(c echo.Context) error {
	request := new(dto.UpdateProductRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest,result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	id, _:= strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	} 
		if request.Name != ""{
		product.Name = request.Name
	} 
		if request.Price != 0 {
		product.Price = request.Price
	} 
		if request.Description != "" {
		product.Description = request.Description
	} 	
		if request.Stock != 0 {
		product.Stock = request.Stock
	} 
		if request.Photo != "" {
		product.Photo = request.Photo
	} 
		data, err := h.ProductRepository.UpdateProduct(product, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		}
		return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProduct(data)})
}

func convProduct (u models.Product) dto.ProductResponse{
	return dto.ProductResponse{
		ID: u.ID,
		Name: u.Name,
		Price: u.Price,
		Description: u.Description,
		Stock: u.Stock,
		Photo: u.Photo,
	}
}
