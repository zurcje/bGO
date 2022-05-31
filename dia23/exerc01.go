package main

import ("fmt"
"log"
"os")

type listaProdutos struct{
	id string
	preco float32
	quantidade int
}

func (l listaProdutos) lista() string{
	return fmt.Sprintf("id, %s, %.2f, %d;", l.id, l.preco, l.quantidade)
}

func (l listaProdutos) salvarLista(path string){
	os.WriteFile(path, []byte(l.lista()), 777)
}

func main(){
	produto := listaProdutos{
		id: "ab12",
		preco: 12.90,
		quantidade: 3,
	} 

	produto.salvarLista("./arquivo.exerc01.txt")

	arquivo, err := os.OpenFile("./arquivo.exerc01.txt", os.O_APPEND | os.O_WRONLY, 777)

	if err != nil{
		log.Printf(err.Error())
	}

	arquivo.WriteString("\nTexto")

	arquivo.Close()
}