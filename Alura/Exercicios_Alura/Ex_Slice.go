package main

import (
	"fmt"
	"reflect"
)

// Slice
func exibeNomes() {

	// slicing
	nomes := []string{"Douglas", "Leo", "Julio"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))

}

func main() {
	exibeNomes()
}

// Tudo no slice é um array. Os slices são um array. O que acontece é que o sistema cria uma abstração, um nível acima dos arays para que a gente não se preocupe
// com tamanho.

// Toda vez que queremos adicionar um elemento no slice, ele dobra a capacidade do slice. Por exemplo, se temos um slice com 3 elementos, se adicionarmos
// o 4 elemento (nome), ele vai para tamanho 6.
