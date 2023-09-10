package main

import "fmt"

func main() {
	v := []int{1, 2, 6, 3, 9, 4, 7, 5}
	//fmt.Println(buscaSoma(v, 9))
	//fmt.Println(buscaSoma(v, 4))
	//fmt.Println(buscaSoma(v, 19))
	buscaSoma(v, 9)

}

func buscaSoma(vetor []int, alvo int) (int, int) {

	//n := len(vetor)
	var limInferior int
	limInferior = 0

	var limSuperior int
	limSuperior = len(vetor) - 1

	//fmt.Println(limInferior)
	//fmt.Println(limSuperior)
	//fmt.Println(n)
	//fmt.Println(alvo)
	//fmt.Println("-----------")

	for limInferior < limSuperior {
		var soma int
		soma = vetor[limInferior] + vetor[limSuperior]
		if soma == alvo {
			fmt.Println(vetor[limInferior], vetor[limSuperior])
			break
		} else if soma < alvo {
			limInferior += 1
			fmt.Println(limInferior)
		} else {
			limSuperior -= 1
			fmt.Println(limSuperior)
		}
		fmt.Println(limInferior)
		fmt.Println(limSuperior)
	}
	return -1, -1
}
