package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularSalario_CategoriaA(t *testing.T) {
	var minutosTrabajados float32 = 150000
	var expected_value float32 = minutosTrabajados / 60 * 3000
	expected_value += expected_value * 0.5
	result := calcularSalario(A, uint32(minutosTrabajados))
	require.Equal(t, expected_value, result)
}

func TestCalcularSalario_CategoriaB(t *testing.T) {
	var minutosTrabajados float32 = 290000
	var expected_value float32 = minutosTrabajados / 60 * 15000
	expected_value += expected_value * 0.2
	result := calcularSalario(A, uint32(minutosTrabajados))
	require.Equal(t, expected_value, result)
}
func TestCalcularSalario_CategoriaC(t *testing.T) {
	var minutosTrabajados float32 = 160020
	var expected_value float32 = minutosTrabajados / 60 * 1000
	result := calcularSalario(A, uint32(minutosTrabajados))
	require.Equal(t, expected_value, result)
}
