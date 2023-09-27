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

}
