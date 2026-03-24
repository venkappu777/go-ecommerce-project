package main

import (
	"go-ecommerce/database"
	"go-ecommerce/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New();// creating a server
	database.ConnectDB()

	e.GET("/", func(c echo.Context) error{
		return c.String(200, "Server running  🚀")
	})
	routes.SetupRoutes(e)

	e.Start(":8080")
}