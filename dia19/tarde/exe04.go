// Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
// De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

// var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

// Por outro lado, você também precisa:
// - Saiba quantos de seus funcionários têm mais de 21 anos.
// - Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
// - Excluir Pedro do mapa.

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	maiorVintUm := 0

	delete(employees, "Pedro")

	employees["Frederico"] = 25

	for _, idade := range employees {
		if idade > 21 {
			maiorVintUm += 1
		}
	}
	fmt.Println("Total de funcionários com mais de 21 anos: ", maiorVintUm, "\n")
	fmt.Printf("Benjamin tem %v anos\n", int(employees[" :="]))
	
}
