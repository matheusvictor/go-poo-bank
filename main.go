package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.saldo += valor
		return "Saque realizado com sucesso! Saldo atual: R$", c.saldo
	}
	return "Depósito inválido! Saldo atual: R$", c.saldo
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
	fmt.Println(conta.saldo, conta.Sacar(600))
	fmt.Println(conta.saldo, conta.Sacar(-100))
	fmt.Println(conta.saldo, conta.Sacar(100))

	fmt.Println(conta2.saldo)
	mensagem, valor := conta2.Depositar(600)
	fmt.Println(mensagem, valor)

}
