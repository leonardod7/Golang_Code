package main

import "fmt"

func main() {

	fmt.Println("Monitorando....")
	sites := []string{"http://random-status-code.herokuapp.com/", "http://www.d7par.com.br", "http://www.alura.com.br"}

	// For que começa em 0, e vai até o menor valor do tamanho
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}
	// o range do site pode retornar 2 coisas: a posição (índice) e o valor
	for indice, site := range sites {
		fmt.Println("Estou passando na posição", indice, "do meu slice e essa posição tem o site:", site)
	}

}

func testaSite(site string) {

}
