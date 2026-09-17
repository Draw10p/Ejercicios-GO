package main

import "fmt"

func main() {
	// Definición e invocación inmediata
	func() {
		fmt.Println("¡Esto es una función anónima!")
	}()
}
