package massiv

import "fmt"

// Поиск элемента

//Проверь, есть ли в массиве определённое число.
// Если есть — выведи индекс

func poiskelementa() {
	// Ввод: [5, 3, 8, 6], искомое число: 8
	// Вывод: Элемент найден на позиции 2
	mass := []int{5, 3, 8, 6}
	fmt.Println(mass)
	for i := 0; i < len(mass); i++ {
		if mass[i] == 8 {
			fmt.Println("Элемент найден на позиции ", i)
		}
	}

}
