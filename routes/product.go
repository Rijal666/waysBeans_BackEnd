package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	ProductRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := handlers.HandlerProduct(ProductRepository)

	e.GET("/products", h.FindProducts)
	e.POST("/product", middleware.UploadFile(h.CreateProducts))
	e.GET("/products/:id", h.GetProducts)
	e.PATCH("/products/:id", h.UpdateProducts)
	e.DELETE("/products/:id", h.DeleteProducts)

}