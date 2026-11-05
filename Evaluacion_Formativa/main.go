package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicamento struct {
	nombre     string
	fabricante string
	fechaFab   time.Time
	vidaUtil   int
}

func (m *Medicamento) SetNombre(n string) error {
	if n == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.nombre = n
	return nil
}

func (m Medicamento) GetNombre() string {
	return m.nombre
}

func (m *Medicamento) SetFabricante(fab string) error {
	if fab == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.fabricante = fab
	return nil
}

func (m Medicamento) GetFabricante() string {
	return m.fabricante
}

func (m *Medicamento) SetFechaFab(f time.Time) error {
	if f.IsZero() {
		return errors.New("la fecha de fabricación no puede estar vacía o ser cero")
	}
	m.fechaFab = f
	return nil
}

func (m Medicamento) GetFechaFab() time.Time {
	return m.fechaFab
}

func (m *Medicamento) SetVidaUtil(meses int) error {
	if meses <= 0 {
		return errors.New("la vida útil debe ser un número positivo")
	}
	m.vidaUtil = meses
	return nil
}

func (m Medicamento) GetVidaUtil() int {
	return m.vidaUtil
}

func NuevoMedicamento(n string, fab string, fecha time.Time, vida int) (Medicamento, error) {
	med := Medicamento{}

	if err := med.SetNombre(n); err != nil {
		return med, err
	}
	if err := med.SetFabricante(fab); err != nil {
		return med, err
	}
	if err := med.SetFechaFab(fecha); err != nil {
		return med, err
	}
	if err := med.SetVidaUtil(vida); err != nil {
		return med, err
	}

	return med, nil
}

func (m Medicamento) CalcularCaducidad() time.Time {
	return m.fechaFab.AddDate(0, m.vidaUtil, 0)
}

type Pastilla struct {
	Medicamento
	dosisPorPastilla int
	requiereReceta   bool
}

func (p *Pastilla) SetDosis(dosis int) error {
	if dosis <= 0 {
		return errors.New("la dosis debe ser mayor a 0 mg")
	}
	p.dosisPorPastilla = dosis
	return nil
}

func (p Pastilla) GetDosis() int {
	return p.dosisPorPastilla
}

func NuevaPastilla(n string, fab string, fecha time.Time, vida int, dosis int, receta bool) (Pastilla, error) {
	medBase, err := NuevoMedicamento(n, fab, fecha, vida)
	if err != nil {
		return Pastilla{}, err
	}

	pastilla := Pastilla{
		Medicamento:    medBase,
		requiereReceta: receta,
	}

	if err := pastilla.SetDosis(dosis); err != nil {
		return Pastilla{}, err
	}

	return pastilla, nil
}

func (p Pastilla) MostrarDetalles() {
	recetaTexto := "No"
	if p.requiereReceta {
		recetaTexto = "Sí"
	}
	fmt.Printf("Pastilla: %s | Fabricante: %s | Dosis: %d mg | Receta: %s\n",
		p.GetNombre(), p.GetFabricante(), p.GetDosis(), recetaTexto)
}

type Jarabe struct {
	Medicamento
	volumen int
	sabor   string
}

func (j *Jarabe) SetVolumen(vol int) error {
	if vol <= 0 {
		return errors.New("el volumen debe ser mayor a 0 ml")
	}
	j.volumen = vol
	return nil
}

func (j Jarabe) GetVolumen() int {
	return j.volumen
}

func NuevoJarabe(n string, fab string, fecha time.Time, vida int, vol int, sabor string) (Jarabe, error) {
	medBase, err := NuevoMedicamento(n, fab, fecha, vida)
	if err != nil {
		return Jarabe{}, err
	}

	jarabe := Jarabe{
		Medicamento: medBase,
		sabor:       sabor,
	}

	if err := jarabe.SetVolumen(vol); err != nil {
		return Jarabe{}, err
	}

	return jarabe, nil
}

func (j Jarabe) MostrarDetalles() {
	fmt.Printf("Jarabe: %s | Fabricante: %s | Volumen: %d ml | Sabor: %s\n",
		j.GetNombre(), j.GetFabricante(), j.GetVolumen(), j.sabor)
}

func main() {
	fecha1 := time.Date(2024, time.May, 10, 0, 0, 0, 0, time.UTC)
	fecha2 := time.Date(2024, time.August, 20, 0, 0, 0, 0, time.UTC)

	pastillaParacetamol, err1 := NuevaPastilla("Paracetamol", "FarmaEcuador", fecha1, 24, 500, false)
	if err1 != nil {
		fmt.Println("Error al registrar pastilla:", err1)
		return
	}

	jarabeTos, err2 := NuevoJarabe("Jarabe para la tos", "Ecuator Lab", fecha2, 12, 120, "Cereza")
	if err2 != nil {
		fmt.Println("Error al registrar jarabe:", err2)
		return
	}

	inventarioPastillas := []Pastilla{pastillaParacetamol}
	inventarioJarabes := []Jarabe{jarabeTos}

	fmt.Println("INVENTARIO DE PASTILLAS")
	for _, p := range inventarioPastillas {
		p.MostrarDetalles()
	}

	fmt.Println("\nINVENTARIO DE JARABES")
	for _, j := range inventarioJarabes {
		j.MostrarDetalles()
	}

	fmt.Println("\nACTUALIZACIÓN DE DATOS")
	fmt.Printf("El volumen original del jarabe es de: %d ml\n", inventarioJarabes[0].GetVolumen())

	errActualizacion := inventarioJarabes[0].SetVolumen(250)

	if errActualizacion != nil {
		fmt.Println("Error al actualizar:", errActualizacion)
	} else {
		fmt.Printf("Nuevo volumen: %d ml\n", inventarioJarabes[0].GetVolumen())
	}

	fmt.Println("\nFECHAS DE CADUCIDAD")
	for _, p := range inventarioPastillas {
		caducidad := p.CalcularCaducidad()
		fmt.Printf("-> %s Este caduca el: %s\n", p.GetNombre(), caducidad.Format("02/01/2006"))
	}

	for _, j := range inventarioJarabes {
		caducidad := j.CalcularCaducidad()
		fmt.Printf("-> %s Este caduca el: %s\n", j.GetNombre(), caducidad.Format("02/01/2006"))
	}
}
