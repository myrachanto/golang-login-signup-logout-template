package models

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"

	_"encoding/json"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"log"
	"fmt"
)
var config Config
var GormDB *gorm.DB
type Config struct {
		EncryptionKey string
		DbHost string
		DbPort string
		DbName string
		DbUsername string
		DbPassword string
	}
type ModelBase struct {
		ID uint64 `gorm:"primary_key"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	//	DeletedAt *time.Time `json:"deleted_at"`
	}

type CustomHTTPSuccess struct {
	Data string `json:"data"`
}
type CustomHTTPError struct {
	Error ErrorType `json:"error"`
}
type ErrorType struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
//todo transfer to env file later
//env file call

func init() {
	//call from env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(err)
	
	//port := os.Getenv("PORT")
	EncryptionKey := os.Getenv("EncryptionKey")
	DbHost := os.Getenv("DbHost")
	DbName := os.Getenv("DbName")
	DbPort := os.Getenv("DbPort")
	DbUsername := os.Getenv("DbUsername")
	DbPassword := os.Getenv("DbPassword")
	config = Config{
		EncryptionKey,
		DbHost,
		DbPort,
		DbName,
		DbUsername,
		DbPassword,
	}
	/*
	config = Config{
		EncryptionKey: "Myrachanto",
		DbHost: "localhost",
		DbPort: "5432",
		DbName: "echoauth",
		DbUsername: "root",
		DbPassword: "",
	}*/
}
func GetConfig() Config{
	return config
}
func ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return
}
func ValidatePassword(password string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(password)
	return
}
func HashPassword(password string)(string, error){
/////////////////////////////////////////////
//hash, _ := HashPassword(password)
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(pwd), err
	
}