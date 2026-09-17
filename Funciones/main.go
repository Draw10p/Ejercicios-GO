package main

import (
	"fmt"
)

func saludar() {
	fmt.Println("¡Hola es mi primera función!")
}
func bienvenida(nombre string) {
	fmt.Printf("¡Bienvenido, %s!\n", nombre)
}

func suma_resta(a, b int) (int, int) {
	suma := a + b
	if b > a {
		return suma, b - a
	}
	return suma, 0
}

func main() {
	var usr string
	var a, b int

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&usr)
	saludar()
	bienvenida(usr)

	fmt.Println("Ingrese el primer número:")
	fmt.Scan(&a)
	fmt.Println("Ingrese el segundo número:")
	fmt.Scan(&b)

	suma, resta := suma_resta(a, b)
	fmt.Println("El resultado de la suma es: ", suma)
	fmt.Println("El resultado de la resta es: ", resta)

	fmt.Println("La sumatoria de los números es: ", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
