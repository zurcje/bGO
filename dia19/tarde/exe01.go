package main

import "fmt"

func main() {
	palavra := "Academica"

	fmt.Println(len(palavra))

	for _, letra := range palavra {

		fmt.Println(string(letra))
	}
}
