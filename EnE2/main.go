package main

import (
	"errors"
	"fmt"
)

type Estudiante struct {
	edad int
}

func (e *Estudiante) GetEdad() int {
	return e.edad
}

func (e *Estudiante) SetEdad(edad int) error {
	if edad < 0 || edad > 80 {
		return errors.New("Edad no válida: debe estar entre 1 y 79 años")
	}
	e.edad = edad
	return nil
}

func NewEstudiante(edad int) (*Estudiante, error) {
	estudiante := &Estudiante{}
	err := estudiante.SetEdad(edad)
	if err != nil {
		return nil, err
	}
	return estudiante, nil
}

func main() {
	student := Estudiante{}
	err := student.SetEdad(100)
	if err != nil {
		fmt.Println("No se puede modificar", err)
	}
	fmt.Println("Hola, su edad es: ", student.GetEdad())
}
