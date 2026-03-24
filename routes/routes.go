package routes

import(
	"go-ecommerce/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo){
	e.POST("/signup", handlers.Signup)
	e.POST("/login",handlers.Login)
}