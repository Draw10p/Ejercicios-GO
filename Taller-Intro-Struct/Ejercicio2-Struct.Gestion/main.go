package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func MostrarInformacion(p Producto) {
	fmt.Println("Producto:", p.Nombre, "| Precio Unitario: $", p.Precio, "| Cantidad:", p.Cantidad)
}

func CalcularValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func TieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func main() {
	p1 := Producto{Nombre: "Monitor", Precio: 150.50, Cantidad: 10}
	p2 := Producto{Nombre: "Teclado Mecánico", Precio: 45.00, Cantidad: 0}
	p3 := Producto{Nombre: "Ratón Inalámbrico", Precio: 25.00, Cantidad: 25}

	inventario := []Producto{p1, p2, p3}

	fmt.Println("--- INVENTARIO DE PRODUCTOS ---")

	for _, prod := range inventario {
		MostrarInformacion(prod)

		valorTotal := CalcularValorTotal(prod)
		fmt.Println("Valor total en stock: $", valorTotal)

		if TieneStock(prod) {
			fmt.Println("Estado: Disponible")
		} else {
			fmt.Println("Estado: Sin stock")
		}

		fmt.Println("-------------------------------")
	}
}
