package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//	fmt.Print(i)
	//}

	numeros := []string{"one", "two", "three", "four"}

	for i, v := range numeros {
		fmt.Println(i, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//for {
	//	println("looping infinito")
	//}

}
