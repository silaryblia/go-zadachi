package massiv

import "fmt"

// Среднее значение

// Вычисли среднее арифметическое массива чисел.
// Практиковать деление суммы на количество элементов (float64).

func srednee() {
	// Ввод: [2, 4, 6, 8]
	// Вывод: Среднее: 5.0
	mass := []int{2, 4, 6, 8}
	fmt.Println(mass)
	sum := 0
	for i := 0; i < len(mass); i++ {
		sum += mass[i]
	}
	srednee := float64(sum) / float64(len(mass))
	fmt.Println("Среднее: ", srednee)
}
