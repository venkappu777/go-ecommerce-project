package handlers

import (
	"context"
	"net/http"
	"go-ecommerce/database"
	"go-ecommerce/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error{
	var user models.User

	// Step1: Bind request body to struct
	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, "All fields are required")
	}

	// Step2: Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Error hashing password")
	}

	// Step3: Insert the data into DB
   	_,err = database.DB.Exec(context.Background(),
    "INSERT INTO users (name,email,password)VALUES($1, $2, $3)",
	user.Name, user.Email, string(hashedPassword),
   )
   if err!=nil{
	return c.JSON(http.StatusInternalServerError, err.Error())
   }

   // Step4: Return response 
   return c.JSON(http.StatusCreated, "User Created Successfully")
}