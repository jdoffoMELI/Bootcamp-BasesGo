package main

import "fmt"

func empleadosMayores(empleados map[string]int) int {
	var contador int = 0
	for _, value := range empleados {
		if value >= 21 {
			contador += 1
		}
	}
	return contador
}

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Obtenemos la edad de Benjamin
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	// Cuantos empleados son mayores a 21
	fmt.Printf("Hay en total %d empleados mayores a 21 años\n", empleadosMayores(employees))

	// Agregamos un nuevo empleado a la lista
	employees["Federico"] = 25
	fmt.Println("Agregamos un nuevo empleado, Federico")
	fmt.Println(employees)

	// Eleminar a Pedro del map
	delete(employees, "Pedro")
	fmt.Println("Quitamos a Pedro del map")
	fmt.Println(employees)
}
