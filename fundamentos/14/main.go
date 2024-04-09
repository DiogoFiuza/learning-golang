package main

import "fmt"

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)

	fmt.Printf("O saldo da conta é %v", conta.saldo)
}
