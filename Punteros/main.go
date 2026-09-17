package main

import "fmt"

func main() {
	numero := 10
	puntero := &numero
	fmt.Println("El valor de numero es:", numero)
	fmt.Println("La dirección de memoria de numero es:", &numero)
	fmt.Println("El valor al que apunta el puntero es:", *puntero)
}
