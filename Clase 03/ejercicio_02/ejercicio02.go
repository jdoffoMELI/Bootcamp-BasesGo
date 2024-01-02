package main

import "fmt"

func main() {
	notas := [5]float32{9.5, 6, 2, 4, 5}
	fmt.Println("Notas: ", notas)
	fmt.Println("Promedio: ", promediarNotas(notas[:]...))
}

func promediarNotas(notas ...float32) float32 {
	var suma float32 = 0
	for _, value := range notas {
		suma += value
	}

	return suma / float32(len(notas))
}
