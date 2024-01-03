package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = make([]Product, 4)

func (p Product) save() error {
	_, err := getById(p.ID)
	if err == nil {
		return fmt.Errorf("la ID del producto ya esta en uso")
	}
	Products = append(Products[:], p)
	return nil
}

func (p Product) GetAll() {
	for _, value := range Products {
		if value.ID != 0 {
			fmt.Println(value)
		}
	}
}

func getById(id int) (Product, error) {
	var result Product
	var err error = nil
	for _, value := range Products {
		if value.ID == id {
			result = value
			return result, nil
		}
	}
	err = fmt.Errorf("no se encontro el producto")
	return result, err
}

func main() {
	Products = append(Products, Product{1, "Apple MacBook", 500, "Computadora portatil", "Laptop"})
	Products = append(Products, Product{2, "Asus 5250", 1000, "Computadora de escritorio", "Pc"})
	Products = append(Products, Product{3, "Ryzen 3000", 200, "Procesador g3", "Microprocesador"})

	newProduct := Product{3, "New Product", 10, "New Product", "Producto de prueba"}
	fmt.Println("Intento agregar un producto con ID existente")
	err := newProduct.save()
	if err != nil {
		fmt.Println(err)
	}

	// Arreglo el ID
	newProduct.ID = 4
	fmt.Println("Corrigo el ID y agrego el nuevo producto")
	err = newProduct.save()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listamos todos los productos")
	newProduct.GetAll()
}
