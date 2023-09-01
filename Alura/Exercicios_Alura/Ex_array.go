// Seu Valdomiro foi à feira e comprou 3 frutas. Sua bolsa tinha capacidade para guardar 4 frutas

package main

import "fmt"

func main() {

	// Array
	var frutas [4]string

	frutas[0] = "Abacaxi"
	frutas[1] = "Laranja"
	frutas[2] = "Morango"
	fmt.Println(frutas[3])

}

// Quando os arrays são criados, eles assumem os valores padrão para os tipos de seus elementos. No caso, o tipo do array
//frutas é string e o valor padrão para cada posição do array será vazio. Portanto, o valor impresso será uma string vazia.
