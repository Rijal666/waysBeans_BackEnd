package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
	e.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))
	e.PATCH("/product/:id", middleware.Auth(middleware.UploadFile(h.UpdateProduct)))
}
