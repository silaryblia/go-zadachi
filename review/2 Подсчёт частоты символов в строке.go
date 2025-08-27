package maps

import "fmt"

// Подсчёт частоты символов в строке

// Прочитай строку и создай map[rune]int,
// где хранятся все символы и их количество.

// Ввод: "banana"
// Вывод: map[b:1 a:3 n:2]

func odchestroki() {
	str := "banana"
	result := make(map[string]int)

	for _, char := range str {
		result[string(char)]++
	}
	fmt.Println(result)
}
