package main

import (
	"github.com/labstack/echo"
	"os"
	 "github.com/joho/godotenv"
	"net/http"
	"log"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	//r "github.com/myrachanto/asearch/routes"
	s "github.com/myrachanto/asearch/support"
	c "github.com/myrachanto/asearch/controllers"
)
var counter int
type (
	Currency float64
	Item struct {
		Name string `json:"name,omitempty"`
		Quantity int `json:"quantity,omitempty"`
		Price Currency `json:"price,omitempty"`
	}
	Store struct {
		Items map[string]Item `json:"items,omitempty"`
	}
	Dashboard struct {
		Users uint `json:"users,omitempty"`
		UsersLoggedIn uint `json:"users_logged_in,omitempty"`
		Inventory *Store `json:"inventory,omitempty"`
	}
)
var dashboard chan *Dashboard
func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	
	PORT := os.Getenv("PORT")
	//http.Handle("/", http.FileServer(http.Dir("client")))
	s.Configs()
	//r.Routes()
	e := echo.New()
	//e.Use(middleware.Static("/client"))
	e.Static("/", "client/public")
	//e.Use(middleware.Logger())
	//e.Use(middleware.CORS())
	
	//echoGroupUseJWT := e.Group("/api/v1")
	//echoGroupUseJWT.Use(middleware.JWT([]byte(config.EncryptionKey)))
	//echoGroupNoJWT := e.Group("api/v1")
	//api v1/users: logged in usrs
	//////generall routes
	e.File("/favicon.ico", "client/public/favicon.ico")
	// e.File("/", "public/index.html")

	e.POST("/users/logout", c.Logout)
	///api/v1/users :public
	e.GET("/users", c.Getusers)
	e.POST("/users/register", c.Register)
	e.POST("/users/login", c.Login)
	
	dashboard = make(chan *Dashboard)
	go UpdateDashboard()
	http.HandleFunc("/dashboard", DashboardHandler)

	e.GET("/home", c.GetHome)
	////////customers///////////
	e.GET("/customers", c.GetCustomers)
	e.Logger.Fatal(e.Start(PORT))
}
//convert to echo on the todo list
func DashboardHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	counter++
	//fmt.Fprintf(w, "data: %v\n\n", counter)
	//fmt.Printf("data: %v\n", counter)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(<- dashboard)
	fmt.Fprintf(w, "data: %v\n\n", buf.String() )
	fmt.Printf("data: %v\n", buf)
}
func UpdateDashboard(){
	for {
		inv := UpdateInvetory()
		db := &Dashboard{
			Users: uint(rand.Uint32()),
			UsersLoggedIn: uint(rand.Uint32() % 200),
			Inventory: inv,
		//db := &dashboard{}}
		}
		dashboard <- db
	}
	
}
func UpdateInvetory() *Store{
	inv:= &Store{}
	inv.Items = make(map[string]Item)
	a:= Item{Name:"Books", Price:500, Quantity:int(rand.Int31()%50)}
	inv.Items["books"] = a
	a= Item{Name:"car", Price:1500000, Quantity:int(rand.Int31()%70)}
	inv.Items["car"] = a
	a= Item{Name:"jacket", Price:1000, Quantity:int(rand.Int31()%56)}
	inv.Items["jacket"] = a
	a= Item{Name:"Tshirt", Price:500, Quantity:int(rand.Int31()%59)}
	inv.Items["tshirt"] = a
	a= Item{Name:"Shoe", Price:1000, Quantity:int(rand.Int31()%100)}
	inv.Items["shoe"] = a
	return inv
}