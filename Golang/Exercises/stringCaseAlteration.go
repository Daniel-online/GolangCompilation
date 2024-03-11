package main

import (
	"fmt" 
	"strings"
)

func main(){
	var w string

	fmt.Println("Deixe uma frase em várias formas diferentes:\n")
	fmt.Scanln(&w)
	fmt.Println(strings.ToUpper(w))
	fmt.Println(strings.ToLower(w))
	fmt.Println(strings.HasPrefix(w, "S"))
	fmt.Println(strings.HasSuffix(w, "h"))
	fmt.Println(strings.Contains(w, "a"))
	fmt.Println(strings.Count(w, "a"))
	fmt.Println(len(w))
	
}