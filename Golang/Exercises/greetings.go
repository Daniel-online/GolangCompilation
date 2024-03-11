package main

import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println("Insira seu nome:")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Olá, %s! Bem vindo ao Golang! ", name)
}