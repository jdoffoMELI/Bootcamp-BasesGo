package main

import "fmt"

func main() {
	var humedad float32 = 82.5
	var temperatura int8 = 17
	var presionAtmosferica int16 = 1019

	fmt.Printf("Humedad: %.2f\nTemperatura: %d\nPresion Atmosferica: %d\n", humedad, temperatura, presionAtmosferica)
}
