package main

import (
	"fmt"
)

//Primeiro Commit:
//func main() {
//fmt.Print("Digite seu nome: ")
//fmt.Scanln(&nome) // fmt.Scan, fmt.Scanln, fmt.Scanf. São usados para ler valores do terminal (entrada padrão).
//fmt.Println("Olá,", nome)
//}

//Segundo Commit:
/*var nome string
var num1, num2 float64
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
}*/

// Terceiro commit:
const (
	_ = iota // ignoramos o 0 (opcional, mas comum)

	// Cada linha recebe um valor crescente de iota
	Soma          = iota + 1       // 1 + 1 = 2
	Subtracao     = (iota - 1) * 2 // (2 - 1) * 2 = 2
	Multiplicacao = iota * 3       // 3 * 3 = 9
	Divisao       = 12 / iota      // 12 / 4 = 3
)

func main() {
	fmt.Println("Constantes geradas com iota e operações:")
	fmt.Println("Soma:         ", Soma)
	fmt.Println("Subtração:    ", Subtracao)
	fmt.Println("Multiplicação:", Multiplicacao)
	fmt.Println("Divisão:      ", Divisao)
}
