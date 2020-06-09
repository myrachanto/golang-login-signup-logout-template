package models

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	
	s "github.com/myrachanto/asearch/models"
)
func Getusers(c echo.Context) error {
	return c.String(200, "welcome to users page")
}

func Register(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err !=nil {
		return err
	}
	name := m["name"].(string)
	email := m["email"].(string)
	password := m["password"].(string)
	confrimPassword := m["confrimPassword"].(string)
	if password == "" || confrimPassword == "" || email == "" || name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "fill all the fields")
	}
	if password != confrimPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Your Passwords do not match")

	}
	if s.ValidateEmail(email) == false {
		return echo.NewHTTPError(http.StatusBadRequest, "check your email")
	}
	if UserExist(email) == true {
		return echo.NewHTTPError(http.StatusBadRequest, "This Use email Is Already taken")
	}
	config := s.GetConfig()
	//enc, err := s.EncryptionString(password, config.EncryptionKey)
	//password, err := bcrypt.GenerateFromPassword([]byte(auser.Password), config.EncryptionKey, bcrypt.DefaultCost)
	hash, err := s.HashPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	auser := s.User{Name:name, Email:email, Password:hash}
	s.GormDB.Create(&auser)
//	token := jwt.New(jwt.SigningMethodES256)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"name": auser.Name,
		"email": auser.Email,
	})
	t, err := token.SignedString([]byte(config.EncryptionKey))
	if err != nil {
		return err
	}
	auth := s.Auth{}
	if s.GormDB.First(&auth, "user_id =?", auser.ID).RecordNotFound(){
		//insert
		s.GormDB.Create(&s.Auth{User:auser, Token:t})
	} else {
		auth.User = auser
		auth.Token = t
		s.GormDB.Save(&auth)
	}
	//return c.JSON(http.StatusOk, map[string]string{
	return c.JSON(200, map[string]string{
		"token": t,
	})
}
func UserExist(email string) bool {
	auser := s.User{}
	if s.GormDB.First(&auser, "email =?", email).RecordNotFound(){
	   return false
	}
	return true
}

	
	/*	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

}*/
func Login(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err !=nil {
		return err
	}
	email := m["email"].(string)
	password := m["password"].(string)

	var auser s.User
	if s.GormDB.First(&auser, "email =?", email).RecordNotFound(){
		_error := s.CustomHTTPError{
			Error: s.ErrorType{
				Code: http.StatusBadRequest,
				Message: "Invalid email &password",
			},
		}
		return c.JSONPretty(http.StatusBadGateway, _error, " ")
	}
	config := s.GetConfig()
	//decrypted, _ := s.DecryptString(auser.Password, config.EncryptionKey)
		//decrypted, _ := s.DecryptString(auser.Password, config.EncryptionKey)
		hash := auser.Password
		//match := s.CheckPasswordHash(password, hash)
	   err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			_error := s.CustomHTTPError{
				Error: s.ErrorType{
					Code: 404,
					Message: "Invalid email &password",
				},
			}
			return c.JSONPretty(http.StatusBadGateway, _error, " ")
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
			"name": auser.Name,
			"email": auser.Email,
			"id": auser.ModelBase.ID,
		})
		t, err := token.SignedString([]byte(config.EncryptionKey))
		if err != nil {
			return err
		}
		
		auth := s.Auth{}
		if s.GormDB.First(&auth, "user_id =?", auser.ID).RecordNotFound(){
			s.GormDB.Create(&s.Auth{User: auser, Token: t})

		}else {
			auth.User = auser
			auth.Token = t
			s.GormDB.Save(&auth)
		}
		return c.JSON(200, map[string]string{
			"token":t,
		})
	}
func Logout(c echo.Context) error{
	tokenRequester := c.Get("user").(*jwt.Token)
	claims := tokenRequester.Claims.(jwt.MapClaims)
	fRequesterID := claims["id"].(float64)
	iRequesterID := int(fRequesterID)
	sRequesterID := strconv.Itoa(iRequesterID)

	requester := s.User{}
	if s.GormDB.First(&requester, "id =?", sRequesterID).RecordNotFound(){
		return echo.ErrUnauthorized
	}
	auth := s.Auth{}
	s.GormDB.Delete(&auth)
	return c.String(http.StatusAccepted, "")
}
/*func ValidateUser(email, password string, c echo.Context)(bool, error){
	fmt.Println("validate")
	var auser s.User
 	if s.GormDB.First(&auser, "email =?", email).RecordNotFound(){
		return false, nil
	}
	//config := s.GetConfig()
	//decrypted, _ := s.DecryptString(auser.Password, config.EncryptionKey)
	hash := auser.Password
	//match := s.CheckPasswordHash(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		_error := s.CustomHTTPError{
			Error: s.ErrorType{
				Code: 404,
				Message: "Invalid email &password",
			},
		}
		return c.JSONPretty(http.StatusBadGateway, _error, " ")
	}*/