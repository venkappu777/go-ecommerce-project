package handlers

import(
	"net/http"

	"github.com/labstack/echo/v4"

)

func Profile(c echo.Context)error{
	userID :=c.Get("user_id")
	email :=c.Get(("email"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userID,
		"email": email,
	})
}