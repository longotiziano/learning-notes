package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*y, *x = *x, *y
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	var maxInd int = -1
	if (len(vector)) > 0 {
		maxInd = 0
	}
	for ind, valor := range vector {
		if valor > vector[maxInd] {
			maxInd = ind
		}
	}
	return maxInd
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	longitud1 := len(vector1)
	longitud2 := len(vector2)

	minimo := longitud2
	if longitud1 < longitud2 {
		minimo = longitud1
	}

	for i := range minimo {
		val1 := vector1[i]
		val2 := vector2[i]

		if val1 != val2 {
			if val1 < val2 {
				return -1
			} else {
				return 1
			}
		}
	}
	if longitud1 == longitud2 {
		return 0
	} else if longitud1 < longitud2 {
		return -1
	} else {
		return 1
	}
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector); i > 1; i-- {
		posMax := Maximo(vector[:i])
		Swap(&vector[posMax], &vector[i-1])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	longitud := len(cadena)

	if longitud <= 1 {
		return true

	} else if longitud == 2 {
		if cadena[0] == cadena[1] {
			return true
		}

	} else if longitud > 2 {
		if cadena[0] == cadena[longitud-1] {
			return EsCadenaCapicua(cadena[1 : longitud-1])
		} else {
			return false
		}

	}
	return false
}
