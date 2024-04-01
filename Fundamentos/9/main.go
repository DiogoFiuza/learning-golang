package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 44, 45, 3, 76))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
