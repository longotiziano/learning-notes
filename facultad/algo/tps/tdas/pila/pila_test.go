package pila_test

import (
	ejer "tdas"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Funcs auxiliares
func validarPilaVacia[T any](t *testing.T, pila TDAPila.Pila[T]) {
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

func probarFlujoBasico[T any](t *testing.T, elem1 T, elem2 T) {
	pila := TDAPila.CrearPilaDinamica[T]()
	validarPilaVacia(t, pila)

	pila.Apilar(elem1)
	require.False(t, pila.EstaVacia())
	require.Equal(t, elem1, pila.VerTope())

	pila.Apilar(elem2)
	require.Equal(t, elem2, pila.VerTope())

	require.Equal(t, elem2, pila.Desapilar())
	require.Equal(t, elem1, pila.VerTope())
	require.Equal(t, elem1, pila.Desapilar())

	validarPilaVacia(t, pila)
}

// Chequea el estado inicial de una pila recién creada.
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	validarPilaVacia(t, pila)
}

// Chequea las operaciones básicas de apilar y desapilar con un elemento.
func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)

	require.False(t, pila.EstaVacia())
	require.Equal(t, 10, pila.VerTope())
	require.Equal(t, 10, pila.Desapilar())

	validarPilaVacia(t, pila)
}

// Chequea el comportamiento de la pila con distintos tipos de datos (int, string, float64).
func TestTiposDeDatos(t *testing.T) {
	probarFlujoBasico(t, 10, 20)
	probarFlujoBasico(t, "Hola", "Ciro!")
	probarFlujoBasico(t, 3.14, 2.71)
}

// Chequea el apilamiento en volumen verificando que el tope se actualice correctamente.
func TestApilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	volumen := 10000

	for i := range volumen {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
		require.False(t, pila.EstaVacia())
	}
}

// Chequea el desapilamiento en volumen y el orden LIFO.
func TestDesapilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	volumen := 10000

	for i := range volumen {
		pila.Apilar(i)
	}

	for j := range volumen {
		esperado := (volumen - 1) - j

		require.Equal(t, esperado, pila.VerTope())
		obtenido := pila.Desapilar()
		require.Equal(t, esperado, obtenido)
	}

	validarPilaVacia(t, pila)
}

func TestOrdenarPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(4)
	pila.Apilar(1)
	pila.Apilar(5)
	pila.Apilar(2)
	pila.Apilar(3)

	// Llamás a tu función usando el nombre del paquete ejer
	ejer.OrdenarPila(pila)

	require.Equal(t, 5, pila.Desapilar())
	require.Equal(t, 4, pila.Desapilar())
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())

	validarPilaVacia(t, pila)
}
