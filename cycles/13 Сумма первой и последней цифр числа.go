package cycles

import "fmt"

// Сумма первой и последней цифр числа

func summacirfch() {
	var number, first, last int
	fmt.Scan(&number)
	last = number % 10
	for number > 0 {
		digit := number % 10
		number /= 10
		first = digit
	}
	fmt.Printf("Первая цифра числа: %d\n", first)
	fmt.Printf("Последняя цифра числа: %d\n", last)
	fmt.Printf("Сумма первой и посдений цифры: %d\n", first+last)
}
