package lista

type Lista[T any] interface {
	// Devuelve si la lista esta vacia
	EstaVacia() bool

	// Inserta un elemento al principio de la lista
	InsertarPrimero(T)

	// Inserta un elemento al final de la lista
	InsertarUltimo(T)

	// Borra el primer elemento de la lista y lo devuelve
	BorrarPrimero() T

	// Devuelve el primer elemento de la lista sin borrarlo
	VerPrimero() T

	// Devuelve el último elemento de la lista sin borrarlo
	VerUltimo() T

	// Devuelve la cantidad de elementos en la lista
	Largo() int

	// Recorre la lista y aplica la función de visita a cada elemento
	Iterar(visitar func(T) bool)

	// Devuelve un iterador para recorrer la lista
	Iterador() IteradorLista[T]
}

// IteradorLista es un iterador para recorrer la lista
type IteradorLista[T any] interface {
	// Devuelve el elemento actual del iterador
	VerActual() T

	// Devuelve si hay más elementos en la lista
	HayAlgoMas() bool

	// Avanza el iterador al siguiente elemento
	Avanzar()

	// Inserta un elemento en la posición actual del iterador
	Insertar(T)

	// Borra el elemento actual del iterador y lo devuelve
	Borrar() T
}
