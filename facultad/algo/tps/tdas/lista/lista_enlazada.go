package lista

const _ERROR_LISTA_VACIA = "La lista esta vacia"
const _ERROR_ITERADOR_TERMINO = "El iterador termino de iterar"

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
	estabaVacia := lista.EstaVacia()

	nuevoNodo := crearNodo(dato)
	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo

	if estabaVacia {
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
		panic(_ERROR_LISTA_VACIA)
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
		panic(_ERROR_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_ERROR_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for datoActual := lista.primero; datoActual != nil && visitar(datoActual.dato); datoActual = datoActual.siguiente {
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: l.primero, anterior: nil, lista: l}
}

// ITERADOR
type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (l *iterListaEnlazada[T]) HayAlgoMas() bool {
	return l.actual != nil
}

func (l *iterListaEnlazada[T]) VerActual() T {
	if !l.HayAlgoMas() {
		panic(_ERROR_ITERADOR_TERMINO)
	}
	return l.actual.dato
}

func (l *iterListaEnlazada[T]) Avanzar() {
	if !l.HayAlgoMas() {
		panic(_ERROR_ITERADOR_TERMINO)
	}

	l.anterior = l.actual
	l.actual = l.actual.siguiente
}

func (l *iterListaEnlazada[T]) Insertar(dato T) {
	nodoNuevo := crearNodo(dato)

	if l.anterior == nil {
		l.lista.primero = nodoNuevo
	} else {
		l.anterior.siguiente = nodoNuevo
	}

	nodoNuevo.siguiente = l.actual
	l.actual = nodoNuevo

	if nodoNuevo.siguiente == nil {
		l.lista.ultimo = nodoNuevo
	}

	l.lista.largo++
}

func (l *iterListaEnlazada[T]) Borrar() T {
	if !l.HayAlgoMas() {
		panic(_ERROR_ITERADOR_TERMINO)
	}

	dato := l.actual.dato

	if l.anterior == nil {
		l.lista.primero = l.actual.siguiente
	} else {
		l.anterior.siguiente = l.actual.siguiente
	}

	if l.actual == l.lista.ultimo {
		l.lista.ultimo = l.anterior
	}

	l.actual = l.actual.siguiente
	l.lista.largo--

	return dato
}
