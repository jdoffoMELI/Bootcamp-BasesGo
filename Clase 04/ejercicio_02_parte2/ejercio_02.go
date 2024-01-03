package main

import "fmt"

/* Tipos de produtos validos */
type ProductType uint32

const (
	TSmall ProductType = iota
	TMedium
	TLarge
)

type Product interface {
	Price() float64
}

/* Implementacion del tipo Small */
type Small struct {
	Precio float64
}

func (s Small) Price() float64 {
	return s.Precio
}

/* Implementacion del tipo Medium */
type Medium struct {
	Precio float64
}

func (m Medium) Price() float64 {
	costoAcopio := m.Precio * 0.25
	return m.Precio + costoAcopio*0.03
}

/* Implementacion del tipo Large */
type Large struct {
	Precio float64
}

func (l Large) Price() float64 {
	costoAcopio := l.Precio * 0.35
	costoEnvio := 2500.0
	return l.Precio + costoAcopio*0.06 + costoEnvio
}

func factory(tipoProducto ProductType, precio float64) (Product, error) {
	switch tipoProducto {
	case TSmall:
		return Small{precio}, nil
	case TMedium:
		return Medium{precio}, nil
	case TLarge:
		return Large{precio}, nil
	default:
		return nil, fmt.Errorf("tipo de producto invalido")
	}

}

func main() {
	smallProduct, err := factory(TSmall, 100)
	if err != nil {
		println(err)
	}
	mediumProduct, err := factory(TMedium, 500)
	if err != nil {
		println(err)
	}
	largeProduct, err := factory(TLarge, 1000)
	if err != nil {
		println(err)
	}
	fmt.Println("Precio del producto pequeño: ", smallProduct.Price())
	fmt.Println("Precio del producto mediano: ", mediumProduct.Price())
	fmt.Println("Precio del producto grande:  ", largeProduct.Price())
}
