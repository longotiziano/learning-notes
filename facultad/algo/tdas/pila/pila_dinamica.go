package pila

const _MENSAJE_ERROR_PILA_VACIA = "La pila esta vacia"
const _TAMANIO_INICIAL_PILA = 10
const _VALOR_REDIMENSION = 2
const _CRITERIO_ACHICAMIENTO = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _TAMANIO_INICIAL_PILA),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(_MENSAJE_ERROR_PILA_VACIA)
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * _VALOR_REDIMENSION)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(_MENSAJE_ERROR_PILA_VACIA)
	}

	elem := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad*_CRITERIO_ACHICAMIENTO <= cap(p.datos) && cap(p.datos) > _TAMANIO_INICIAL_PILA {
		p.redimensionar(cap(p.datos) / _VALOR_REDIMENSION)
	}

	return elem
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos[:p.cantidad])
	p.datos = nuevosDatos
}
