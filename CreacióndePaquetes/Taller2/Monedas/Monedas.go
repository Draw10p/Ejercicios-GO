package Monedas

import (
	"fmt"
)

func Convertir() {
	var dolares float64
	var moneda int
	var resultado float64
	var nombreMoneda string

	fmt.Print("Ingresa el valor en dólares (USD): ")
	fmt.Scan(&dolares)
	fmt.Println("\nEscoge la moneda a la que quieres convertir:")
	fmt.Println("1. EUROS")
	fmt.Println("2. LB ")
	fmt.Println("3. Won")
	fmt.Println("4. BTC")
	fmt.Print("Selecciona una opción: ")
	fmt.Scan(&moneda)

	switch moneda {

	case 1:
		resultado = dolares * 0.92
		nombreMoneda = "EUROS"
	case 2:
		resultado = dolares * 0.79
		nombreMoneda = "LB"
	case 3:
		resultado = dolares * 1350
		nombreMoneda = "WON"
	case 4:
		resultado = dolares * 0.000015
		nombreMoneda = "BTC"
	default:
		fmt.Println("Error: moneda no válida.")
		return
	}

	fmt.Printf("%.2fUSD equivalen a %.8f %s.\n", dolares, resultado, nombreMoneda)
}
