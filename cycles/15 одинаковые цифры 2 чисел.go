package cycles

import "fmt"

// Найти одинаковые цифры двух чисел

func odinikov() {
	var a, b int

	fmt.Println("Введите два числа: ")
	fmt.Scan(&a, &b)

	fmt.Print("Общие цифры: ")

	for digit := 0; digit <= 9; digit++ {
		foundInA := false
		foundInB := false

		for temp := a; temp > 0; temp /= 10 {
			if temp%10 == digit {
				foundInA = true
				break
			}
		}

		for temp := b; temp > 0; temp /= 10 {
			if temp%10 == digit {
				foundInB = true
				break
			}
		}

		if foundInA && foundInB {
			fmt.Print(digit, " ")
		}
	}
	fmt.Println()
}
