package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      uint32
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre:\t\t", a.Nombre)
	fmt.Println("Apellido:\t", a.Apellido)
	fmt.Println("DNI:\t\t", a.DNI)
	fmt.Println("Fecha\t\t", a.Fecha)
}

func main() {
	alumno := Alumno{"Jonh", "Doe", 4241300, "1998/01/01"}
	alumno.detalle()
}
