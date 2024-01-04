package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horasTrabajadas int32, valorHora float32) (float32, error) {
	if horasTrabajadas < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	var salario float32 = float32(horasTrabajadas) * valorHora
	if salario > 150000 {
		salario *= 0.9
	}
	return salario, nil
}

func main() {
	var salario, valorHora float32
	var horasTrabajadas int32
	fmt.Println("Insert your montly hours and their value")
	fmt.Scanf("%d %f", &horasTrabajadas, &valorHora)
	salario, err := calcularSalario(horasTrabajadas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Su salario es: ", salario)
}
