package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPosicionPico(t *testing.T) {
	arreglo := []int{1, 3, 20, 4, 1, 0}

	require.Equal(t, 2, posicionPico(arreglo))
}
