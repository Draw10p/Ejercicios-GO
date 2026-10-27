package main

import "fmt"

// 1. Estructura Producto
type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

// 2. Constructor de Producto
func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

// 3. Métodos para la estructura Producto
func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("ID:", p.ID, "| Nombre:", p.Nombre, "| Precio: $", p.Precio, "| Cantidad en stock:", p.Cantidad)
}

// 4. Estructura Inventario
type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[int]Producto),
	}
}

// 5. Funciones (Métodos) para la estructura Inventario
func (i *Inventario) AgregarProducto(p Producto) {
	i.Productos[p.ID] = p
	fmt.Println("Se agregó el producto:", p.Nombre)
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.Productos, id)
	fmt.Println("Se eliminó el producto con ID:", id)
}

func (i Inventario) BuscarProducto(id int) (Producto, bool) {
	prod, existe := i.Productos[id]
	return prod, existe
}

func (i Inventario) ListarProductos() {
	fmt.Println("--- LISTADO DE PRODUCTOS EN INVENTARIO ---")
	for _, prod := range i.Productos {
		prod.MostrarInfo()
	}
	fmt.Println("------------------------------------------")
}

func main() {
	// Inicializar el inventario
	tienda := NuevoInventario()

	prod1 := NuevoProducto(1, "Laptop", 850.50, 10)
	prod2 := NuevoProducto(2, "Mouse", 25.00, 50)
	prod3 := NuevoProducto(3, "Teclado", 45.00, 30)

	// Agregar productos al inventario
	tienda.AgregarProducto(prod1)
	tienda.AgregarProducto(prod2)
	tienda.AgregarProducto(prod3)

	fmt.Println()

	tienda.ListarProductos()
	fmt.Println()

	fmt.Println("Buscando el producto con ID 1...")
	prodEncontrado, existe := tienda.BuscarProducto(1)
	if existe {
		prodEncontrado.ActualizarPrecio(800.00)
		prodEncontrado.ActualizarCantidad(5)
		tienda.AgregarProducto(prodEncontrado)
	}

	fmt.Println()
	tienda.ListarProductos()
	fmt.Println()

	// Eliminar un producto
	tienda.EliminarProducto(2)

	fmt.Println()
	tienda.ListarProductos()
}
