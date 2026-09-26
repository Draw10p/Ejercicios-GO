package main

import (
	"Taller2/Taller2/Monedas"
	"Taller2/Taller2/Vocales"
	"fmt"
)

func main() {
	var opcion int

	for {
		fmt.Println("\n===== MENÚ PRINCIPAL =====")
		fmt.Println("1. Conversor de monedas")
		fmt.Println("2. Contador de vocales")
		fmt.Println("3. Salir")
		fmt.Print("Selecciona una opción: ")
		fmt.Scan(&opcion)
		fmt.Scanln()

		switch opcion {
		case 1:
			Monedas.Convertir()
		case 2:
			Vocales.Contar()
		case 3:
			fmt.Println("Programa finalizado.")
			return
		default:
			fmt.Println("Error: opción no válida.")
		}
	}
}
