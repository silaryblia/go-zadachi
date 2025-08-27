package slice

import (
	"fmt"
)

// Удаление элемента по индексу

//Дан слайс из целых чисел.
// Удали элемент по заданному индексу,
// не используя дополнительные библиотеки.

func deleteindex() {
	// Ввод: s := []int{1, 2, 3, 4, 5}, index = 2
	// Вывод: [1 2 4 5]
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	index := 2

	if index >= 0 && index < len(s) {
		s = append(s[:index], s[index+1:]...)
	}

	fmt.Println(s)
}
