package main

import "fmt"

// Creación Menú
func main() {
	var opcion string

	for {
		fmt.Println("\n--- MENÚ ---")
		fmt.Println("1. Promedio de notas del curso")
		fmt.Println("2. Suma de números del 1 al n")
		fmt.Println("3. Convertir Celsius a Fahrenheit")
		fmt.Println("4. Convertir Fahrenheit a Celsius")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción o escriba 'salir': ")
		fmt.Scan(&opcion)
		// Opciones del Menú
		switch opcion {
		case "1":
			opcionPromedioCurso()
		case "2":
			sumadeN()
		case "3":
			C_a_F()
		case "4":
			F_a_C()
		case "0", "salir":
			fmt.Println("Programa finalizado.")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Parte 1: Promedio del Curso
func averageGrade(notas []float64) float64 {
	total := 0.0
	count := 0
	for _, nota := range notas {
		total += nota
		count++
	}
	return total / float64(count)
}

func opcionPromedioCurso() {
	var numestudiantes int

	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&numestudiantes)

	if numestudiantes <= 0 {
		fmt.Println("La cantidad de estudiantes debe ser mayor que 0.")
		return
	}
	notas := make([]float64, 0, numestudiantes)
	for estudiante := 1; estudiante <= numestudiantes; estudiante++ {
		var nota float64

		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", estudiante)
		fmt.Scan(&nota)

		if nota < 0 || nota > 100 {
			fmt.Println("La nota debe estar entre 0 y 100.")
			estudiante--
		} else {
			notas = append(notas, nota)
		}
	}

	average := averageGrade(notas)
	fmt.Printf("Promedio del curso: %.2f\n", average)

	if average >= 70 {
		fmt.Println("Curso aprobado")
	} else {
		fmt.Println("Curso reprobado")
	}
	switch {
	case average >= 90:
		fmt.Println("Excelente desempeño")
	case average >= 80:
		fmt.Println("Buen desempeño")
	case average >= 70:
		fmt.Println("Desempeño satisfactorio")
	default:
		fmt.Println("Necesita mejorar")
	}
}

// Parte 2: Suma de números del 1 al n
func sumadeN() {
	var n int

	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)

	suma := 0
	for numero := 1; numero <= n; numero++ {
		suma += numero
	}

	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, suma)
}

// Parte 3: De Celsius a Fahrenheit
func C_a_F() {
	var celsius float64

	fmt.Print("Ingrese la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("La temperatura en Fahrenheit es: %.2f\n", fahrenheit)
}

// Parte 4: De Fahrenheit a Celsius
func F_a_C() {
	var fahrenheit float64

	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("La temperatura en Celsius es: %.2f\n", celsius)
}
