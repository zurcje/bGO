package main

import "fmt"

var umidade, pressao float32
var temperatura int8

func main() {

	umidade, temperatura, pressao = 11.5, 20, 35.9

	fmt.Println(umidade, temperatura, pressao)
}
