package main

import "fmt"

func main(){
	user := map[string] string{
		"name": "Daniel",
		"age": "24",
		"job": "Golang developer",
	}

	fmt.Println(user)
}