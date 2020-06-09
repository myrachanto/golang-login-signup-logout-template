package support

import (
	"github.com/jinzhu/gorm"
	// "os"
	// "github.com/joho/godotenv"
	// "log"
	// "fmt"
	m "github.com/myrachanto/asearch/models"

	_"github.com/jinzhu/gorm/dialects/mysql"
	// _"github.com/jinzhu/gorm/dialects/postgres"
	// _"github.com/jinzhu/gorm/dialects/sqlite"
	//_"github.com/jinzhu/gorm/dialects/mssql"
)
func Configs(){
	//config := m.GetConfig()
	//gormParams := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbName, config.DbUsername, config.DbPassword)
	//gormParams := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbName, config.DbUsername, config.DbPassword)
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// fmt.Println(err)
	
	// PORT := os.Getenv("PORT")
	//DbName := os.Getenv("DbName")
	//DbUsername := os.Getenv("DbUsername")
	//, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	dbURI := "root@/search?charset=utf8&parseTime=True&loc=Local"
	//GormDB, err := gorm.Open("mysql", gormParams)
	GormDB, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic("failed to connect database")
	}
	
	defer GormDB.Close()
	m.GormDB = GormDB
	//migrate tables
	m.GormDB.AutoMigrate(&m.User{})
	m.GormDB.AutoMigrate(&m.Auth{})
	m.GormDB.AutoMigrate(&m.Customer{})
	m.GormDB.AutoMigrate(&m.Invoice{})
	m.GormDB.AutoMigrate(&m.InvoiceItem{})
	///////////////////////////////////////////////////////////////
	//rememdber to include soft delete
	//////////////////////////////////////////////////////////////////////////\
	m.GormDB.Model(&m.Auth{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}