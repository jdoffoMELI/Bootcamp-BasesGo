package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func leerArchivo(path string) {
	defer println("Ejecucion finalizada")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Println(strings.Fields(fileScanner.Text()))
	}
}

func main() {
	filePath := "/Users/jdoffo/Desktop/Practica Bootcamp/Bootcamp-GO/Clase 05/ejercicio_02_parte2/customers.txt"
	leerArchivo(filePath)
}
