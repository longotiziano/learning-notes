package pila

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, 10),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * 2)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	elem := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad*4 <= cap(p.datos) && cap(p.datos) > 10 {
		p.redimensionar(cap(p.datos) / 2)
	}

	return elem
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos[:p.cantidad])
	p.datos = nuevosDatos
}
