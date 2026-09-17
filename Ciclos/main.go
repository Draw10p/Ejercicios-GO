package main

import "fmt"

func main() {

	fmt.Println("Bienvenidos a la Clase de Ciclos en Go")

	fmt.Println("******Bucle Normal******")

	for i := 0; i < 11; i++ {
		fmt.Println("Los valor en i son:", i)
	}
	for {
		fmt.Println("Este es un bucle infinito")
		break
	}
	for rango := range [10]int{} {
		fmt.Println("Los valores en rango son:", rango)
	}
}
