package main

import "fmt"
import "os" // função para retorno de erro e exit

// Função de Boas Vindas
func exibeIntroducao() {
	nome := "Leonardo"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

// Função para leitura do comando
func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

// Função exibe menu
func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0) // Saindo do programa
	default: // utilizado para quando não cair em nenhum dos casos
		fmt.Println("Não conheço este comando")
		os.Exit(-1) // o -1 indica que ocorreu algum erro inesperado
	}

}
