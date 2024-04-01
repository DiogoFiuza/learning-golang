package main

import "fmt"

func main() {
	defer fmt.Println("Teste 1")
	defer fmt.Println("Teste 2")
	fmt.Println("Teste 3")
}
