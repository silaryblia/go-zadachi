package funcc

import "fmt"

func obmen() {
	fmt.Println(swap(7, 9))
	// или
	a, b := swap(3, 5)
	fmt.Println(a, b)

}

func swap(a int, b int) (int, int) {
	return b, a
}
