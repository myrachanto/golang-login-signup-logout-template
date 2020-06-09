package models

import (

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