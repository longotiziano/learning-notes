package lista

const _ERROR_LISTA_VACIA = "La lista esta vacia"

type IteradorLista[T any] interface {
	VerActual() T
	HayAlgoMas() bool
	Avanzar()
	Insertar(T)
	Borrar() T
}

type listaIter[T any] struct {
	actual   *listaNodo[T]
	anterior *listaNodo[T]
	lista    Lista[T]
}

func (l *listaIter[T]) HayAlgoMas() bool {
	return l.actual != nil
}

func (l *listaIter[T]) VerActual() T {
	return l.actual.dato
}

func (l *listaIter[T]) Avanzar() {
	if !l.HayAlgoMas() {
		panic(_ERROR_LISTA_VACIA)
	}
	l.anterior = l.actual
	l.actual = l.actual.siguiente
}

// nodo nuevo -> anterior -> nodo nuevo -> actual
func (l *listaIter[T]) Insertar(dato T) {
	nodoNuevo := nodoCrear(dato)
	nodoNuevo.siguiente = l.actual
	l.actual = nodoNuevo
	if l.anterior != nil {
		l.anterior.siguiente = nodoNuevo
	}
	lista.Insertar()
}

// anterior -> actual -> siguiente | anterior -----------> siguiente
func (l *listaIter[T]) Borrar() T {
	actual := l.actual
	l.anterior.siguiente = actual.siguiente

	return actual.dato
}
