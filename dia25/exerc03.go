// 1. Desenvolva as funções necessárias para permitir que a empresa calcule:
// a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
// - A função receberá as horas trabalhadas no mês e o valor da hora como
// parâmetro.
// - Esta função deve retornar mais de um valor (salário calculado e erro).
// - No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
// descontados como imposto.

//quantidade de horas trabalhadas
//valor da hora
//se salario < 20.00 - 10% de desconto

package main

import "fmt"

type info struct {
    salario float64 
	horas float32
	valorHora float32
}

func (i *info) calcSalario(horas, valorHora float64){
	i.horas = horas
	i.valorHora = valorHora
}
