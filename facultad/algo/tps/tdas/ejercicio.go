package ejer

import (
	"cmp"
	TDAPila "tdas/pila"
)

// (★★★★) ♠♠ Implementar una función que ordene de manera ascendente una pila de enteros sin conocer
// su estructura interna y utilizando como estructura auxiliar sólo otra pila auxiliar.
// Por ejemplo, la pila [ 4, 1, 5, 2, 3 ] debe quedar como [ 1, 2, 3, 4, 5 ]
// (siendo el último elemento el tope de la pila, en ambos casos).
// Indicar y justificar el orden de la función.

// mi idea es intentar siempre dejar el maximo hacia abajo. En caso de encontrar un nuevo maximo, entonces se guarda en la variable auxiliar,
// se vacia la pila auxiliar, se pone ese maximo, y se vuelve a comparar
// queremos que la pilaAux quede como [ 5 4 3 2 1 ]
func OrdenarPila[T cmp.Ordered](pila TDAPila.Pila[T]) {
	pilaAux := TDAPila.CrearPilaDinamica[T]()

	// queremos que la pilaAux quede como [ 5 4 3 2 1 ]
	for !pila.EstaVacia() {
		desapilado := pila.Desapilar() // el elemento que estamos evaluando ya no esta en la pila
		// mientras el desapilado sea mayor al tope de la pila auxiliar, movemos elementos de una pila a otra o esta q la auxiliar este vacia
		// este ultimo caso significaria que el desapilado
		for !pilaAux.EstaVacia() {
			if desapilado > pilaAux.VerTope() {
				pila.Apilar(pilaAux.Desapilar())
			}
		}
		pilaAux.Apilar(desapilado)
	}
	for !pilaAux.EstaVacia() {
		pila.Apilar(pilaAux.Desapilar())
	}
}
