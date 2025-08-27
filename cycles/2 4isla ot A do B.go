package cycles

import "fmt"

// Кубы чисел от A до B

func adob() {
	var a, b int
	a, b = 2, 10
	for i := a; i <= b; i++ {
		fmt.Print(i*i*i, " ")
	}

}
