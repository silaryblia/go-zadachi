package massiv

import "fmt"

// Количество чётных чисел
// Подсчитай, сколько в массиве из 8 элементов чётных чисел.
// Практика работы с условиями и остатком от деления (`%`).

func kolichestvochetnih() {
	// Ввод: [2, 7, 8, 3, 4, 9, 6, 1]
	// Вывод: 4
	mass := [8]int{2, 7, 8, 3, 4, 9, 6, 1}
	fmt.Println(mass)
	count := 0
	for i := 0; i < 8; i++ {
		if mass[i]%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
