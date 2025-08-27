package cycles

import "fmt"

// Вычисление факториала числа

func factorial() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	result := 1
	for i := 1; i <= a; i++ {
		result *= i
	}

	fmt.Printf("Факториал числа: %d, равен: %d\n", a, result)
}
