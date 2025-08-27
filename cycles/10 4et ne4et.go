package cycles

import "fmt"

func chetnechet() {
	var number int
	fmt.Println("Введите ваше число:")
	fmt.Scan(&number)
	chet := 0
	nechet := 0
	for number > 0 {
		digit := number % 10
		if digit%2 == 0 && digit != 0 {
			chet++
		} else if digit%2 != 0 && digit != 0 {
			nechet++
		}

		number /= 10
	}
	fmt.Printf("Количество четных чисел: %d\n", chet)
	fmt.Printf("Количество четных чисел: %d\n", nechet)
}
