package funcc

import "fmt"

func doublee() {
	fmt.Println(double(3))
}

func double(n int) int {
	return n * 2
}
