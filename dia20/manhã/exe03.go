package main

import "fmt"

const (
	categoriaA = 3000.00
	categoriaB = 1500.00
	categoriaC = 1000.00

	bonusA = 0.5
	bonusB = 0.2

	horaExtra = 160
)

func main() {
	fmt.Printf("Salario total: %.2f\n", calcularSalario("B", 159))
}

func calcularSalario(categoria string, horas float64) float64 {
	var salarioBruto float64

	if categoria == "A" {
		salarioBruto = categoriaA * horas

		if horas > horaExtra {
			salarioBruto += salarioBruto * bonusB
		}
	}

	if categoria == "B" {
		salarioBruto = categoriaB * horas

		if horas > horaExtra {
			salarioBruto += salarioBruto * bonusB
		}
	}

	if categoria == "C" {
		salarioBruto = categoriaC * horas
	}

	return salarioBruto
}


