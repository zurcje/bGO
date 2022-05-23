// Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
// qual possamos passar N quantidade de números inteiros e devolva a média.
// Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
package main

import ("fmt"
"errors"
)

func main (){
	resultado, err := notasAlunos (6.5, 7, 8.5, 6, 5)

	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Printf("média: %.1f\n", resultado)
	}
}
func notasAlunos(notas ...float32) (float32, error) {
	var mediaNotas float32 = 0.0

	for _, nota := range notas {
		if nota < 0.0 {
			return 0.0, errors.New("Nota inválida")
		}
		mediaNotas += nota
	}

	return mediaNotas / float32(len(notas)), nil
}