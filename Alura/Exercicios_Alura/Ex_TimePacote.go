// Fábio é um personal trainer e encomendou um programa para marcar o tempo de seus clientes.
// O desenvolvedor está tendo dificuldades para fazer o seguinte programa compilar. Você pode ajudá-lo?

package main

import (
	"fmt"
	"time" // importando, não tinha originalmente no exercício
)

const aquecimento = 5
const corridaLeve = 10
const corridaForte = 10
const alongamento = 1

func main() {
	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Minute)

	fmt.Println("Período de aquecimento...")
	time.Sleep(aquecimento * time.Minute)

	fmt.Println("Período de corrida leve...")
	time.Sleep(corridaLeve * time.Minute)

	fmt.Println("Período de corrida forte...")
	time.Sleep(corridaForte * time.Minute)

	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Minute)
}

//Falta importar o pacote time...
//import "time"
//...e incluir a constante alongamento:
//const alongamento = 1
