package handlers

import (
	"net/http"
	"go-ecommerce/database"
	"go-ecommerce/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"go-ecommerce/utils"
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
   	_,err = database.DB.Exec(c.Request().Context(),
    "INSERT INTO users (name,email,password)VALUES($1, $2, $3)",
	user.Name, user.Email, string(hashedPassword),
   )
   if err!=nil{
	return c.JSON(http.StatusInternalServerError, err.Error())
   }

   // Step4: Return response 
   return c.JSON(http.StatusCreated, "User Created Successfully")
}

func Login(c echo.Context) error{
	var input models.User;


	// step 1: Bind request
	if err := c.Bind(&input); err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}
    
	// Validate input 
	if input.Email=="" || input.Password==""{
		return c.JSON(http.StatusBadRequest, "Email and password required")
	}

	// Step 3: Get user from db
	var user models.User

	err := database.DB.QueryRow(c.Request().Context(),
		"SELECT id, name, email, password FROM users WHERE email=$1",
		input.Email,).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		// fill the data row return from database to struct of users using Scan

	if err != nil{
		return	c.JSON(http.StatusUnauthorized, "Invalid email or password") 
	}

	// Step 4: Compare Password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil{
		return c.JSON(http.StatusUnauthorized, "Invalid password")
	}

	// Step 5: Generate JWT token
	tokenString,err := utils.GenerateToken(user.ID, user.Email)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "Could generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})

}