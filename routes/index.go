package routes

import (
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	c "github.com/myrachanto/asearch/controllers"
)
func Routes(){
	e := echo.New()
	//e.Use(middleware.Static("/client"))
	e.Static("/", "client")
	//e.Use(middleware.Logger())
	//e.Use(middleware.CORS())
	
	//echoGroupUseJWT := e.Group("/api/v1")
	//echoGroupUseJWT.Use(middleware.JWT([]byte(config.EncryptionKey)))
	//echoGroupNoJWT := e.Group("api/v1")
	//api v1/users: logged in usrs
	//////generall routes
	// e.File("/favicon.ico", "images/favicon.ico")
	// e.File("/", "public/index.html")

	e.POST("/users/logout", c.Logout)
	///api/v1/users :public
	e.GET("/users", c.Getusers)
	e.POST("/users/register", c.Register)
	e.POST("/users/login", c.Login)

	e.GET("/home", c.GetHome)
	////////customers///////////
	e.GET("/customers", c.GetCustomers)
	
}