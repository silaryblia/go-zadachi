package maps

import "fmt"

// Подсчёт количества цифр в числе

// Напиши функцию, которая принимает целое число и
// возвращает map[int]int, где ключ — цифра,
// а значение — сколько раз она встречается.

// Ввод: 112358
// Вывод: map[1:2 2:1 3:1 5:1 8:1]

func mapa(n int) map[int]int {
	result := make(map[int]int)

	for n > 0 {
		digit := n % 10
		result[digit]++
		n /= 10
	}
	return result
}

func podchet() {
	fmt.Println(mapa(112358))
	fmt.Println(mapa(122333445))
}
