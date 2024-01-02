package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinFunc(t *testing.T) {
	var values = [4]float32{1, -22, 3, -5}
	var expected_value float32 = -22
	var result = minFunc(values[:]...)
	require.Equal(t, expected_value, result)
}

func TestMaxFunc(t *testing.T) {
	var values = [4]float32{1, -22, 3, -5}
	var expected_value float32 = 3
	var result = maxFunc(values[:]...)
	require.Equal(t, expected_value, result)
}

func TestMeanFunc(t *testing.T) {
	var values = [4]float32{2, -22, 3, -5}
	var expected_value float32 = -5.5
	var result = meanFunc(values[:]...)
	require.Equal(t, expected_value, result)
}
