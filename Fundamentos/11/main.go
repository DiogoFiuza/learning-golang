package main

import "fmt"

type Endereco struct {
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Client struct {
	Nome  string
	idade int
	Ativo bool
	Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	diogo := Client{
		Nome:  "Diogo",
		idade: 26,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", diogo.Nome, diogo.idade, diogo.Ativo)
}
