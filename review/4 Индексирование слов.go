package maps

import (
	"fmt"
	"strings"
)

// Индексирование слов

//Разбей строку на слова и составь мапу
// map[string]int, где каждому слову соответствует
// индекс его первого вхождения в списке.

// Ввод: "go is fun and go is fast"
// Вывод: map[go:0 is:1 fun:2 and:3 fast:6]

func indexWords(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)

	for index, word := range words {
		if _, exists := result[word]; !exists {
			result[word] = index
		}
	}
	return result
}

func iiindex() {
	input := "go is fun and go is fast"
	result := indexWords(input)

	fmt.Println(result)

}
