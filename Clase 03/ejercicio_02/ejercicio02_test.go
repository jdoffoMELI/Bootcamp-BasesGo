package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPromediarNotas_SliceVacia(t *testing.T) {
	var notas = []float32{}
	var expected_value float32 = 0

	result, err := promediarNotas(notas[:]...)
	require.Nil(t, err)
	require.Equal(t, expected_value, result)
}

func TestPromediarNotas_SliceValoresNegativos(t *testing.T) {
	var notas = [3]float32{1, 0, -1}
	var expected_value float32 = 0
	result, err := promediarNotas(notas[:]...)
	require.NotNil(t, err)
	require.Equal(t, expected_value, result)
}

func TestPromediarNotas_SliceValida(t *testing.T) {
	var notas = [5]float32{10, 10, 10, 10, 10}
	var expected_value float32 = 10
	result, err := promediarNotas(notas[:]...)
	require.Nil(t, err)
	require.Equal(t, expected_value, result)
}
