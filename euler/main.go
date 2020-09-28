package main

import "fmt"

func main() {
	var e float64

	for i := 0; i < 50; i++ {
		fact := 1
		for j := i; j > 0; j-- {
			fact *= j
		}
		e = e + (1.0 / float64(fact))
	}
	fmt.Print(e)
}
