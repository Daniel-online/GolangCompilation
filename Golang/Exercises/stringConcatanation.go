package main

import "fmt"

func main(){

	var primeiroNome string
	var sobreNome string

	fmt.Println("Digite seu primeiro nome: ")
	fmt.Scanln(&primeiroNome)
	fmt.Println("Digite seu sobrenome: ")
	fmt.Scanln(&sobreNome)

	fmt.Printf("Seu nome é  " + primeiroNome + " " +sobreNome )
}