package funcc

import "fmt"

func prevetstvie() {
	greet("Stan")
	greet("John")
}

func greet(name string) {
	fmt.Printf("Привет, %s!\n", name)
}
