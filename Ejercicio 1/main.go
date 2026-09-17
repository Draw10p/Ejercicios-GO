package main

import "fmt"

func main() {
	n := 0
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	switch {
	case n <= 10:
		fmt.Println("Número Pequeño")
	case n <= 20:
		fmt.Println("Número Mediano")
	default:
		fmt.Println("Número Grande")
	}
}
