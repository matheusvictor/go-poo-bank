package main

import (
	"fmt"

	"alura/go-poo-bank/clientes"
	"alura/go-poo-bank/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaCorrente := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Peter Parker",
			Cpf:       "1234.456.789-00",
			Profissao: "Homem-aranha",
		},
		Agencia: 001,
		Conta:   1234,
	}

	fmt.Println(contaCorrente.ConsultarSaldo())
	contaCorrente.Depositar(500)
	PagarBoleto(&contaCorrente, 100)
	fmt.Println(contaCorrente.ConsultarSaldo())

	titular := clientes.Titular{
		Nome:      "Dr. Estranho",
		Cpf:       "0000.1111.2222-33",
		Profissao: "Doutor, só que estranho",
	}

	contaPoupanca := contas.ContaPoupanca{
		Titular:  titular,
		Agencia:  002,
		Conta:    4444,
		Operacao: 007,
	}

	fmt.Println(contaPoupanca.ConsultarSaldo())
	contaPoupanca.Depositar(300)
	PagarBoleto(&contaPoupanca, 100)
	fmt.Println(contaPoupanca.ConsultarSaldo())

}
