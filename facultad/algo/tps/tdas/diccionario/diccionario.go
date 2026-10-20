package diccionario

type Diccionario[K comparable, V any] interface {
	// Guardar una clave y su valor asociado en el diccionario. Si la clave ya existe, se reemplaza el valor.
	Guardar(clave K, dato V)

	// Devuelve true si la clave pertenece al diccionario, false en caso contrario.
	Pertenece(clave K) bool

	// Devuelve el valor asociado a la clave. Si la clave no pertenece al diccionario, se lanza un panic.
	Obtener(clave K) V

	// Borra la clave y su valor asociado del diccionario. Si la clave no pertenece al diccionario, se lanza un panic.
	Borrar(clave K) V

	// Devuelve la cantidad de claves almacenadas en el diccionario.
	Cantidad() int

	// Itera sobre todas las claves y valores del diccionario, llamando a la función visitar para cada par clave-valor.
	Iterar(func(clave K, dato V) bool)

	// Devuelve un iterador para recorrer el diccionario.
	Iterador() IterDiccionario[K, V]
}

// IterDiccionario es un iterador para recorrer un diccionario. Permite avanzar a través de los elementos del diccionario y acceder a la clave y el valor del elemento actual.
type IterDiccionario[K comparable, V any] interface {
	// Devuelve true si hay más elementos para iterar, false en caso contrario.
	HayAlgoMas() bool

	// Devuelve la clave y el valor del elemento actual. Si no hay más elementos, se lanza un panic.
	VerActual() (K, V)

	// Avanza al siguiente elemento del diccionario. Si no hay más elementos, se lanza un panic.
	Avanzar()
}
