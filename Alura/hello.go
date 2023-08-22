package main

import "fmt"
import "reflect"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Meu primeiro programa em Go!")

	// Variável
	var nome string = "Douglas"
	var versao float32 = 1.1

	// No Go, quando não atribuímos nenhum valor a variável, ela começa com o valor zero! Dependendo do tipo da variável, ele retorna um valor
	// No caso de string, quanado não declaramos ela, o Go retorna um espaço vazio

	var idadeTeste int
	var idade int = 23

	fmt.Println("Olá Sr", nome, "sua idade é:", idadeTeste)
	fmt.Println("Olá Sr", nome, "sua idade é:", idade)
	fmt.Println("Este programa está na versão", versao)

	// o Go consegue inferir o tipo da variável. Dessa forma, se não declararmos os tipos das variáveis, ele automaticamente entende
	var idade2 = 23
	var nome2 = "Leo"
	fmt.Println("Olá Sr", nome2, "sua idade é:", idade2)

	// Para as variáveis do tipo float, é importante declararmos o seu tipo pois há 2 tipos e o Go automaticamente vai selecionar a maior (no caso, o float64)
	// Isso pode ser um problema, pois dependendo do tamanho da variável pode ser mais do que ela precisa de armazenagem em memória

	// Vamos usar o reflect para descobrir o tipo da variável
	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))

	// o Go tem um declarador de variável curto, que é o :=
	// Esse formato, quando usado, significa que estamos declarando uma variável e inferindo o seu tipo automaticamente

	sobrenome := "Santos"
	idade3 := 24

	fmt.Println("a idade do", sobrenome, "é de", idade3)

}

// Como convesão da linguagem go, quando importamos o pacote chamamos ele e colocamos a função na sequencia com letra maiúscula.
// go build "nome do arquivo.go", torna o arquivo executável

// go run hello.go
