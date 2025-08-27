package funcc

import "fmt"

var name string = "Molly"

func changeName() {
	name = "Joshua"
}

func proverka() {
	fmt.Println("Было:", name) // Molly
	changeName()
	fmt.Println("Стало:", name) // Joshua
}
