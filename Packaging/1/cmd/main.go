package main

import (
	"fmt"
	"github.com/DiogoFiuza/learning-golang/7/1/math"
)

func main() {

	calc := math.Math{2, 3}
	value := calc.Add()

	fmt.Printf("The result is: %d\n", value)

	fmt.Println("Hello, world!")
}
