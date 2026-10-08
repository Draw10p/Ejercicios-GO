package main

import "fmt"

func main() {
	// Array con las notas de 6 estudiantes en 4 materias.
	notas := [6][4]float64{
		{8.5, 7.0, 9.0, 8.0},
		{6.5, 8.0, 7.5, 9.0},
		{9.0, 9.5, 8.5, 9.0},
		{7.0, 6.5, 8.0, 7.5},
		{8.0, 7.5, 7.0, 8.5},
		{9.5, 8.5, 9.0, 10.0},
	}

	nombres := []string{"Ana", "Luis", "Marta", "Pedro", "Sofia", "Diego"}
	materias := []string{"Matematicas", "Lengua", "Ciencias", "Ingles"}

	fmt.Println("NOTAS DE LOS ESTUDIANTES")
	for estudiante := 0; estudiante < len(notas); estudiante++ {
		// Esta fila del array se convierte en un slice.
		notasEstudiante := notas[estudiante][:]
		suma := 0.0

		fmt.Println("\nEstudiante:", nombres[estudiante])
		for materia := 0; materia < len(notasEstudiante); materia++ {
			fmt.Println("Materia:", materias[materia], "Nota:", notasEstudiante[materia])
			suma += notasEstudiante[materia]
		}

		promedio := suma / float64(len(notasEstudiante))
		fmt.Println("Promedio:", promedio)
	}

	fmt.Println("\nPROMEDIO POR MATERIA")
	for materia := 0; materia < len(materias); materia++ {
		notasMateria := make([]float64, len(notas))
		suma := 0.0

		for estudiante := 0; estudiante < len(notas); estudiante++ {
			notasMateria[estudiante] = notas[estudiante][materia]
			suma += notasMateria[estudiante]
		}

		promedio := suma / float64(len(notasMateria))
		fmt.Println(materias[materia], "Promedio:", promedio)
	}
}
