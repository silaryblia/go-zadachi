package main

import "fmt"

type User struct {
	Name    string
	Balance float64
}

// Авто
type Car struct {
	Name     string
	Fuel     float64
	MaxFuel  float64
	FuelType string
	Owner    *User
}

// Станция