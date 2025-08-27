package cycles

import (
	"fmt"
	"math"
)

// Возведение числа в степень
func mastepen() {
	var a, b float64
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	fmt.Println("Введите степень:")
	fmt.Scan(&b)
	result := math.Pow(a, b)
	fmt.Println(result)

}
