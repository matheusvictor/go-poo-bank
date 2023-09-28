package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	conta := ContaCorrente{
		titular: "Fulano",
		agencia: 001,
		conta:   1234,
		saldo:   100.50,
	}

	conta2 := ContaCorrente{
		"Beltrano",
		001,
		5678,
		100.50,
	}

	fmt.Println(conta, conta2)

	var conta3 *ContaCorrente
	fmt.Println(conta3)
	conta3 = new(ContaCorrente)
	conta3.titular = "Cris"
	fmt.Println("Conteúdo NO endereço de memória: ", conta3)
	fmt.Println("Conteúdo DO endereço de memória: ", *conta3)

}
