package funcc

import "fmt"

func recurs() {
	countdown(5)
}

func countdown(n int) {
	if n <= 0 {
		return
	}
	fmt.Print(n, " ")
	countdown(n - 1)
}
