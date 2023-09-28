package main

import (
	"alura/go-poo-bank/contas"
	"fmt"
)

func main() {
	conta := contas.ContaCorrente{
		Titular: "Fulano",
		Agencia: 001,
		Conta:   1234,
		Saldo:   100.50,
	}

	conta2 := contas.ContaCorrente{
		"Beltrano",
		001,
		5678,
		100.50,
	}

	fmt.Println(conta, conta2)
	fmt.Println(conta.Saldo, conta.Sacar(600))
	fmt.Println(conta.Saldo, conta.Sacar(-100))
	fmt.Println(conta.Saldo, conta.Sacar(100))

	fmt.Println(conta2.Saldo)
	mensagem, valor := conta2.Depositar(600)
	fmt.Println(mensagem, valor)

	statusTransferencia := conta2.Transferir(200, &conta)
	fmt.Println(statusTransferencia, conta, conta2)

}
