package main

import "fmt"

func main() {

	// o valor 50 aponta para a variável a, que aponta para um espaço na memória. O lugar da memória vai guardar o valor 50
	// O lugar na memória possui um endereço. Toda vez que chamamos a variável a o a vai chamar o endereço para achar o valor.
	// memoria  < -------- a <---------- 50

	// o endereço da memória fica catalogada no ponteiro

	a := 10 // o a é apontado para a memória
	// print(&a) // vai mostrar onde a variável a está sendo guardada (0xc000056768)

	var ponteiro *int = &a // esse ponteiro sabe aonde o a está armazenando valor. Esse ponteiro possui o endereço da memória
	fmt.Println(ponteiro)  // tráz o endereço da memória

	// se colocarmos um * antes, ele vai trazer o valor que está no endereço da memória
	fmt.Println(*ponteiro)

	*ponteiro = 50 // alteramos o valor que estava armazenado nesse endereço da memória
	fmt.Println(*ponteiro)

	// Observe que se printarmos o a, como a está apontado para o mesmo endereço da memória, o valor ira atualizar
	fmt.Println(a)
}
