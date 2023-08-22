package main

import "fmt"
import "os"
import "net/http" // fazendo requisicões web com Go

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

// Função que retorna múltiplos valores
func devolveNomeEIdade() (string, int) {
	nome := "Leonardo"
	idade := 37
	return nome, idade
}

func iniciarMonitoramento() {
	site := "http://www.d7par.com.br"
	resp, _ := http.Get(site) // a função get retorna a resposta do site e um possível erro (caso consiga para que a gente veja).
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := lerComando()
	iniciarMonitoramento()

	//
	//nome, idade := devolveNomeEIdade()
	//fmt.Println(nome, idade)
	//
	//// Podemos também ao invés de ter que usar obrigatoriamente as duas variáveis, usar somente uma dispensando a outra.
	//nome2, _ := devolveNomeEIdade()
	//fmt.Println(nome2)

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
