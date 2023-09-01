package main

import (
	"fmt"
)

func hanoi(n_discos int, partida string, auxiliar string, destino string) int {

	var movimentos int
	movimentos = 0

	switch n_discos {
	case 1:
		movimentos++
		fmt.Printf("Deslocar 1 de %s para %s\n", partida, destino)
		return movimentos
	default:
		movimentos += hanoi(n_discos-1, partida, destino, auxiliar)
		movimentos++
		fmt.Printf("Deslocar disco %d de %s para %s\n", n_discos, partida, destino)
		movimentos += hanoi(n_discos-1, auxiliar, partida, destino)
	}

	return movimentos
}

func main() {

	var discos int
	fmt.Println("Digite a quantidade de discos para movimentação:")
	fmt.Scan(&discos)

	var movimentos int
	movimentos = hanoi(discos, "A", "B", "C")

	fmt.Printf("Movimentos feitos:%d\n", movimentos)
}
