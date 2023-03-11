package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.ConnDB)
	userRepository := repositories.RepositoryUser(mysql.ConnDB)
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	cartRepository := repositories.RepositoryCart(mysql.ConnDB)
	h := handlers.HandlerTransaction(transactionRepository, userRepository, productRepository, cartRepository)

	e.GET("/transactions", middleware.Auth(h.FindTransactions))
	e.GET("/transaction/:id", middleware.Auth(h.GetTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
}
