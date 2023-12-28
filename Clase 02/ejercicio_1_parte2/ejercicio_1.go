package main

import "fmt"

func imprimirPalabra(palabra string) {
	largo_palabra := len(palabra)
	fmt.Println("Largo de la palabra: ", largo_palabra)
	fmt.Println("Sus letras son:")
	for _, char := range palabra {
		fmt.Println(string(char))
	}
}

func main() {
	var palabraEntrada string
	fmt.Scanf("%s", &palabraEntrada)
	imprimirPalabra(palabraEntrada)
}
