package cycles

import "fmt"

// Удалить цифру из числа

func deletenumberr() {
	var number, delNumber int
	result, multiplier := 0, 1

	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	fmt.Println("Введите какую цифру удалить из числа: ")
	fmt.Scan(&delNumber)
	for number > 0 {
		digit := number % 10
		number /= 10
		if digit != delNumber {
			result += digit * multiplier
			multiplier *= 10
		}
	}
	fmt.Printf("Результат: %d\n", result)

}
