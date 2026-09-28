package ejercicios

/*
Debido a la ingerencia de rayos cósmicos, la integridad de esta estructura enlazada puede verse afectada. Estos rayos
pueden generar que el siguiente de cada NodoLista se vea afectado y en vez de apuntar al nodo siguiente que le
corresponde, apuntar a otro anterior, generando ciclos.

Se pide implementar una primitiva que detecte si la lista tiene un ciclo. La función debe tener una complejidad temporal
de O(n) (siendo n la cantidad de elementos de la lista) y espacial de O(1).

Pista: pensar en el ejercicio (y resolución) para obtener el K-último de una lista sin tener largo, en una única iteración.
*/

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func (l *ListaEnlazada[T]) TieneCiclosCortos() bool {
	res := false
	var anterior *nodoLista[T]
	actual := l.primero
	for actual != nil && !res {
		// me resguardo el caso de que haya un único elemento
		if anterior != nil {
			// significan que el nodo apunta a la misma dirección de memoria,
			// por lo que son el mismo nodo (siguiente = anterior)
			if anterior == actual.siguiente {
				res = true
			}
		}
		anterior = actual
		actual = actual.siguiente
	}
	return res
}
