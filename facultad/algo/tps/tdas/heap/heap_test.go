package cola_prioridad_test

import (
	TDAColaPrioridad "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERROR_COLA_VACIA = "La cola esta vacia"
const _VOLUMEN = 100000

func compararNums(n, m int) int {
	if n > m {
		return 1
	} else if n < m {
		return -1
	}
	return 0
}

func validarColaVacia[T any](t *testing.T, cola TDAColaPrioridad.ColaPrioridad[T]) {
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _ERROR_COLA_VACIA, func() {
		cola.VerMax()
	})
	require.PanicsWithValue(t, _ERROR_COLA_VACIA, func() {
		cola.Desencolar()
	})
	require.Equal(t, cola.Cantidad(), 0)
}

func TestCrearHeapVacio(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	validarColaVacia(t, cola)
}

func TestCrearHeapifyVacio(t *testing.T) {
	var arr []int
	cola := TDAColaPrioridad.CrearHeapArr(arr, compararNums)
	validarColaVacia(t, cola)
}

func TestHeapSort(t *testing.T) {
	var arrDesordenado []int
	arr := []int{3, 4, 1, 0, 2}
	copy(arr, arrDesordenado)
	TDAColaPrioridad.HeapSort(arr, compararNums)
	require.Equal(t, []int{0, 1, 2, 3, 4}, arr)
}

func TestHeapSortVolumen(t *testing.T) {
	var arrDesordenado []int
	for i := 0; i < _VOLUMEN; i++ {
		arrDesordenado = append(arrDesordenado, i)
	}
	arr := make([]int, _VOLUMEN)
	copy(arr, arrDesordenado)
	TDAColaPrioridad.HeapSort(arr, compararNums)
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, arr[i])
	}
}
