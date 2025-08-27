package funcc

import "fmt"

func localperem() {
	scopedVar()
}

func scopedVar() {
	x := 5
	fmt.Println(x)
}
