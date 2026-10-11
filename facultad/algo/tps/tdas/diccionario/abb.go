package diccionario

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

func CrearABB[K comparable, V any](funcion_cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

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
