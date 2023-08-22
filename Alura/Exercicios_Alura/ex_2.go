// Observe o código abaixo

package main

import "fmt"

func main() {
	palavraDigitada := lePalavra()
	fmt.Println("A palavra digitada foi", palavraDigitada)
}

func lePalavra() string {
	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)
	return palavra
}

// O que está faltando para que ele compile e a função retorne a palavra digitada para a função principal?
// R.: Nada, pois o código está compilado normalmente
