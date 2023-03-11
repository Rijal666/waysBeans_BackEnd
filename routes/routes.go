package routes

import "github.com/labstack/echo/v4"

func Routes(e *echo.Group) {
	UserRoutes(e)
	ProfileRoutes(e)
	ProductRoutes(e)
	AuthRoutes(e)
	CartRoutes(e)
	TransactionRoutes(e)
}
