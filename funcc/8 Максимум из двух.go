package funcc

import "fmt"

func maxizdvyh() {
	fmt.Println(max(5, 4))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
