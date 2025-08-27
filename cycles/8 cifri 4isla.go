package cycles

import (
	"fmt"
	"strconv"
)

// Извлечение цифр числа

func izvlechenie() {
	var number int
	fmt.Println("Введите целое число:")
	fmt.Scan(&number)

	numberStr := strconv.Itoa(number)

	fmt.Printf("Число %d состоит из цифр: ", number)

	for _, digitChar := range numberStr {
		digit, _ := strconv.Atoi(string(digitChar))
		fmt.Print(digit, " ")
	}
}
