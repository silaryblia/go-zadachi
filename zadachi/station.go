package main

import "fmt"

// Машина
type Car struct {
	Model       string
	FuelLevel   int // уровень топлива
	MaxFuelSize int // максимальный уровень топлива
}

// Заправка
type GasStation struct{}

func (gs GasStation) Refuel(car *Car, amount int) {
	if car.FuelLevel >= car.MaxFuelSize {
		fmt.Printf("%s уже заправлен, бак полный\n", car.Model)
		return
	}

	need := car.MaxFuelSize - car.FuelLevel
	if amount > need {
		amount = need
	}

	car.FuelLevel += amount
	fmt.Printf("Заправляем %s на %d литров. Сейчас в баке: %d/%d\n",
		car.Model, amount, car.FuelLevel, car.MaxFuelSize)

	if car.FuelLevel == car.MaxFuelSize {
		fmt.Printf("%s полностью заправлен\n", car.Model)
	}
}

func main() {
	car := Car{Model: "BMW", FuelLevel: 15, MaxFuelSize: 40}

	station := GasStation{}

	station.Refuel(&car, 15)
	station.Refuel(&car, 30)
	station.Refuel(&car, 10)
}
