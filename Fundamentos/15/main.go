package main

import "fmt"

func main() {
	var minhaVar interface{} = "10"
	println(minhaVar.(string))
	val, ok := minhaVar.(string)
	fmt.Printf("Minha variável tem um valor %v e o status é %v", val, ok)
}
