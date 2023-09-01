// Reinaldo é um Scrum Master que deseja implementar a idéia no Go. Começou por montar um slice com os pontos do Scrum.
// Como Reinaldo pode saber quantas cartas ele vai precisar para fazer a técnica? Dica: verifique o tamanho do slice.

package main

import "fmt"

func main() {
	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
	tamanho := len(pontosPlanningPoker)
	fmt.Println(tamanho)

	pontosPlanningPoker = append(pontosPlanningPoker, 40)
	fmt.Println(cap(pontosPlanningPoker))

}

// Contudo, ele se esqueceu de colocar uma última pontuação, 40. Em vez de mudar o código de inicialização do slice, usou a função append.
// R."Quando é necessário colocar mais elementos do que sua capacidade atual, o slice dobra a capacidade.
