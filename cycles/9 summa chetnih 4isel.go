package cycles

import "fmt"

// Найти сумму четных цифр числа

func summachetnih() {
	var number int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&number)
	sum := 0
	for number > 0 {
		digit := number % 10
		if digit%2 == 0 {
			sum += digit
		}
		number /= 10
	}
	fmt.Printf("Сумма четных цифр числа равна: %d\n", sum)
}
