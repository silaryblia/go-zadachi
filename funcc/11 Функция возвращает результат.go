package funcc

import "fmt"

func vozvratresult() {
	result1 := square(5) + square(4)
	result2 := square(6) + square(5)
	result3 := square(7) + square(6)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func square(n int) int {
	return n * n
}
