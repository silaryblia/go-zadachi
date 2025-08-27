package slice

import "fmt"

// Поиск максимального элемента

// Найди наибольшее значение в массиве из 10 чисел.
// Ввод: [1, 4, 6, 8, 3, 9, 2, 0, 5, 7]
// Вывод: 9

func poiskmax() {
	s := []int{1, 4, 6, 8, 3, 9, 2, 0, 5, 7}
	max := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	fmt.Println(max)
}
