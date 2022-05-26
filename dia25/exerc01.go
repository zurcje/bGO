// Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
// tipo “int”.
// 2. Crie um erro personalizado com uma struct que implemente “Error()” com a
// mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
// disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
// console a mensagem “Necessário pagamento de imposto”.

package main

import (
	"fmt"
	"os"
)

var salario = 1000

type ErrorSalario struct {
	msg string
}

func (e *ErrorSalario) Error() string {
	return fmt.Sprintf(e.msg)
}

func erroPersonalizado(salario int) error {
	if salario < 15000 {
		return &ErrorSalario{
			msg: "erro: O salário digitado não está dentro do valor mínimo\n",
		}
	}
	return nil
}

func main() {
	err := erroPersonalizado(salario)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Necessário pagamento de imposto sobre o salário R$%d\n", salario)
}
