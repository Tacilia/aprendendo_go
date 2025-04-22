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
/*const (
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
}*/

// Quarto commit:
/*func main() {
	var opcao int

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1 - Dizer Olá")
		fmt.Println("2 - Somar 2 + 2")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")

		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Olá! Tudo bem com você?")
		case 2:
			fmt.Println("2 + 2 =", 2+2)
		case 3:
			fmt.Println("Saindo do programa... Até logo!")
			break // Sai do switch, mas não do for
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

		// Verifica se o usuário escolheu sair
		if opcao == 3 {
			break // Agora sim: sai do laço for
		}
	}
}

// O primeiro break dentro do switch só sai do switch, não do for.
//A verificação com if opcao == 3 é onde usamos o break para sair do laço for, encerrando o programa.

//Quinto commit - Exemplo for range (para listas, arrays, strings, mapas...):
/*nomes := []string{"Ana", "Bruno", "Carlos"}

for i, nome := range nomes {
    fmt.Println("Índice:", i, "Nome:", nome)
}
//for com três partes = ideal para contador fixo
//for com condição = ideal para repetição até algo mudar
//for infinito + break = ideal para menus ou interações contínuas
//for range = ideal para percorrer coleções de dados
*/

// Sexto commit: Exemplo de breack condicao que para o loop imediatamente
/*func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // para tudo quando i for 5
		}
		fmt.Println("Valor de i:", i)
	}
}*/

// Setimo commt - pulando número proibido e continuando o loop
func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("3 é proibido!")
			continue
		}
		fmt.Println("Número:", i)
	}
}
