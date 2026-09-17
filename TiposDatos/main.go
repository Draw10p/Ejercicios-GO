package main

import "fmt"

func main() {
	var edad int = 76
	var temperatura float64 = 36.5
	var activo bool = true
	var mensaje string = "Hola, soy tu papá"
	var dato byte = 200
	var dias [6]string = [6]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Printf("Edad: %v---TipoDato: %T\n", edad, edad)
	fmt.Printf("Temperatura: %v---TipoDato: %T\n", temperatura, temperatura)
	fmt.Printf("El estado de Ashley es: %v---TipoDato: %T\n", activo, activo)
	fmt.Printf("El mensaje de Ashley es: %v---TipoDato: %T\n", mensaje, mensaje)
	fmt.Printf("El dato de Ashley es: %v---TipoDato: %T\n", dato, dato)
	fmt.Printf("Los días de Ashley son: %v---TipoDato: %T\n", dias, dias)
	fmt.Printf("Los números de Ashley son: %v---TipoDato: %T\n", numeros, numeros)

	fmt.Println("La edad de Ashley es:", edad, "El estado de Ashley es:", activo)
	fmt.Println("La temperatura de Ashley es:", temperatura)
	fmt.Println("El mensaje de Ashley es:", mensaje)
	fmt.Println("El dato de Ashley es:", dato)
	fmt.Println("Los días de Ashley son:", dias)
	fmt.Println("Los números de Ashley son:", numeros)

	var entero int
	var flotante float64

	fmt.Print("Ingresa un número entero: ")
	fmt.Scan(&entero)
	fmt.Print("Ingresa un número flotante: ")
	fmt.Scan(&flotante)

	Multiplicacion := float64(entero) * flotante
	Cuadrado := Multiplicacion * Multiplicacion
	fmt.Println("El resultado es:", Cuadrado)

}
