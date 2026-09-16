package ejercicios

func esPico(v []int, pos int) bool {
	return v[pos-1] < v[pos] && v[pos] > v[pos+1]
}

// se tiene un arreglo de como minimo 3 elementos, donde es creciente hasta p, y a partir de p decreciente
// pico = punto donde o bien v[i-1] < v[i] < v[i+1] o bien inicio > final
// 0 < p < N - 1 con N el tamaño del arreglo
func posicionPico(v []int) int {
	inicio := 0
	final := len(v)
	for inicio < final {
		medio := inicio + (final-inicio)/2
		if esPico(v, medio) {
			return medio
		}
		if v[medio-1] < v[medio] {
			inicio = medio + 1
		} else {
			final = medio - 1
		}
	}
	return -1
}
