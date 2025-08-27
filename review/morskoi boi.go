package main

import "fmt"

// Морской бой
// Есть игровое поле 10х1, на нем размещается один
// корабль 4 палубы, пользователь вводит координату,
// получает ответ в формате «вы попали» если попал,
// поле обновляется. «Вы выиграли» когда корабль убит.
func main() {
	field := make([]rune, 10)
	for i := range field {
		field[i] = '@'
	}

	shipStart := 3
	shipEnd := shipStart + 4
	ship := make(map[int]bool)
	for i := shipStart; i < shipEnd; i++ {
		ship[i] = true
	}

	hits := 0

	for {
		fmt.Println(string(field))
		fmt.Println("Введите координату (0-9): ")
		var shot int
		fmt.Scan(&shot)

		if shot < 0 || shot >= len(field) {
			fmt.Println("Некорректная координата!")
			continue
		}

		if ship[shot] {
			if field[shot] == 'X' {
				fmt.Println("Вы уже сюда стреляли")
				continue
			}
			field[shot] = 'X'
			hits++
			fmt.Println("Вы попали!")
		} else {
			if field[shot] == '.' {
				fmt.Println("Вы уже сюда стреляли")
				continue
			}
			field[shot] = '.'
			fmt.Println("Вы промахнулись!")
		}

		if hits == len(ship) {
			fmt.Println(string(field))
			fmt.Println("Вы выиграли!")
			break
		}
	}
}
