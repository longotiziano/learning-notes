package diccionario

import TDAPila "tdas/pila"

const _ERROR_CLAVE_INEXISTENTE = "La clave no pertenece al diccionario"

// si 1era clave < 2da, devuelve < 0
type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iteradorAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   funcCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

// devuelve un puntero al puntero del nodo que contiene la clave, 
// o un puntero al puntero donde debería estar el nodo con la clave
func (abb *abb[K, V]) obtenerEnlace(nodo **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	if *nodo == nil || abb.cmp(clave, (*nodo).clave) == 0 {
		return nodo
	}
	if abb.cmp(clave, (*nodo).clave) < 0 {
		return abb.obtenerEnlace(&((*nodo).izquierdo), clave)
	}
	return abb.obtenerEnlace(&((*nodo).derecho), clave)
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo := abb.obtenerEnlace(&abb.raiz, clave)
	if *nodo == nil {
		abb.cantidad++
		*nodo = crearNodoABB(clave, dato)
	} else {
		(*nodo).dato = dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *abb.obtenerEnlace(&abb.raiz, clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.obtenerEnlace(&abb.raiz, clave)
	if *nodo != nil {
		return (*nodo).dato
	}
	panic(_ERROR_CLAVE_INEXISTENTE)
}

// borra el nodo con la clave y devuelve el dato que contenía. Si no existe, lanza un panic.
// maneja los 3 casos: nodo hoja, nodo con un hijo, nodo con dos hijos
func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := abb.obtenerEnlace(&abb.raiz, clave)
	if *nodo == nil {
		panic(_ERROR_CLAVE_INEXISTENTE)
	}
	dato := (*nodo).dato

	if (*nodo).derecho == nil && (*nodo).izquierdo == nil {
		*nodo = nil

	} else if (*nodo).derecho != nil && (*nodo).izquierdo != nil {
		izquierdo := (*nodo).izquierdo
		derecho := (*nodo).derecho

		enlace := &izquierdo
		for (*enlace).derecho != nil {
			enlace = &(*enlace).derecho
		}

		nodoARemplazar := *enlace
		subarbolIzquierdo := nodoARemplazar.izquierdo

		*enlace = subarbolIzquierdo

		if nodoARemplazar == izquierdo {
			nodoARemplazar.izquierdo = subarbolIzquierdo
		} else {
			nodoARemplazar.izquierdo = izquierdo
		}

		nodoARemplazar.derecho = derecho
		*nodo = nodoARemplazar

	} else if (*nodo).derecho != nil {
		*nodo = (*nodo).derecho
	} else {
		*nodo = (*nodo).izquierdo
	}

	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.raiz.iterar(visitar)
}

// itero in-order
func (nodo *nodoAbb[K, V]) iterar(visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}
	if !nodo.izquierdo.iterar(visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	return nodo.derecho.iterar(visitar)
}

// iterador interno

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(K, V) bool, cmp funcCmp[K]) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || cmp(nodo.clave, *desde) > 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	enRangoDesde := desde == nil || cmp(nodo.clave, *desde) >= 0
	enRangoHasta := hasta == nil || cmp(nodo.clave, *hasta) <= 0

	if enRangoDesde && enRangoHasta {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || cmp(nodo.clave, *hasta) < 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	return true
}

// iterador externo

func (iter *iteradorAbb[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.derecho
		} else {
			iter.pila.Apilar(nodo)
			nodo = nodo.izquierdo
		}
	}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iteradorAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		cmp:   abb.cmp,
	}
	iter.apilarIzquierdos(abb.raiz)
	return iter
}

func (iter *iteradorAbb[K, V]) HayAlgoMas() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	if iter.hasta != nil && iter.cmp(iter.pila.VerTope().clave, *iter.hasta) > 0 {
		return false
	}
	return true
}

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iteradorAbb[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	desapilado := iter.pila.Desapilar()
	iter.apilarIzquierdos(desapilado.derecho)
}
