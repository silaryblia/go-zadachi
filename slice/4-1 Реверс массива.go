package slice

// Реверс массива

// Ввод: [1, 2, 3, 4, 5]
// Вывод: [5, 4, 3, 2, 1]

import "fmt"

func reverss() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
