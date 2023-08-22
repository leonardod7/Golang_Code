package main

import (
	"EstruturaDados_IBMEC/IBMEC/AC2/utils"
	"fmt"
)

func main() {

	// Resposta 1):

	// Criar contato
	contatoOriginal := utils.Contato{Nome: "leo", Email: "leonardo@d7par.com.br"}
	fmt.Println("Contato original:", contatoOriginal)

	// Alterar email
	contatoOriginal.ModificarEmail("leonardosantos@d7par.com.br")
	fmt.Println("Contato atualizado:", contatoOriginal)
	fmt.Println("--------------------------------------------------")

	// Resposta 2):

}
