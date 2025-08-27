package cycles

import "fmt"

// "Переворот" числа

func perevorot() {
	var n int
	fmt.Scan(&n)
	reversed := 0
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	fmt.Println(reversed)
}
