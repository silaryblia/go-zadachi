package cycles

import "fmt"

// Числа Фибоначчи

func fib() {
	var n int
	fmt.Scan(&n)
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}

}
