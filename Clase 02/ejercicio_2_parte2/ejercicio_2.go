package main

import "fmt"

func otorgarPrestamo(edad uint8, esEmpleado bool, añosAntiguedad uint8, sueldo float32) {
	switch {
	case edad <= 2:
		fmt.Println("No cumple el requisito de edad")
	case !esEmpleado:
		fmt.Println("No cumple el requisito de ser empleado")
	case añosAntiguedad <= 1:
		fmt.Println("No cumple con el requisito de atiguedad")
	default:
		if sueldo > 100000 {
			fmt.Println("¡Se le otorgara el prestamo SIN itereses!")
		} else {
			fmt.Println("¡Se le otorgara el prestamo CON interes!")
		}
	}
}

func main() {
	var edad uint8 = 23
	var esEmpleado bool = true
	var añosAntiguedad uint8 = 2
	var sueldo float32 = 1000000

	otorgarPrestamo(edad, esEmpleado, añosAntiguedad, sueldo)
}
