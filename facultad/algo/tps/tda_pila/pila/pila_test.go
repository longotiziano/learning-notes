package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Chequea que una pila recién creada esté vacía
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
}

// Chequea que al momento de apilar, la pila no esté vacía y el tope sea el elemento creado
func TestApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 10, pila.VerTope())
}

// Chequea que el desapilamiento devuelva el elemento esperado
func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.Equal(t, 10, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

// Chequea que el desapilamiento paniquee si la pila está vacía
func TestDesapilarVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

// Chequea apilamiento en volumen
func TestApilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	volumen := 10000
	for i := range volumen {
		pila.Apilar(i)
	}
	require.False(t, pila.EstaVacia())
	require.Equal(t, volumen-1, pila.VerTope())
}

// Chequea desapilamiento en volumen y comportamiento LIFO
func TestDesapilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	volumen := 10000

	for i := range volumen {
		pila.Apilar(i)
	}

	for j := range volumen {
		require.Equal(t, (volumen-1)-j, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "la pila esta vacia", func() {
		pila.Desapilar()
	})
}
