package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese una frase o texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	palabras := strings.Fields(texto)
	conteo := make(map[string]int)

	for _, palabra := range palabras {
		palabraLimpia := strings.ToLower(palabra)
		conteo[palabraLimpia]++
	}

	palabraMasRepetida := ""
	maxRepeticiones := 0

	fmt.Println()
	fmt.Println("--- Resultados del Conteo ---")
	for clave, valor := range conteo {
		fmt.Println("Palabra:", clave, "| Repeticiones:", valor)

		if valor > maxRepeticiones {
			maxRepeticiones = valor
			palabraMasRepetida = clave
		}
	}

	fmt.Println()
	fmt.Println("-> La palabra que más se repite es", palabraMasRepetida, "con", maxRepeticiones, "repeticiones.")
}
