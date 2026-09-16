package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _MENSAJE_ERROR_COLA_VACIA = "La cola esta vacia"
const _VOLUMEN = 60000

// Funcs auxiliares
func validarColaVacia[T any](t *testing.T, cola TDACola.Cola[T]) {
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_ERROR_COLA_VACIA, func() {
		cola.VerPrimero()
	})
	require.PanicsWithValue(t, _MENSAJE_ERROR_COLA_VACIA, func() {
		cola.Desencolar()
	})
}

// recibe 2 elementos y
// 1) Crea una cola y valida que se comporte como tal
// 2) Encola elem1  y verifica que este se encuentre primero
// 3) Encola el elem2 y verifica que el primero siga siendo primero
// 4) Vacia la cola controlando el comportamiento FIFO
// 5) Controla que la cola vacia se comporte como tal
func probarFlujoBasico[T any](t *testing.T, elem1 T, elem2 T) {
	cola := TDACola.CrearColaEnlazada[T]()
	validarColaVacia(t, cola)

	cola.Encolar(elem1)
	require.False(t, cola.EstaVacia())
	require.Equal(t, elem1, cola.VerPrimero())

	cola.Encolar(elem2)
	require.Equal(t, elem1, cola.VerPrimero())

	require.Equal(t, elem1, cola.Desencolar())
	require.Equal(t, elem2, cola.VerPrimero())
	require.Equal(t, elem2, cola.Desencolar())

	validarColaVacia(t, cola)
}

// Chequea el estado inicial de una cola recién creada.
func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	validarColaVacia(t, cola)
}

// Chequea las operaciones básicas de encole y desencole con un elemento.
func TestAColarYDesaColar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 10, cola.VerPrimero())
	require.Equal(t, 10, cola.Desencolar())

	validarColaVacia(t, cola)
}

// Chequea el comportamiento de la Cola con distintos tipos de datos (int, string, float64).
func TestTiposDeDatos(t *testing.T) {
	probarFlujoBasico(t, 10, 20)
	probarFlujoBasico(t, "Hola", "Ciro!")
	probarFlujoBasico(t, 3.14, 2.71)
}

// Chequea el aColamiento en volumen verificando que el tope se actualice correctamente.
func TestAColarVolumen(t *testing.T) {
	Cola := TDACola.CrearColaEnlazada[int]()

	for i := range _VOLUMEN {
		Cola.Encolar(i)
		require.Equal(t, 0, Cola.VerPrimero())
		require.False(t, Cola.EstaVacia())
	}
}

// Chequea el desaColamiento en volumen y el orden LIFO.
func TestDesaColarVolumen(t *testing.T) {
	Cola := TDACola.CrearColaEnlazada[int]()

	for i := range _VOLUMEN {
		Cola.Encolar(i)
	}

	for j := range _VOLUMEN {
		require.Equal(t, j, Cola.VerPrimero())
		obtenido := Cola.Desencolar()
		require.Equal(t, j, obtenido)
	}

	validarColaVacia(t, Cola)
}
