package cycles

import "fmt"

// Вывести таблицу значений функции

func tablica() {

	fmt.Println("Таблицу значений функции y=5-x^2/2")
	fmt.Println("------------------")
	fmt.Println("   x   |   y")
	fmt.Println("------------------")

	for x := -5.0; x <= 5.0; x += 0.5 {
		y := 5 - x*x/2
		fmt.Printf("%3.2f  | %3.2f\n", x, y)
	}

}
