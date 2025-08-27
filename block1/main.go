package main

import "fmt"

func main() {
	fmt.Println("Совершенные числа от 1 до 10000:")
	for n := 1; n <= 10000; n++ {
		sum := 0
		for i := 1; i <= n/2; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		if sum == n {
			fmt.Println(n)
		}
	}
}
