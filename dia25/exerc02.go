// Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
// “Error()”, seja implementado “errors.New()”.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var salario = 1000

func verificaSalario(salario int) error {
	if salario < 15000 {
		return errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}
	return nil
}

func main() {
	err := verificaSalario(salario)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Necessário pagamento de imposto sobre o salário R$%d\n", salario)
}
