package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

func NewProducto(nombre string, precio float64, stock int) Producto {
	return Producto{
		Nombre: nombre,
		precio: precio,
		stock:  0,
	}
}
func calcularStock(producto Producto) float64 {
	return float64(producto.stock) * producto.precio
}

func (producto *Producto) NuevoStock(cantidad int) {
	if cantidad > 0 {
		producto.stock += cantidad
	}
}

func (producto *Producto) Vender(cantidad int) bool {
	if cantidad <= 0 || cantidad > producto.stock {
		return false
	}
	producto.stock -= cantidad
	return true
}

func main() {
	product := Producto{}
	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	prodcut3 := NewProducto("Mouse", 25.99, 50)
	fmt.Println(prodcut3)

	fmt.Println(product)
	fmt.Println(product2)
	fmt.Println("El valor del inventario del producto", product.Nombre, " es de:", calcularStock(product))
	fmt.Println("El valor del inventario del producto", product2.Nombre, " es de:", calcularStock(product2))

	cantidadVendida := 50
	ventaRealizada := product.Vender(cantidadVendida)
	if ventaRealizada {
		fmt.Println("\nVerificación de venta:", product.Nombre)
		fmt.Println("Vendido:", cantidadVendida)
		fmt.Println("Stock sobrante:", product.stock)
		fmt.Println("Valor del inventario restante:", calcularStock(product))
	} else {
		fmt.Println("\nLa venta no se pudo realizar.")
		fmt.Println("Stock disponible:", product.stock)
	}

	cantidadVendida = 20
	ventaRealizada = product2.Vender(cantidadVendida)
	if ventaRealizada {
		fmt.Println("\nVerificación de venta:", product2.Nombre)
		fmt.Println("Vendido:", cantidadVendida)
		fmt.Println("Stock sobrante:", product2.stock)
		fmt.Println("Valor del inventario restante:", calcularStock(product2))
	} else {
		fmt.Println("\nLa venta no se pudo realizar para:", product2.Nombre)
		fmt.Println("Solicitado:", cantidadVendida)
		fmt.Println("Stock disponible:", product2.stock)
	}
}
