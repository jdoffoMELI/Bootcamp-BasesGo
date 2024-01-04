package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	Name        string
	ID          uint32
	PhoneNumber uint32
	Home        string
}

var (
	NameError  = errors.New("nombre de cliente invalido")
	IDError    = errors.New("ID de cliente invalida")
	PhoneError = errors.New("telefono de cliente invalido")
	HomeError  = errors.New("dirrecion de cliente invalida")
)

var Clientes []Client

func agregarCliente(c Client, filePath string) {
	/* Handler del panic */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime -", r)
		}
		fmt.Println("End of execution")
	}()

	/* Leemos el archivo */
	Clientes = cargarClientes(filePath)

	if clienteExiste(c) {
		panic("Error: client already exists")
	}
	if flag, err := clienteTieneDatosNulos(c); flag {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newLine := "\n" + fmt.Sprintf(`"%s" %d %d "%s"`, c.Name, c.ID, c.PhoneNumber, c.Home)
	_, err = file.Write([]byte(newLine))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client %s added succesfully\n", c.Name)
}

func cargarClientes(filePath string) []Client {
	var clientes []Client
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		nuevoCliente := parseStringToClient(fileScanner.Text())
		clientes = append(clientes, nuevoCliente)
	}
	return clientes
}

func parseStringToClient(s string) Client {
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ' '
	clientFields, err := r.Read()
	if err != nil {
		panic(err)
	}

	idCliente, err := strconv.Atoi(clientFields[1])
	if err != nil {
		panic(err)
	}
	numeroCliente, err := strconv.Atoi(clientFields[2])
	if err != nil {
		panic(err)
	}
	return Client{
		Name:        clientFields[0],
		ID:          uint32(idCliente),
		PhoneNumber: uint32(numeroCliente),
		Home:        clientFields[3],
	}
}

func clienteExiste(c Client) bool {
	for _, client := range Clientes {
		if client.ID == c.ID {
			return true
		}
	}
	return false
}

func clienteTieneDatosNulos(c Client) (bool, error) {
	switch {
	case c.ID == 0:
		return true, IDError
	case c.Name == "":
		return true, NameError
	case c.Home == "":
		return true, HomeError
	case c.PhoneNumber == 0:
		return true, PhoneError
	default:
		return false, nil
	}
}

func main() {
	filePath := "/Users/jdoffo/Desktop/Practica Bootcamp/Bootcamp-GO/Clase 05/ejercicio_03_parte_2/customers.txt"
	Clientes := cargarClientes(filePath)

	/* Cargamos un cliente existente para que arroje un panic */
	cliente := Clientes[0]
	agregarCliente(cliente, filePath)

	/* Cargamos un cliente nuevo con campos vacios */
	clienteCamposVacios := Client{"Pedro Sanchez", 4, 0, "Paraguay 300"}
	agregarCliente(clienteCamposVacios, filePath)

	/* Cargamos un cliente nuevo */
	clienteValido := Client{"Marco Aurelio", 4, 111, "Roma"}
	agregarCliente(clienteValido, filePath)
}
