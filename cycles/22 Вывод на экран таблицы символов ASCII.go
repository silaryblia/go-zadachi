package cycles

import "fmt"

// Вывод на экран таблицы символов ASCII
// символы с 32-го по 126-й включительно

func acii() {
	fmt.Println("Таблица ASCII (символы с 32-го по 126-й):")
	for i := 32; i <= 126; i++ {
		fmt.Printf("%3d:%c ", i, i)

		// переход на новую строку
		if (i-31)%5 == 0 {
			fmt.Println()
		}
	}

	if (126-31)%5 != 0 {
		fmt.Println()
	}
}
