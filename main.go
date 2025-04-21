package main

import (
	"fmt"
)

var nome string
var num1, num2 float64

//func main() {
//fmt.Print("Digite seu nome: ")
//fmt.Scanln(&nome) // fmt.Scan, fmt.Scanln, fmt.Scanf. São usados para ler valores do terminal (entrada padrão).
//fmt.Println("Olá,", nome)
//}

func main() {
	// Entrada do nome
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	// Entrada dos números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	// Soma e resposta formatada
	resultado := num1 + num2
	mensagem := fmt.Sprintf("Olá, %s! A soma de %.2f + %.2f é igual a %.2f.", nome, num1, num2, resultado)

	// Exibe a resposta
	fmt.Println(mensagem)
}
