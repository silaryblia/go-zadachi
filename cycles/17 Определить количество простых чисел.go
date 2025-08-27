package cycles

import "fmt"

// Определить количество простых чисел

// C клавиатуры вводятся целые числа до первого числа, которое меньше двух.
//  Написать программу, которая определяет сколько простых чисел было введено.

func natur() {
	count := 0
	var num int
	fmt.Println("Введите числа: ")
	for {
		fmt.Scan(&num)
		if num < 2 {
			break
		}

		isPrime := true
		// Т.к. делитель не может быть больше половины числа
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime && num >= 2 {
			count++
		}
	}
	fmt.Println("Найдено простых чисел:", count)
}
