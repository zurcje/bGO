
package main

import "fmt"

type usuario struct {
    nome string 
	sobrenome string
	idade int
	email string
	senha string
}

func (u *usuario) alterarNome(nome, sobrenome string){
	u.nome = nome
	u.sobrenome = sobrenome
}

func (u *usuario) alterarIdade(idade int){
	u.idade = idade
}

func (u *usuario) alterarEmail(email string){
	u.email = email
}

func (u *usuario) alterarSenha(senha string) {
	u.senha = senha
}

func main(){
	user := &usuario{
		nome: "jessica",
		sobrenome: "cruz",
		idade: 30,
		email: "jwdcruz@gmail.com",
		senha: "abc123",
	}
	user.alterarNome("heymer", "carbonara")
	user.alterarIdade(65)
	user.alterarEmail("heymer@hotmail.com")
	user.alterarSenha("123abc")

	fmt.Printf("Nome: %s\n", user.nome)
	fmt.Printf("Sobrenome: %s\n", user.sobrenome)
	fmt.Printf("Idade: %d\n", user.idade)
	fmt.Printf("E-mail: %s\n", user.email)
	fmt.Printf("Senha: %s\n", user.senha)
}




