package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() *pilaDinamica[T] {
	return &pilaDinamica[T]{make([]T, 0, 10), 0} // capacidad inicial de 10
}

// helper privado para redimensionar el arreglo
func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	redimensionado := make([]T, pila.cantidad, nuevaCapacidad)
	copy(redimensionado, pila.datos)
	pila.datos = redimensionado
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(dato T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(pila.cantidad * 2)
	}
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*4 == cap(pila.datos) {
		pila.redimensionar(pila.cantidad * 2)
	}
	dato := pila.datos[pila.cantidad]
	pila.cantidad--
	return dato

}
