package slice

import "fmt"

// Фильтрация чётных чисел
// Напиши функцию, которая принимает []int и возвращает
// новый слайс, содержащий только чётные числа.

// Ввод: [1 2 3 4 5 6]
// Вывод: [2 4 6]

func filter(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func filtra() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	evenS := filter(s)
	fmt.Println(evenS)
}
