package models
import (
	// "github.com/revel/revel"
)


type Application struct {
	IdApplication         int
	NomApplication, Address    string
	City, State, Zip string
	Country          string
	Price            int
}
