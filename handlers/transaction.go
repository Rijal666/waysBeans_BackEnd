package handlers

import (
	dto "backEnd/dto/result"
	transactionsdto "backEnd/dto/transaction"
	"backEnd/models"
	"backEnd/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
	UserRepository        repositories.UserRepository
	ProductRepository     repositories.ProductRepository
	CartRepository        repositories.CartRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository, UserRepository repositories.UserRepository, ProductRepository repositories.ProductRepository, CartRepository repositories.CartRepository) *handlerTransaction {
	return &handlerTransaction{
		TransactionRepository: TransactionRepository,
		UserRepository:        UserRepository,
		ProductRepository:     ProductRepository,
		CartRepository:        CartRepository,
	}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	if userAdmin {
		transactions, err := h.TransactionRepository.FindTransactions()
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}

		if len(transactions) > 0 {
			return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Data for all transactions was successfully obtained", Data: transactions})
		} else {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
		}
	} else {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
	}
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Transaction data successfully obtained", Data: transaction})
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactionsdto.TransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	userCart := user.Cart
	totalQuantity := 0
	for _, cart := range userCart {
		totalQuantity += cart.OrderQuantity
	}
	totalPrice := 0
	for _, cart := range userCart {
		product, err := h.ProductRepository.GetProduct(cart.ProductID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		multiplied := cart.OrderQuantity * product.Price
		totalPrice += multiplied
	}

	var userTransaction models.UserTransactionResponse
	userTransaction.ID = user.ID
	userTransaction.Name = user.Name
	userTransaction.Email = user.Email

	var productTransaction []models.ProductTransaction
	for _, cart := range userCart {
		product, err := h.ProductRepository.GetProduct(cart.ProductID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		var cartNew models.ProductTransaction
		cartNew.ProductID = product.ID
		cartNew.ProductName = product.Name
		cartNew.ProductPrice = product.Price
		cartNew.OrderQuantity = cart.OrderQuantity
		productTransaction = append(productTransaction, cartNew)
	}

	transaction := models.Transaction{
		UserID:             int(userId),
		User:               userTransaction,
		ProductTransaction: productTransaction,
		TotalQuantity:      totalQuantity,
		TotalPrice:         totalPrice,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	carts, err := h.CartRepository.FindCarts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	for _, cart := range carts {
		if cart.UserID == int(userId) {
			cartToDelete, err := h.CartRepository.GetCart(cart.ID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
			_, err = h.CartRepository.DeleteCart(cartToDelete)
			if err != nil {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
			}
		}
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Transaction data created successfully", Data: convertResponseTransaction(data)})
}

func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Message: "Cart data updated successfully", Data: convertResponseTransaction(data)})
}

func convertResponseTransaction(u models.Transaction) transactionsdto.TransactionResponse {
	return transactionsdto.TransactionResponse{
		UserID:        u.UserID,
		TotalQuantity: u.TotalQuantity,
		TotalPrice:    u.TotalPrice,
	}
}
