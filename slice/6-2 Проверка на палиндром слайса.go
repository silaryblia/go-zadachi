package slice

import (
	"fmt"
)

// Проверка на палиндром слайса
// Проверь, является ли слайс палиндромом
// (одинаковый слева направо и справа налево).

// Ввод: [1 2 3 2 1]
// Вывод: true
// Ввод: [1 2 3 4]
// Вывод: false

func isPalindrome(s []int) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func palin() {
	fmt.Println(isPalindrome([]int{1, 2, 3, 2, 1})) // true
	fmt.Println(isPalindrome([]int{1, 2, 3, 4}))    // false
}
