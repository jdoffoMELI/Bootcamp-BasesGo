package main

import "fmt"

func main() {
	var salarioEmpleado float32 = 1000000
	fmt.Printf("Impuesto $%.2f\n", calcularImpuesto(salarioEmpleado))
}

func calcularImpuesto(salario float32) float32 {
	var porcentaje float32 = 0.0
	if salario > 50000 {
		porcentaje += 0.17
	}
	if salario > 150000 {
		porcentaje += 0.1
	}
	return salario * porcentaje
}
