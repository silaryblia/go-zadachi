package slice

import "fmt"

// Циклический сдвиг влево
// Сдвинь все элементы слайса влево на 1, а первый — в конец.

func sdvig() {
	// Ввод: [10 20 30 40]
	// Вывод: [20 30 40 10]

	s := []int{10, 20, 30, 40}
	fmt.Println(s)
	if len(s) > 0 {
		s = append(s[1:], s[0])
	}
	fmt.Println(s)

}
