package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello world!")
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()
	/// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	//arquivo2, err := os.Open("arquivo.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//reader := bufio.NewReader(arquivo2)
	//buffer := make([]byte, 10)
	//for {
	//	n, err := reader.Read(buffer)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(buffer[:n]))
	//}

	// Removendo arquivos
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
