package main

import "fmt"

func main() {
	salarios := map[string]int{"Diogo": 8000, "João": 3000, "Lucas": 4000}
	delete(salarios, "Lucas")
	salarios["Pedro"] = 5000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}
}
