package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGato(t *testing.T) {
	animalFunc, msg := animal(gato)
	require.Equal(t, "", msg)
	var cantidadAnimales uint32 = 15
	var expectedValue uint32 = cantidadAnimales * 5000
	var result uint32 = animalFunc(cantidadAnimales)
	require.Equal(t, expectedValue, result)
}

func TestPerro(t *testing.T) {
	animalFunc, msg := animal(perro)
	require.Equal(t, "", msg)
	var cantidadAnimales uint32 = 25
	var expectedValue uint32 = cantidadAnimales * 10000
	var result uint32 = animalFunc(cantidadAnimales)
	require.Equal(t, expectedValue, result)
}

func TestHamnster(t *testing.T) {
	animalFunc, msg := animal(hamnster)
	require.Equal(t, "", msg)
	var cantidadAnimales uint32 = 99
	var expectedValue uint32 = cantidadAnimales * 250
	var result uint32 = animalFunc(cantidadAnimales)
	require.Equal(t, expectedValue, result)
}

func TestTarantula(t *testing.T) {
	animalFunc, msg := animal(tarantula)
	require.Equal(t, "", msg)
	var cantidadAnimales uint32 = 11
	var expectedValue uint32 = cantidadAnimales * 150
	var result uint32 = animalFunc(cantidadAnimales)
	require.Equal(t, expectedValue, result)
}
