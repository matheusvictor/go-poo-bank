package main

import (
	"fmt"

	"alura/go-poo-bank/clientes"
	"alura/go-poo-bank/contas"
)

func main() {
	conta := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Peter Parker",
			Cpf:       "1234.456.789-00",
			Profissao: "Homem-aranha",
		},
		Agencia: 001,
		Conta:   1234,
		Saldo:   250000.0,
	}

	titular := clientes.Titular{
		Nome:      "Dr. Estranho",
		Cpf:       "0000.1111.2222-33",
		Profissao: "Doutor, só que estranho",
	}

	conta2 := contas.ContaCorrente{
		Titular: titular,
		Agencia: 002,
		Conta:   4444,
		Saldo:   0,
	}

	fmt.Println(conta, conta2)

}
