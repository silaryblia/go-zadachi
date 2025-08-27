package funcc

import "fmt"

func funcvnytrifunc() {
	add := func(a, b int) int { return a + b }
	result := add(3, 5)
	fmt.Println(result)
}
