package ejercicios

/*
Se quiere implementar un nuevo TDA Pila de Prioridad que maneje 2 tipos de prioridades para sus elementos (“prioritario”
y “no prioritario”).

Las primitivas deben ser EstaVacia, VerTope, Desapilar, ApilarPrioritario y ApilarNoPrioritario.

Al Desapilar, si hay algún elemento prioritario, este debe salir siempre antes que un elemento no prioritario (y VerTope debe mostrar
el siguiente a ser desapilado dado el estado actual).

Dentro de la misma prioridad, los elementos deben cumplir con la propiedad
LIFO. Implementar este TDA (incluyendo estructura interna) de tal forma que todas las primitivas funcionen en tiempo constante.
*/

type pilaDePrioridad[T any] struct {
	datosPrioridad      []T
	datosNoPrioridad    []T
	cantidadPrioridad   int
	cantidadNoPrioridad int
}

const _ERROR_PILA_VACIA = "La pila esta vacia"

func (p *pilaDePrioridad[T]) noHayPrioritarios() bool {
	return p.cantidadPrioridad == 0
}

func (p *pilaDePrioridad[T]) EstaVacia() bool {
	return p.noHayPrioritarios() && p.cantidadNoPrioridad == 0
}

func (p *pilaDePrioridad[T]) VerTope() T {
	if p.EstaVacia() {
		panic(_ERROR_PILA_VACIA)
	}
	if p.noHayPrioritarios() {
		return p.datosNoPrioridad[p.cantidadNoPrioridad-1]
	}
	return p.datosPrioridad[p.cantidadPrioridad-1]
}

func (p *pilaDePrioridad[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(_ERROR_PILA_VACIA)
	}
	if p.noHayPrioritarios() {
		dato := p.datosNoPrioridad[p.cantidadNoPrioridad-1]
		p.cantidadNoPrioridad--
		return dato
	}
	dato := p.datosPrioridad[p.cantidadPrioridad-1]
	p.cantidadPrioridad--
	return dato
}

func (p *pilaDePrioridad[T]) ApilarPrioritario(dato T) {
	p.datosPrioridad[p.cantidadPrioridad] = dato
	p.cantidadPrioridad++
}

func (p *pilaDePrioridad[T]) ApilarNoPrioritario(dato T) {
	p.datosNoPrioridad[p.cantidadNoPrioridad] = dato
	p.cantidadNoPrioridad++
}
