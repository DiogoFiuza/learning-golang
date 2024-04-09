package main

import (
	"curso-golang/Fundamentos/18/matematica"
	"fmt"
)

func main() {
	resultadoSoma := matematica.Soma(4, 5)
	fmt.Printf("O valor da soma é igual a %v \n", resultadoSoma)
}
