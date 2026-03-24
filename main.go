package main

import (
	"log"
	"os"
	"go-ecommerce/database"
	"go-ecommerce/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	err:=godotenv.Load()

	if err!=nil{
		log.Fatal("Error loadong .env file")
	}

	e := echo.New();// creating a server
	database.ConnectDB()

	e.GET("/", func(c echo.Context) error{
		return c.String(200, "Server running  🚀")
	})
	routes.SetupRoutes(e)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}