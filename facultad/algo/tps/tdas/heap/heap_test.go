package cola_prioridad_test

import (
	TDAColaPrioridad "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERROR_COLA_VACIA = "La cola esta vacia"
const _VOLUMEN = 100000

// Auxiliares
func compararNums(n, m int) int {
	if n > m {
		return 1
	} else if n < m {
		return -1
	}
	return 0
}

func compararStrings(s1, s2 string) int {
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
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

func validarFlujoBasico[T any](t *testing.T, cola TDAColaPrioridad.ColaPrioridad[T], arr []T) {
	for _, v := range arr {
		cola.Encolar(v)
	}
	require.Equal(t, len(arr), cola.Cantidad())
	for i := 0; i < len(arr); i++ {
		maximo := cola.VerMax()
		require.Equal(t, maximo, cola.Desencolar())
		require.Equal(t, len(arr)-i-1, cola.Cantidad())
	}
	validarColaVacia(t, cola)
}

// Tests
func TestCrearHeapVacio(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	validarColaVacia(t, cola)
}

func TestCrearHeapifyVacio(t *testing.T) {
	var arr []int
	cola := TDAColaPrioridad.CrearHeapArr(arr, compararNums)
	validarColaVacia(t, cola)
}

func TestInsertarYVerMax(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	cola.Encolar(1)
	cola.Encolar(0)
	require.Equal(t, 1, cola.VerMax())
}

func TestInsertarYDesencolar(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	cola.Encolar(1)
	cola.Encolar(0)
	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 0, cola.Desencolar())
	validarColaVacia(t, cola)
}

func TestFlujoBasicoInts(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	arr := []int{3, 4, 1, 0, 2}
	validarFlujoBasico(t, cola, arr)
}

func TestFlujoBasicoStrings(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararStrings)
	arr := []string{"c", "d", "a", "b"}
	validarFlujoBasico(t, cola, arr)
}

func TestHeapify(t *testing.T) {
	arr := []int{3, 4, 1, 0, 2}
	cola := TDAColaPrioridad.CrearHeapArr(arr, compararNums)
	for i := 4; i >= 0; i-- {
		require.Equal(t, i, cola.VerMax())
		require.Equal(t, i, cola.Desencolar())
	}
	validarColaVacia(t, cola)
}

func TestFlujoBasicoVolumen(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	var arr []int
	for i := 0; i < _VOLUMEN; i++ {
		arr = append(arr, i)
	}
	validarFlujoBasico(t, cola, arr)
}

func TestHeapifyVolumen(t *testing.T) {
	arr := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		arr[i] = i
	}
	cola := TDAColaPrioridad.CrearHeap(compararNums)
	for _, v := range arr {
		cola.Encolar(v)
	}
	colaArr := TDAColaPrioridad.CrearHeapArr(arr, compararNums)
	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, cola.VerMax())
		require.Equal(t, i, colaArr.VerMax())
		require.Equal(t, i, cola.Desencolar())
		require.Equal(t, i, colaArr.Desencolar())
	}
	validarColaVacia(t, cola)
	validarColaVacia(t, colaArr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 4, 1, 0, 2}
	TDAColaPrioridad.HeapSort(arr, compararNums)
	require.Equal(t, []int{0, 1, 2, 3, 4}, arr)
}

func TestHeapSortStrings(t *testing.T) {
	arr := []string{"c", "d", "a", "b"}
	TDAColaPrioridad.HeapSort(arr, compararStrings)
	require.Equal(t, []string{"a", "b", "c", "d"}, arr)
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
