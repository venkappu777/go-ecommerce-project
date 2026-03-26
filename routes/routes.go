package routes

import (
	"go-ecommerce/handlers"
	"go-ecommerce/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)
	// Protected Route example
	e.GET("/profile", handlers.Profile, middleware.JWTMiddleware)
}
