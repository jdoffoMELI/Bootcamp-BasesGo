package main

import "fmt"

func main() {
	notas := [5]float32{9.5, 6, 2, 4, 5}
	promedio, err := promediarNotas(notas[:]...)
	if err != nil {
		fmt.Println("Notas: ", notas)
		fmt.Println("Promedio: ", promedio)
	}
}

func promediarNotas(notas ...float32) (float32, error) {
	var suma float32 = 0

	if len(notas) == 0 {
		return 0, nil
	}

	for _, value := range notas {
		if value < 0 {
			return 0, fmt.Errorf("no se pueden introducir valores negativos")
		}
		suma += value
	}

	return suma / float32(len(notas)), nil
}
