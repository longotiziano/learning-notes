package era

// dado un arreglo ordenado y sin elementos repetidos de valores enteros positivos,
// obtener el minimo valor que no se encuentre en el arreglo
// obtener en tiempo logaritmico
func minimoExcluido(slice []int, inicio int, final int) int {
	largo := len(slice)

	if largo == 0 {
		return 0
	} else if largo-1 == slice[largo-1] {
		return largo
	}

	medio := inicio + (final-inicio)/2

	for i := 0; i < largo; i++ {

		// ej 100 / 2 = 50 || slice[50] = 60 -> 50 > 60
		if medio == slice[medio] {
			// busco en la segunda mitad
			inicio = medio + 1
		} else if medio < slice[medio] {
			// si la posicion 50 es menor al elemento en slice[50], entonces hay numeros que se estan salteando
			final = medio - 1
		}
	}
}
