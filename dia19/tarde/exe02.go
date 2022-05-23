//clientes com mais de 22 anos, empregados e com mais de um ano de atividade.
//não cobra juros para quem tem um salário superior a US$ 100.000.

// É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
// de acordo com cada caso.
// Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

//Se for maior que 22 anos e tiver mais que um ano de carteira = true
//Se for maior que 22 e tiver menos que 1 ano = false
//Se o salario for superior a US$ 100.000 isento de juros

package main

import "fmt"

func main() {

	idade := 29
	tempTrab := 2
	salario := 9000

	if idade > 22 && tempTrab > 1 && salario < 100000 {
		fmt.Println("Elevível para empréstimo")
	} else if idade > 22 && tempTrab > 1 && salario > 100000 {
		fmt.Println("Elevível para empréstimo sem juros")
	} else {
		fmt.Println("Ilegível para empréstimos")
	}

}
