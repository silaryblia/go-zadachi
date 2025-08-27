package maps

import "fmt"

// Инвертирование мапы

// Есть мапа map[string]string.
// Поменяй местами ключи и значения.

// Ввод: map["cat"]="кот", map["dog"]="собака"
// Вывод: map["кот"]="cat", map["собака"]="dog"

func invert() {
	origin := map[string]string{
		"cat": "кот",
		"dog": "пес",
	}
	fmt.Println(origin)

	inverted := make(map[string]string)

	for key, value := range origin {
		inverted[value] = key
	}
	fmt.Println(inverted)
}
