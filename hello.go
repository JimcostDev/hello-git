package main

import "fmt"

// GetGreeting retorna el mensaje de saludo
func GetGreeting() string {
	return "Hola Mundo en Go!"
}

// PrintGreeting imprime el mensaje de saludo
func PrintGreeting() {
	fmt.Println(GetGreeting())
}

func main() {
	// Hola Mundo
	PrintGreeting()
}
