package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularImpuesto(t *testing.T) {
	/* Caso de prueba salario < 50000 */
	var salario float32 = 40000
	var expectedResult float32 = 0
	result := calcularImpuesto(salario)
	require.Equal(t, expectedResult, result, "Verifica calculo de impuesto por debajo del umbral.")

	/* Caso de prueba 50000 <= salario < 150000 */
	salario = 75000
	expectedResult = salario * 0.17
	result = calcularImpuesto(salario)
	require.Equal(t, expectedResult, result, "Verifica calculo de impuesto para rangos intermedios.")

	/* Caso de prueba salario >= 150000 */
	salario = 180000
	expectedResult = salario * 0.27
	result = calcularImpuesto(salario)
	require.Equal(t, expectedResult, result, "Verifica calculo de impuesto para rangos mayores al umbral maximo.")
}
