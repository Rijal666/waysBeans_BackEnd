package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.ConnDB)
	profileRepository := repositories.RepositoryProfile(mysql.ConnDB)
	cartRepository := repositories.RepositoryCart(mysql.ConnDB)
	transactionRepository := repositories.RepositoryTransaction(mysql.ConnDB)
	h := handlers.HandlerUser(userRepository, profileRepository, cartRepository, transactionRepository)

	e.GET("/users", middleware.Auth(h.FindUsers))
	e.GET("/user/:id", middleware.Auth(h.GetUser))
	e.PATCH("/user", middleware.Auth(h.UpdateUser))
	e.DELETE("/user", middleware.Auth(h.DeleteUser))
}
