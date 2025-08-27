package funcc

import "fmt"

func slojenievfunc() {
	result1 := sum(10, 5)
	result2 := sum(11, 6)
	result3 := sum(2, 4)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func sum(a int, b int) int {
	return a + b
}
