package middleware

import(
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{


		// step 1: Get Authorization header
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == ""{
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		// Step 2: Extract token (Bearer <token>)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Step 3: Parse & Validate token
		// this piece of code like it will parse tokenString into Token obejct
		// then after converting to token object it is passed as input to this function func(t *jwt.Token)
		// then based on secret it will try to check whether same signuature then valid token is true 
		token, err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("JWT_SECRET")),nil
		})
		 if err!=nil || !token.Valid{
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		 }

		 // Step 4: Extract claims
		 claims, ok:= token.Claims.(jwt.MapClaims)
		 if !ok{
			return c.JSON(http.StatusUnauthorized, "Invalid token claims")
		 }

		 // Step 5: Store user info in context
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		// Step 6: Continue to next handler
		return next(c)
	}
}

// Step we are following in this function

// Middleware intercepts the request
// Extracts JWT token from Authorization header
// Validates token using secret key
// If invalid → returns 401
// If valid → extracts claims
// Stores user data in context
// Calls next handler using next(c)