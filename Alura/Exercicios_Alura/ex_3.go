// 1) Manoela está fazendo um programa para imprimir informações sobre determinada cidade. Por enquanto o código está assim:

package main

import "fmt"

func main() {
	cidade, populacao, capital := devolveCidadeEPopulacao()
	if capital {
		fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
	} else {
		fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
	}
}

func devolveCidadeEPopulacao() (string, int, bool) {
	return "Vila Sem Nome", 4328, true
}

// Quando ela executar o programa, o que será impresso no terminal?
// R.:"Vila Sem Nome", 4328, true.
// Repare que o valor retornado pela função devolveCidadeEPopulacao para a variável capital é true.

// 2) Manoela se empolgou com seu pequeno programa e previu uma melhoria. Agora ela quer que a função retorne também o estado da cidade.
func devolveCidadeEPopulacao2() (string, int, bool, string) {
	return "Vila Sem Nome", 4328, true, "RJ"
}

// Mas por enquanto não quer utilizá-la no programa. O que ela deve fazer na função main abaixo para que o programa compile sem utilizar o estado?
// R.: cidade, populacao, capital, _ := devolveCidadeEPopulacao()
