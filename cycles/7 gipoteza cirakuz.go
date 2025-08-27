package cycles

// Проверка гипотезы Сиракуз

import "fmt"

func cirakuz() {

	fmt.Println("Гипотеза Сиракуз (20-30):")
	for n := 20; n <= 30; n++ {
		fmt.Printf("\n%d: ", n)
		x := n
		for x != 1 {
			fmt.Printf("%d ", x)
			if x%2 == 0 {
				x /= 2
			} else {
				x = (x*3 + 1) / 2
			}
		}
		fmt.Print("1")
	}
}
