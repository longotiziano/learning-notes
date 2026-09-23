package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	largo   int
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{
		dato, nil,
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return listaEnlazada[T]{0, nil, nil}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: l.primero, anterior: nil, lista: l}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}
