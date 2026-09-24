package lista

type IteradorLista[T any] interface {
	VerActual() T
	HayAlgoMas() bool
	Avanzar()
	Insertar(T)
	Borrar() T
}

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
		panic(ErrorListaVacia)
	}
	return l.actual.dato
}

func (l *iterListaEnlazada[T]) Avanzar() {
	if !l.HayAlgoMas() {
		panic(ErrorListaVacia)
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
		panic(ErrorListaVacia)
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
