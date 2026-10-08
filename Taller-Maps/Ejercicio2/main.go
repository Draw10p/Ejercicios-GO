package main

import "fmt"

func actividadGanadora(votos map[string]int) string {
	ganadora := ""
	mayorVotos := 0

	for actividad, cantidad := range votos {
		if cantidad > mayorVotos {
			ganadora = actividad
			mayorVotos = cantidad
		}
	}

	return ganadora
}

func main() {
	votos := map[string]int{
		"Deportes":    0,
		"Videojuegos": 0,
		"Cine":        0,
		"Musica":      0,
	}

	fmt.Println("Elige una actividad para cada voto:")
	fmt.Println("1. Deportes")
	fmt.Println("2. Videojuegos")
	fmt.Println("3. Cine")
	fmt.Println("4. Musica")

	for voto := 1; voto <= 5; voto++ {
		var opcion int
		fmt.Println("Voto", voto, ":")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			votos["Deportes"]++
		case 2:
			votos["Videojuegos"]++
		case 3:
			votos["Cine"]++
		case 4:
			votos["Musica"]++
		default:
			fmt.Println("Opcion no valida")
		}
	}

	fmt.Println("Resultados:")
	for actividad, cantidad := range votos {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	fmt.Println("Actividad ganadora:", actividadGanadora(votos))
}
