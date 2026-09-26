package Vocales

import (
	"bufio"
	"fmt"
	"os"
)

func Contar() {
	fmt.Print("Ingresa una frase: ")
	escaner := bufio.NewScanner(os.Stdin)
	escaner.Scan()
	frase := escaner.Text()

	var a, e, i, o, u int

	for _, caracter := range frase {
		switch caracter {
		case 'a', 'A', 'á', 'Á', 'à', 'À', 'ä', 'Ä':
			a++
		case 'e', 'E', 'é', 'É', 'è', 'È', 'ë', 'Ë':
			e++
		case 'i', 'I', 'í', 'Í', 'ì', 'Ì', 'ï', 'Ï':
			i++
		case 'o', 'O', 'ó', 'Ó', 'ò', 'Ò', 'ö', 'Ö':
			o++
		case 'u', 'U', 'ú', 'Ú', 'ù', 'Ù', 'ü', 'Ü':
			u++
		}
	}

	fmt.Println("Cantidad de cada vocal:")
	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
