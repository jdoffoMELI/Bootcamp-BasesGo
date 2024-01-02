package main

import "fmt"

type TAnimal string

const (
	perro     TAnimal = "perro"
	gato      TAnimal = "gato"
	hamnster  TAnimal = "hamnster"
	tarantula TAnimal = "tarantula"
)

func main() {

	animalDog, msg := animal(perro)
	if msg != "" {
		return
	}
	animalCat, msg := animal(gato)
	if msg != "" {
		return
	}
	var amount uint32
	amount += animalDog(10)
	amount += animalCat(10)
	fmt.Printf("Cantidad de alimento total: %.2f kg\n", float32(amount)/1000.0)

}

func animal(nombreAnimal TAnimal) (func(uint32) uint32, string) {
	var msg string
	var animalFunc func(uint32) uint32

	switch nombreAnimal {
	case perro, gato, hamnster, tarantula:
		animalFunc, msg = calcularAlimento(nombreAnimal), ""
	default:
		animalFunc, msg = nil, "El animal no existe"
	}
	return animalFunc, msg
}

func calcularAlimento(nombreAnimal TAnimal) func(uint32) uint32 {
	distribucionAlimento := map[TAnimal]uint32{
		perro:     10 * 1000,
		gato:      5 * 1000,
		hamnster:  250,
		tarantula: 150,
	}

	result_func := func(cantidadAnimal uint32) uint32 {
		return distribucionAlimento[nombreAnimal] * cantidadAnimal
	}

	return result_func
}
