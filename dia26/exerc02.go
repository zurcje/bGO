package main

import (
	"fmt"
	"os"

	guuid "github.com/google/uuid"
)

//math/rand
//random numbers

type Cliente struct {
	arquivo   string
	nome      string
	sobrenome string
	rg        string
	telefone  string
	endereco  string
}

func GerarID() string {
	newId := guuid.New().String()
	if newId == "" {
		panic("Erro ao gerar um novo ID.")
	}
	return newId
}

func VerificarSeNomeJaExiste(nome string) bool {
	LerArquivo()
	return false
}

func LerArquivo() {
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado.")
	}
}

func handlePanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

func main() {

	handlePanic()

	if !VerificarSeNomeJaExiste("Heymer") {

		cliente := Cliente{
			arquivo:   GerarID(),
			nome:      "Heymer",
			sobrenome: "Britto",
			rg:        "123456",
			telefone:  "1732555555",
			endereco:  "não te interessa",
		}

		fmt.Printf("Nome do cliente: %s \n ", cliente.nome)
	}
}
