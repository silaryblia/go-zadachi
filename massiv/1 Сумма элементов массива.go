package massiv

import "fmt"

// Сумма элементов массива
// Создай массив из 5 целых чисел и вычисли сумму всех его элементов.
// **Цель:** Научиться обращаться к элементам массива и использовать цикл `for`.

func summaelementov() {
	mass := [5]int{3, 5, 2, 7, 1}
	fmt.Println(mass)
	//	var mass [5]int
	//	mass[5] = [3, 5, 2, 7, 1]
	sum := 0
	for i := 0; i < 5; i++ {
		sum += mass[i]
	}
	fmt.Println(sum)

}
