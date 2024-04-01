package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	result := a + b
	if result >= 50 {
		return result, errors.New("A soma é maior que 50")
	}
	return result, nil
}
