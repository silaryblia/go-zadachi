package cycles

import "fmt"

// Сумма и произведение цифр числа

func summaiproizvedenie() {
	var number int
	sum := 0
	mult := 1
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	for number > 0 {
		digit := number % 10
		sum += digit
		mult *= digit
		number /= 10
	}

	fmt.Printf("Сумма чисел равна: %d\n", sum)
	fmt.Printf("Произведение чисел равно: %d\n", mult)

}
