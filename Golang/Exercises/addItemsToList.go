package main

import "fmt"

func main(){
	var item string
	var choice int
	itemList := []string{}
	
	fmt.Println("Insira um item na sua lista: ")
	fmt.Scanln(&item)
	itemList = append(itemList, item)

	fmt.Printf("Item adicionado com sucesso! Para ver sua lista, pressione 1: ")
	fmt.Scanln(&choice)
	
	fmt.Println(itemList)

}