package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERROR_LISTA_VACIA = "La lista esta vacia"
const _VOLUMEN = 150000

// Funcs auxiliares
func validarListaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {
	require.False(t, lista.EstaVacia())
	require.True(t, lista.Largo() == 0)
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		lista.BorrarPrimero()
	})
}

func validarIterVacio[T any](t *testing.T, iter TDALista.IteradorLista[T]) {
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		iter.VerActual()
	})
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		iter.Borrar()
	})
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		iter.Avanzar()
	})
}

func probarInsercion[T any](t *testing.T, lista TDALista.Lista, elem1 T, elem2 T) {
	 := TDALista.CrearListaEnlazada[T]()
	validarListaVacia(t, lista)

	lista.InsertarPrimero(elem1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, elem1, lista.VerTope())

	lista.Alistar(elem2)
	require.Equal(t, elem2, lista.VerTope())

	require.Equal(t, elem2, lista.Desalistar())
	require.Equal(t, elem1, lista.VerTope())
	require.Equal(t, elem1, lista.Desalistar())

	validarListaVacia(t, lista)
}

// Chequea el estado inicial de una lista recién creada.
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	validarListaVacia(t, lista)
}

func ValidarIterVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	validarIterVacio(t, iter)
}

func Validar

// Chequea las operaciones básicas de alistar y desalistar con un elemento.
func TestAlistarYDesalistar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Alistar(10)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 10, lista.VerTope())
	require.Equal(t, 10, lista.Desalistar())

	validarListaVacia(t, lista)
}

// Chequea el comportamiento de la lista con distintos tipos de datos (int, string, float64).
func TestTiposDeDatos(t *testing.T) {
	probarFlujoBasico(t, 10, 20)
	probarFlujoBasico(t, "Hola", "Ciro!")
	probarFlujoBasico(t, 3.14, 2.71)
}

// Chequea el alistamiento en volumen verificando que el tope se actualice correctamente.
func TestAlistarVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := 10000

	for i := range volumen {
		lista.Alistar(i)
		require.Equal(t, i, lista.VerTope())
		require.False(t, lista.EstaVacia())
	}
}

// Chequea el desalistamiento en volumen y el orden LIFO.
func TestDesalistarVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumen := 10000

	for i := range volumen {
		lista.Alistar(i)
	}

	for j := range volumen {
		esperado := (volumen - 1) - j

		require.Equal(t, esperado, lista.VerTope())
		obtenido := lista.Desalistar()
		require.Equal(t, esperado, obtenido)
	}

	validarListaVacia(t, lista)
}
