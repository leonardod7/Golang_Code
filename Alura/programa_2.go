package main

import "fmt"

func main() {
	nome := "Leonardo"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	// No Go, ele não aceita na estrutura do if qualquer retorno que não seja uma variável booleana
	//if comando == 1 {
	//	fmt.Println("Monitorando")
	//} else if comando == 2 {
	//	fmt.Println("Exibindo Logs")
	//} else if comando == 0 {
	//	fmt.Println("Saindo do programa")
	//} else {
	//	fmt.Println("Não conheço este comando")
	//}

	// Controle de fluxo com o Switch ao invés da estrutura acima if
	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do programa")
	default: // utilizado para quando não cair em nenhum dos casos
		fmt.Println("Não conheço este comando")
	}

	// Na linguagem Go, não precisamos colocar o break dentro do switch, pois ao encontrar o valor do case, ele automaticamente sai da estrutura.
}
