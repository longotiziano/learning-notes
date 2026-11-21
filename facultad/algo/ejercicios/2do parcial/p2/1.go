package ejercicios

type ab[T any] struct {
	izq  *ab[T]
	der  *ab[T]
	dato T
}

func (ab *ab[T]) DistanciaD() int {
	return aux(ab, 0)
}

func aux[T any](ab *ab[T], altura int) int {
	if ab == nil {
		return altura
	}
	return min(aux(ab.izq, altura+1), aux(ab.der, altura+1))
}
