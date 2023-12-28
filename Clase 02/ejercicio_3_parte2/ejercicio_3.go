package main

import "fmt"

var meses = map[uint8]string{
	1:  "Enero",
	2:  "Febrero",
	3:  "Marzo",
	4:  "Abril",
	5:  "Mayo",
	6:  "Junio",
	7:  "Julio",
	8:  "Agosto",
	9:  "Septiembre",
	10: "Octubre",
	11: "Noviembre",
	12: "Diciembre"}

func nombreMes(numero uint8) string {
	var nombreMes, ok = meses[numero]
	if !ok {
		nombreMes = "Numero de mes invalido"
	}
	return nombreMes
}

func main() {
	var numero uint8 = 15
	fmt.Println(nombreMes(numero))
}

/*
	Existen multiples formas de resolver este ejericicio. Podemos consultar cada valor utilizando un switch
	o un if, pero este tipo de soluciones repite muchas lineas de codigo y resulta en un programa extenso.
*/
