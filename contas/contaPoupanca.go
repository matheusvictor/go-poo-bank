package contas

import "alura/go-poo-bank/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) ConsultarSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	}
	return "saldo insuficiente!"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor >= 0 {
		c.saldo += valor
		return "Saque realizado com sucesso! saldo atual: R$", c.saldo
	}
	return "Depósito inválido! saldo atual: R$", c.saldo
}
