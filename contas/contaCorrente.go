package contas

import "alura/go-poo-bank/clientes"

type ContaCorrente struct {
	Titular clientes.Titular // atributos iniciando com letras maiúsculas "equivale" à vibilidade
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.Saldo
	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.Saldo += valor
		return "Saque realizado com sucesso! Saldo atual: R$", c.Saldo
	}
	return "Depósito inválido! Saldo atual: R$", c.Saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valor > 0 && valor <= c.Saldo

	if podeTransferir {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
	}

	return podeTransferir
}
