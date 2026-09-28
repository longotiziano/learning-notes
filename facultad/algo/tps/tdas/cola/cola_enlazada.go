package cola

import (
	"reflect"
)

const _MENSAJE_ERROR_COLA_VACIA = "La cola esta vacia"

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{
		dato, nil,
	}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		nil, nil,
	}
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(_MENSAJE_ERROR_COLA_VACIA)
	}
	return c.primero.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := nodoCrear(elem)

	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
	}

	c.ultimo = nuevoNodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(_MENSAJE_ERROR_COLA_VACIA)
	}

	primerNodo := c.primero
	c.primero = primerNodo.prox

	if c.primero == nil {
		c.ultimo = nil
	}

	return primerNodo.dato
}

/*
Ejer 2 de parcial

Implementar la primitiva Colarse para el TDA ColaEnlazada que, dado un dato, lo mueva al desde su lugar (asumir
que hay solo una aparición de cada dato) hasta el inicio de la Cola, para que pueda ser usada por Hollander para mover
clientes.
Indicar y justificar la complejidad de la primitiva.
*/
func (c *colaEnlazada[T]) Colarse(dato T) {
	var anterior *nodoCola[T]
	actual := c.primero // o(1)
	for actual != nil { // o(n)
		if reflect.DeepEqual(dato, actual.dato) { // o(1)
			if actual != c.primero {
				if actual == c.ultimo {
					c.ultimo = anterior
				}
				anterior.prox = actual.prox // o(1)
				actual.prox = c.primero
				c.primero = actual
			}
			return
		}
		anterior = actual
		actual = actual.prox // o(1)
	}
}
