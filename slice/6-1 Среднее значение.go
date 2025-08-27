package slice

import "fmt"

// Среднее значение

// Вычисли среднее арифметическое массива чисел.
// Ввод: [2, 4, 6, 8]
// Вывод: Среднее: 5.0

func sredneee() {
	s := []int{2, 4, 6, 8}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	fmt.Println(float64(sum) / float64(len(s)))
}
