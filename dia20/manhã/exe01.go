//criar uma função que retorne o
// imposto de um salário.

//um salário de R$50.000 e será descontado 17% do salário. 

//Um outro funcionário ganha um salário de $150.000, e nesse
// caso será descontado mais 10%.

//60 mil - 10200

//60.000 x (17/100)

package main

import "fmt"

const(
	descontoImposto1 = 0.17
	descontoImposto2 = 0.27
)


func main(){

	salarioFuncionario := 60000.00

	fmt.Printf("Desconto de R$%.2f sobre o salário R$%.2f\n", calcImposto(float32(salarioFuncionario)), salarioFuncionario)
	
}

func calcImposto(salario float32) float32{
	if salario >= 150000{
		return salario * descontoImposto2
	}
	return salario * descontoImposto1
}