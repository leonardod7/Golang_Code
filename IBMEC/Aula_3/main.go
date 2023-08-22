package main

import (
	"EstruturaDados_IBMEC/IBMEC/Aula_3/utils"
	outroutils "EstruturaDados_IBMEC/IBMEC/Aula_3/utils/outroutils"
	f "fmt"
)

func main() {
	f.Println("Olá, mundo!")

	f.Println(utils.Somar(4.5, 2.2))
	f.Println(utils.Multiplicar(4.5, 2.2))
	f.Println(outroutils.Dividir(8.2, 4.1))
}
