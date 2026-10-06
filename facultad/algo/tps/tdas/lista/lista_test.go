package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERROR_LISTA_VACIA = "La lista esta vacia"
const _VOLUMEN = 150000

func validarListaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
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

func insertarPrimerElemento[T any](t *testing.T, lista TDALista.Lista[T], elem T) {
	largoPrevio := lista.Largo()
	var ultimo T
	if !lista.EstaVacia() {
		ultimo = lista.VerUltimo()
	}

	lista.InsertarPrimero(elem)

	require.False(t, lista.EstaVacia())
	require.Equal(t, largoPrevio+1, lista.Largo())
	require.Equal(t, elem, lista.VerPrimero())

	if largoPrevio > 0 {
		require.Equal(t, ultimo, lista.VerUltimo())
	} else {
		require.Equal(t, elem, lista.VerUltimo())
	}
}

func insertarUltimoElemento[T any](t *testing.T, lista TDALista.Lista[T], elem T) {
	largoPrevio := lista.Largo()
	var primero T
	if !lista.EstaVacia() {
		primero = lista.VerPrimero()
	}

	lista.InsertarUltimo(elem)

	require.False(t, lista.EstaVacia())
	require.Equal(t, largoPrevio+1, lista.Largo())
	require.Equal(t, elem, lista.VerUltimo())

	if largoPrevio > 0 {
		require.Equal(t, primero, lista.VerPrimero())
	}
}

func insertarVariosPrimeros(t *testing.T, lista TDALista.Lista[int], cantidad int) {
	for i := 0; i < cantidad; i++ {
		insertarPrimerElemento(t, lista, i)
	}
}

func insertarVariosUltimos(t *testing.T, lista TDALista.Lista[int], cantidad int) {
	for i := 0; i < cantidad; i++ {
		insertarUltimoElemento(t, lista, i)
	}
}

func borrarPrimerElemento[T any](t *testing.T, lista TDALista.Lista[T]) {
	primero := lista.VerPrimero()
	ultimo := lista.VerUltimo()
	largoPrevio := lista.Largo()
	borrado := lista.BorrarPrimero()
	require.Equal(t, primero, borrado)
	require.Equal(t, largoPrevio-1, lista.Largo())
	if largoPrevio > 1 {
		require.Equal(t, ultimo, lista.VerUltimo())
	}
}

func obtenerElementos[T any](lista TDALista.Lista[T]) []T {
	var elementos []T
	lista.Iterar(func(dato T) bool {
		elementos = append(elementos, dato)
		return true
	})
	return elementos
}

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	validarListaVacia(t, lista)
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarPrimerElemento(t, lista, 10)
}

func TestInsertarVariosPrimeros(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosPrimeros(t, lista, 3)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, []int{2, 1, 0}, obtenerElementos(lista))
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoElemento(t, lista, 10)
}

func TestInsertarVariosUltimos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, []int{0, 1, 2}, obtenerElementos(lista))
}

func TestInsertarPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarPrimerElemento(t, lista, 20)
	insertarPrimerElemento(t, lista, 10)
	insertarUltimoElemento(t, lista, 30)
	insertarUltimoElemento(t, lista, 40)
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 40, lista.VerUltimo())
	require.Equal(t, []int{10, 20, 30, 40}, obtenerElementos(lista))
}

func TestBorrarUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarUltimoElemento(t, lista, 10)
	require.Equal(t, 10, lista.BorrarPrimero())
	validarListaVacia(t, lista)
}

func TestBorrarPrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 2)
	borrarPrimerElemento(t, lista)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestBorrarTodosLosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	require.Equal(t, 0, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	validarListaVacia(t, lista)
}

func TestVolumenBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, _VOLUMEN)

	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	validarListaVacia(t, lista)
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cantidad := 0
	lista.Iterar(func(dato int) bool {
		cantidad++
		return true
	})
	require.Equal(t, 0, cantidad)
}

func TestIterarTodaLaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	require.Equal(t, []int{0, 1, 2}, obtenerElementos(lista))
}

func TestIterarCortaCuandoVisitarDevuelveFalse(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 4)
	var elementos []int
	lista.Iterar(func(dato int) bool {
		elementos = append(elementos, dato)
		return dato != 1
	})
	require.Equal(t, []int{0, 1}, elementos)
}

func TestIterarCortaEnPrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	var elementos []int
	lista.Iterar(func(dato int) bool {
		elementos = append(elementos, dato)
		return false
	})
	require.Equal(t, []int{0}, elementos)
}

func TestVolumenIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, _VOLUMEN)

	esperado := 0
	lista.Iterar(func(dato int) bool {
		require.Equal(t, esperado, dato)
		esperado++
		return true
	})

	require.Equal(t, _VOLUMEN, esperado)
}

func TestIterExtVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	validarIterVacio(t, iter)
}

func TestIterExtVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	iter := lista.Iterador()
	require.Equal(t, 0, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 1, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 2, iter.VerActual())
	iter.Avanzar()
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		iter.VerActual()
	})
}

func TestIterExtAvanzar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 2)
	iter := lista.Iterador()
	require.Equal(t, 0, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 1, iter.VerActual())
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, _ERROR_LISTA_VACIA, func() {
		iter.Avanzar()
	})
}

func TestIterExtInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 2)
	iter := lista.Iterador()
	iter.Insertar(10)
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 10, iter.VerActual())
	require.Equal(t, []int{10, 0, 1}, obtenerElementos(lista))
}

func TestIterExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 2)
	iter := lista.Iterador()
	iter.Avanzar()
	iter.Insertar(10)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 10, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 1, iter.VerActual())
	require.Equal(t, []int{0, 10, 1}, obtenerElementos(lista))
}

func TestIterExtInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 1)
	iter := lista.Iterador()
	iter.Avanzar()
	iter.Insertar(10)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 10, iter.VerActual())
	require.Equal(t, []int{0, 10}, obtenerElementos(lista))
}

func TestIterExtInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(10)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 10, iter.VerActual())
}

func TestIterExtBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	iter := lista.Iterador()
	require.Equal(t, 0, iter.Borrar())
	require.Equal(t, 1, iter.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, []int{1, 2}, obtenerElementos(lista))
}

func TestIterExtBorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	iter := lista.Iterador()
	iter.Avanzar()
	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 2, iter.VerActual())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, []int{0, 2}, obtenerElementos(lista))
}

func TestIterExtBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 3)
	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()
	require.Equal(t, 2, iter.Borrar())
	require.False(t, iter.HayAlgoMas())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, []int{0, 1}, obtenerElementos(lista))
}

func TestIterExtBorrarUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 1)
	iter := lista.Iterador()
	require.Equal(t, 0, iter.Borrar())
	validarListaVacia(t, lista)
	require.False(t, iter.HayAlgoMas())
}

func TestIterExtBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	validarIterVacio(t, iter)
}

func TestIterExtInsertarYBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, 2)
	iter := lista.Iterador()
	iter.Avanzar()
	iter.Insertar(10)
	require.Equal(t, []int{0, 10, 1}, obtenerElementos(lista))
	require.Equal(t, 10, iter.Borrar())
	require.Equal(t, []int{0, 1}, obtenerElementos(lista))
	require.Equal(t, 1, iter.VerActual())
}

func TestVolumenIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	insertarVariosUltimos(t, lista, _VOLUMEN)

	iter := lista.Iterador()
	esperado := 0

	for iter.HayAlgoMas() {
		require.Equal(t, esperado, iter.VerActual())
		esperado++
		iter.Avanzar()
	}

	require.Equal(t, _VOLUMEN, esperado)
}
