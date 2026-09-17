package main

import "fmt"

func main() {
	n := 0
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&n)

	for i := 1; i < 11; i++ {

		fmt.Printf("%d x %d = %d\n", n, i, n*i)

	}

	if n%2 == 0 {

		fmt.Printf("El número es Par\n")

	} else {

		fmt.Printf("El númeroes Impar\n")

	}

	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Número pequeño")
	case n >= 6 && n <= 10:
		fmt.Println("Número mediano")
	default:
		fmt.Println("Número grande")
	}
}
