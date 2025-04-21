package main

import (
	"fmt"
)

var nome string

func main() {
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome) // fmt.Scan, fmt.Scanln, fmt.Scanf. São usados para ler valores do terminal (entrada padrão).
	fmt.Println("Olá,", nome)
}
