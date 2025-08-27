package cycles

import "fmt"

// Вывести таблицу умножения

func tablicaymnojeniya() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			// %-4d выравнивает числа по левому краю с шинирой 4
			fmt.Printf("%-4d", i*j)
		}
		fmt.Println()
	}
}
