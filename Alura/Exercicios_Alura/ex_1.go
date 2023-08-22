package main

// Mário foi contratado por um restaurante para desenvolver um programa que solicita a escolha do cardápio.
// Ele implementou na linguagem Go e seu código está listado abaixo

import "fmt"

func main() {
	var prato string
	fmt.Println("Digite seu prato preferido...")
	fmt.Println("P - Pizza")
	fmt.Println("H - Hambúrguer")
	fmt.Println("B - Bife com fritas")
	fmt.Println("S - Salada Caesar")
	fmt.Println("F - Salada de Frutas")
	fmt.Println("E - Estrogonofe")
	fmt.Println("O - Outros")
	fmt.Scan(&prato)

	switch prato {
	case "B":
		fmt.Println("Com batatas Palito ou Noisete?")
	case "H":
		fmt.Println("Com Queijo ou com Ovo?")
	case "P":
		fmt.Println("Calabresa ou Napolitana?")
	case "S":
		fmt.Println("Alface ou Rúcula?")
	case "F":
		fmt.Println("Kiwi ou Frango?")
	case "E":
		fmt.Println("Carne ou Frango?")
	case "O":
		fmt.Println("Não gostou de nosso cardápio?")
	default:
		fmt.Println("Não entendi seu paladar...")
	}
}

// Quando o cliente digitar b, o que será impresso?

// Observe que como não houve nenhum tratamento para a letra minúscula, ele entrará no default

// Em outras linguagens, após cada alternativa é necessário colocar a palavra reservada break. Quais são os comportamentos do compilador da linguagem Go?
// A - No Go no uso do break não é obrigatório
// B - Se o desenvolvedor colocar break, o compilador não vai reclamar
