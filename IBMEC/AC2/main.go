package main

import (
	"EstruturaDados_IBMEC/IBMEC/AC2/utils"
	"fmt"
)

func main() {

	// Declarando o array para permitir as operações de incluir e excluir contatos
	var contatos [5]*utils.Contato
	var opcao int

	// Criando interface
	fmt.Println("1. Contatos Listagem")
	fmt.Println("2. Adicione um Contato")
	fmt.Println("3. Exclua um Contato")
	fmt.Println("4. Logout")
	fmt.Println("--------------------")
	fmt.Print("Escolha uma das opções: ")
	fmt.Scan(&opcao)

	switch opcao {

	// Lista de contatos
	case 1:
		listaDadosUsuarios := utils.VisualizarListaContatos(contatos)
		for _, dadoContato := range listaDadosUsuarios {
			fmt.Println(dadoContato)
		}

	// se for o comando adicionar
	case 2:
		var nome, email string
		fmt.Println("Digite o Nome:  ")
		fmt.Scan(&nome)
		fmt.Println("----------------")
		fmt.Println("Digite o Email: ")
		fmt.Scan(&email)
		contato := utils.Contato{nome, email}
		utils.AdicionarContato(contato, &contatos)
		fmt.Println(contatos) // array com os contatos e as posições
		fmt.Println("Contato inserido")

	// se for para excluir o contato
	case 3:
		utils.RemoverContato(&contatos)
		fmt.Println("Contato excluído")

	}
	fmt.Println(utils.VisualizarListaContatos(contatos))

}
