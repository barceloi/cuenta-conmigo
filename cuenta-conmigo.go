package main

import (
	"fmt"
)

const (
	letterA = `
  A
 A A
AAAAA
A   A
A   A
`
	letterB = `
BBBB
B   B
BBBB
B   B
BBBB
`
)

func main() {
	fmt.Println("Hola")

	var inputLetter string
	fmt.Print("Ingresa una letra: ")
	fmt.Scanln(&inputLetter)

	switch inputLetter {
	case "A", "a":
		fmt.Println(letterA)
	case "B", "b":
		fmt.Println(letterB)
	// Agrega más casos para otras letras aquí
	default:
		fmt.Println("La letra ingresada no tiene una imagen ASCII asociada.")
	}

}
