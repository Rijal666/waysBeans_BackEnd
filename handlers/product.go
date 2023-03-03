package handlers

import (
	"backEnd/dto"
	"backEnd/dto/result"
	"backEnd/models"
	"backEnd/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
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
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	price, _:= strconv.Atoi(c.FormValue("price"))
	stock, _:= strconv.Atoi(c.FormValue("stock"))

	request := models.Product{
		Name: c.FormValue("name"),
		Description: c.FormValue("desciption"),
		Price: price,
		Stock: stock,	
		Photo: dataFile,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error(
			
		)})
	}
	product := models.Product{
		Name: request.Name,
		Price: request.Price,
		Description: request.Description,
		Stock: request.Stock,
		Photo: request.Photo,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
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

func (h *productHandler) DeleteProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err :=h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	data, err := h.ProductRepository.DeleteProduct(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
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
