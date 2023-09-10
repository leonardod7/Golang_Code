package main

import "fmt"

func busca(matriz [][]int, n int, k int) bool {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matriz[i][j] == k {
				return true
			}
		}
	}
	return false
}

func main() {
	matriz := [][]int{{1, 2, 3}, {3, 0, 1}, {4, 7, 8}}
	fmt.Println(matriz)
	busca(matriz, 3, 3)
}
