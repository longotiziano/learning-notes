package lista

const (
	ErrorListaVacia = "La lista esta vacia"
)

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato, nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)
	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo

	if lista.ultimo == nil {
		lista.ultimo = nuevoNodo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}

	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(ErrorListaVacia)
	}

	dato := lista.primero.dato

	lista.primero = lista.primero.siguiente
	lista.largo--
	if lista.EstaVacia() {
		lista.ultimo = nil
	}

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(ErrorListaVacia)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(ErrorListaVacia)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	datoActual := lista.primero

	for datoActual != nil && visitar(datoActual.dato) {
		datoActual = datoActual.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: l.primero, anterior: nil, lista: l}
}
