package main

import (
	"fmt"
)

type People struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Persona  People
}

func (e Employee) printEmployee() {
	fmt.Println("Nombre: ", e.Persona.Name)
	fmt.Println("Fecha de nacimiento: ", e.Persona.DateOfBirth)
}

func main() {
	pablo := People{1, "Pablo", "1992/1/1"}
	empleadoPablo := Employee{1, "Cajero", pablo}
	empleadoPablo.printEmployee()
}
