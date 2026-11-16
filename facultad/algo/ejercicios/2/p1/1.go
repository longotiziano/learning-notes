package ejercicios

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (ab *ab) sumaHijosIzquierdos() int {
	return aux(ab, false)
}

func aux(ab *ab, esIzq bool) int {
	if ab == nil {
		return 0
	}
	aumento := 0
	if esIzq {
		aumento = ab.dato
	}
	sumaIzq := aux(ab.izq, true)
	sumaDer := aux(ab.der, false)
	return aumento + sumaIzq + sumaDer
}
