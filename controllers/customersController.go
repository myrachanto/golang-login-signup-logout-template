package models

import (
	 "github.com/labstack/echo"
	// "fmt"
	// "net/http"
	// "strconv"
	// jwt "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	
	// s "github.com/myrachanto/asearch/models"
)
func GetCustomers(c echo.Context) error {
	return c.String(200, "welcome to users page")
}