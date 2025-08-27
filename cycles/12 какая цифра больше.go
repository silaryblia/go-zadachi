package cycles

import "fmt"

// Определить, какая цифра числа больше

func kakaiabolshe() {
	var number int
	fmt.Scan(&number)
	max := 0
	for number > 0 {
		digit := number % 10
		if digit > max {
			max = digit
		}
		number /= 10
	}
	fmt.Printf("Наибольшая цифра в числе: %d\n", max)
}
