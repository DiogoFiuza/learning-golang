package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 44, 45, 3, 76) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
