package funcc

import "fmt"

var counter int = 5

func globa() {
	increment()
	increment()
	increment()

}

func increment() {
	fmt.Println(counter)
}
