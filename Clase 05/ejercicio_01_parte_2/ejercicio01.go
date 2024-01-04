package main

import (
	"fmt"
	"os"
)

func leerArchivo(path string) {
	defer println("Ejecucion finalizada")
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func main() {
	leerArchivo("./customers.txt")
}
