package main

import "fmt"

type Categoria uint8

const (
	A Categoria = iota
	B
	C
)

func main() {
	var categoria Categoria = A
	var minutosTrabajados uint32 = 60
	fmt.Printf("Salario calculado: $%.2f\n", calcularSalario(categoria, minutosTrabajados))

}

func calcularSalario(categoria Categoria, minutosTrabajados uint32) float32 {
	var salario float32 = 0.0
	var valorPorHora float32 = 0
	var adicionalMensual float32 = 0.0
	var horasTrabajadas float32 = float32(minutosTrabajados / 60)
	switch categoria {
	case A:
		valorPorHora = 3000
		adicionalMensual = 0.5
	case B:
		valorPorHora = 1500
		adicionalMensual = 0.2
	case C:
		valorPorHora = 1000
		adicionalMensual = 0.0
	default:
		return 0
	}
	salario = horasTrabajadas * float32(valorPorHora)
	salario += salario * adicionalMensual
	return salario
}
