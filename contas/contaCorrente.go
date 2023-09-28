package contas

import "alura/go-poo-bank/clientes"

type ContaCorrente struct {
	Titular        clientes.Titular // atributos iniciando com letras maiúsculas "equivale" à vibilidade
	Agencia, Conta int
	saldo          float64
}

func (c *ContaCorrente) ConsultarSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	}
	return "saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.saldo += valor
		return "Saque realizado com sucesso! saldo atual: R$", c.saldo
	}
	return "Depósito inválido! saldo atual: R$", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valor > 0 && valor <= c.saldo

	if podeTransferir {
		c.saldo -= valor
		contaDestino.Depositar(valor)
	}

	return podeTransferir
}
