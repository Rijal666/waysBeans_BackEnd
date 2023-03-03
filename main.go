package main

import (
	"backEnd/database"
	"backEnd/pkg/mysql"
	"backEnd/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DataBaseInit()
	database.RunMigration()

	routes.Routes(e.Group("api/v1"))

	port := "5000"
	fmt.Println("server running on port", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}