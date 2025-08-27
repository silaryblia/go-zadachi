package cycles

import "fmt"

// Комбинации из трех чисел, дающие в сумме заданное число
// Написать программу, которая находит все комбинации из
// трех чисел до определенного предела, которые в сумме дают
// другое число.Предел трех чисел и заданное число-сумму вводить
//  с клавиатуры.

func combothree() {
	var limit, summa int
	fmt.Println("Введите верхний предел для чисел:")
	fmt.Scan(&limit)
	fmt.Println("Введите заданную сумму:")
	fmt.Scan(&summa)

	fmt.Printf("Комбинация чисел от 1 до %d, дающие в сумме %d:\n", limit, summa)

	count := 0
	for i := 1; i <= limit; i++ {
		for j := i; j <= limit; j++ {
			for k := j; k <= limit; k++ {
				if i+j+k == summa {
					fmt.Printf("%d + %d + %d = %d\n", i, j, k, summa)
					count++
				}
			}
		}
	}
	if count == 0 {
		fmt.Println("Комбинаций не найдено")
	} else {
		fmt.Printf("Всего найдено комбинаций: %d\n", count)
	}
}
