package main

import "fmt"

func main() {
	flag := false
	var user string
	if flag {
		user = "Admin"
	} else {
		user = "Guest"
	}
	edad := 0

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)

	if edad >= 18 && user == "Admin" {
		fmt.Println("Bienvenido Admin, puede acceder al sistema.")
	} else if edad < 18 && user == "Guest" {
		fmt.Println("Bienvenido tu usario tiene permisos limitados.")
	} else {
		fmt.Println("No tienes acceso al sistema.")
	}
}
