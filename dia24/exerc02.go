package main

import (
	"fmt"
)

type usuario struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

func NovoProduto(nome string, preco float64) *produto {
	return &produto{
		nome:  nome,
		preco: preco,
	}
}

func Adicionar(u *usuario, p *produto, q int) {
	p.quantidade += q
	u.produtos = append(u.produtos, *p)
}

//recebe (nos parenteses)
func Deletar(u usuario) {
	u.produtos = nil
}

func main() {
	user := usuario{
		nome:      "Manu",
		sobrenome: "Moreira",
		email:     "manu@gmail.com",
	}
	Adicionar(&user, NovoProduto("Cama", 1000.00), 1)

	fmt.Printf("Produto: %s", user.produtos[0].nome)
}
