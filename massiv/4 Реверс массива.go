package massiv

import "fmt"

// Реверс массива
//Выведи массив в обратном порядке.
//**Цель:** Научиться индексировать массив "с конца".

func reverse() {
	// Ввод: [1, 2, 3, 4, 5]
	// Вывод: [5, 4, 3, 2, 1]
	mass := []int{1, 2, 3, 4, 5}
	fmt.Println(mass)

	fmt.Print("reverse: [")
	for i := len(mass) - 1; i >= 0; i-- {
		fmt.Print(mass[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
