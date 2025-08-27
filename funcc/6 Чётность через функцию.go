package funcc

import "fmt"

func cetnostcherezf() {
	fmt.Println(isEven(4))
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
