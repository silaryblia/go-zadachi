package massiv

import "fmt"

// Поиск максимального элемента
// Найди наибольшее значение в массиве из 10 чисел.
// **Цель:** Понять сравнение значений в цикле.

func poiskmax() {
	// Ввод: [1, 4, 6, 8, 3, 9, 2, 0, 5, 7]
	// Вывод: 9
	mass := [10]int{1, 4, 6, 8, 3, 9, 2, 0, 5, 7}
	fmt.Println(mass)
	max := mass[0]
	for i := 0; i < 10; i++ {
		if mass[i] > max {
			max = mass[i]
		}
	}
	fmt.Println(max)
}
