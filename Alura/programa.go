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

	// Função que captura o input do usuário e salva em uma variável
	var comando int

	// na declaração do scan, o primeiro item (%d) é sempre o movificador. O segundo (&comando), é o endereço da variável que queremos salvar o resultado
	// Não podemos passar somente o comando, pois o comando nesse caso é apenas zero, pois não foi declarado. O que se espera é um endereço, um ponteiro para aquela variável.
	// Esse endereço da memória da variável, para decobrirmos, basta colocarmos o & na frente da variável

	fmt.Scanf("%d", &comando)                                       // o %d significa que é um inteiro
	fmt.Println("O endereço da minha variável comando é", &comando) // endereço do lugar na memória do computador
	fmt.Println("O comando escolhido foi", comando)

	// Podemos usar a variável Scan sem precisarmos definir o modificador, pois ao declararmos o tipo da variável ele automaticamente vai entender o seu tipo na entrada
	var comando2 int
	fmt.Scan("o endereço da minha variável comando2 é", &comando2)

	// Se por algum motivo digitarmos o valor "a" tendo declarado comando2 como uma variável do tipo int, ele automaticamente recusa a string digitada e assume zero.

}
