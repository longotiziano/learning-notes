package ejercicios

/*
Implementar una función que reciba un slice de enteros ordenado y un valor K y devuelva cuántas veces aparece ese
valor en el mismo. Indicar y justificar la complejidad del algoritmo implementado.
*/
func ObtenerApariciones(arr []int, k int) int {
	cotaInf := buscarCota(arr, 0, len(arr), k, false)
	cotaSup := buscarCota(arr, 0, len(arr), k, true)

	return cotaSup - cotaInf
}

/*
Gracias a que trata de un algoritmo de División y Conquista, puedo utilizar el Teorema Maestro para calcular
la complejidad de esta búsqueda.

Análisis de complejidad:
- A = 2 (cantidad de llamados recursivos en ObtenerApariciones(...))
- B = 2 (porque divido el arreglo en 2 en cada llamado)
- C = 0 (ya que el resto de operaciones son O(1), por lo tanto n^C = 0 -> C = 1)
Como logA(B) = 1 < 0, entonces por Teorema Maestro, que dice que si logA(B) < C, entonces la complejidad de este algoritmo
es O(n^C * log n) = O(n^0 * log n) = O(log n)
*/
func buscarCota(arr []int, inicio int, fin int, k int, cotaSuperior bool) int {
	if inicio == fin {
		return inicio
	}

	medio := (inicio + fin) / 2

	criterio := k <= arr[medio]

	if cotaSuperior {
		criterio = k < arr[medio]
	}

	if criterio {
		return buscarCota(arr, inicio, medio, k, cotaSuperior)
	}

	return buscarCota(arr, medio+1, fin, k, cotaSuperior)
}
