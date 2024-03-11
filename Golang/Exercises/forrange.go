package main

import "fmt"

func main(){
	a:= "Hello, D@n!3l"
	for i, c:= range a{
		fmt.Printf("%d: %s \n", i, string(c))
	}
	fmt.Printf("Tamanho de a", len(a))
}