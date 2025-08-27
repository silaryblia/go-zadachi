package slice

import "fmt"

// Количество чётных чисел

// Подсчитай, сколько в массиве из 8 элементов чётных чисел.

func kolvochetnihslice() {
	// Ввод: [2, 7, 8, 3, 4, 9, 6, 1]
	// Вывод: 4
	s := []int{2, 7, 8, 3, 4, 9, 6, 1}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			count++
		}
	}
	fmt.Println(count)

}
