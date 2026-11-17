package practica

type cmpFunc[T any] func(x, y T) int

type abb[T any] struct {
	izq  *abb[T]
	der  *abb[T]
	dato T
	cmp  cmpFunc[T]
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (abb *abb[T]) EsAvl() bool {
	if abb == nil {
		return true
	}
	var esAvl bool
	auxEsAvl(abb, &esAvl)
	return esAvl
}

// recibe árboles no nulos
// si cmp > 0, entonces los params van en orden ascendente
func auxEsABB[T any](abb *abb[T], esAvl *bool) {
	*esAvl = true
	if abb.izq != nil && !(abb.cmp(abb.izq.dato, abb.dato) > 0) {
		*esAvl = false
		return
	}
	if abb.der != nil && !(abb.cmp(abb.dato, abb.der.dato) > 0) {
		*esAvl = false
		return
	}
}

func auxEsAvl[T any](abb *abb[T], esAvl *bool) int { // O(N)
	if abb == nil {
		return 0
	}
	auxEsABB(abb, esAvl)
	if *esAvl {
		return 0
	}
	altIzq := auxEsAvl(abb.izq, esAvl)
	altDer := auxEsAvl(abb.der, esAvl)
	difAltura := altIzq - altDer
	if absInt(difAltura) > 1 {
		*esAvl = false
		return 0
	}
	return 1 + max(altIzq, altDer)
}
