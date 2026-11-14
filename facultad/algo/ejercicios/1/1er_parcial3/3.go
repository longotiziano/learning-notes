package ejercicios

import (
	"strconv"
)

/*
Una importadora de vinos franceses almacena todos sus productos con su fecha de añejamiento en formato MM-DD-AAAA.

Necesitan ordenarlos para saber cuáles están listos para distribuir (los más añejos primero).

Implementar un ordenamiento lineal que, dado un arreglo con todos los vinos (tipo de dato), devuelva un nuevo arreglo ordenado por fecha de añejamiento
de la más antigua a la más reciente.

Además, se desea saber de qué año es el vino más añejo de todo el lote. Tanto el arreglo ordenado como el año deben ser
devueltos por la función.

El tipo de dato Vino tiene implementado el método Añejamiento(), que devuelve la fecha del vino como un string con
formato MM-DD-AAAA. Indicar y justificar la complejidad del algoritmo implementado, desarrollándola de forma completa.
*/

type Vino struct {
	fecha string
}

func (v *Vino) Aniejamiento() string {
	return v.fecha
}

func counting[T any](arr []T, rango int, obtenerDigito func(elem T) int) []T {
	frec := make([]int, rango)
	for _, elem := range arr {
		frec[obtenerDigito(elem)]++
	}
	arrPosiciones := make([]int, rango)
	for i := 1; i < rango; i++ {
		arrPosiciones[i] = arrPosiciones[i-1] + frec[i-1]
	}
	res := make([]T, len(arr))
	for _, elem := range arr {
		d := obtenerDigito(elem)
		pos := arrPosiciones[d]
		res[pos] = elem
		arrPosiciones[d]++
	}
	return res
}

func obtenerAnio(v Vino) int {
	a, _ := strconv.Atoi(v.Aniejamiento()[6:10])
	return a
}

func OrdenarPorAniejo(arr []Vino) ([]Vino, int) {
	arr = counting(arr, 31, func(elem Vino) int {
		d, _ := strconv.Atoi(elem.Aniejamiento()[3:5])
		return d
	})
	arr = counting(arr, 12, func(elem Vino) int {
		m, _ := strconv.Atoi(elem.Aniejamiento()[0:2])
		return m
	})
	arr = counting(arr, 2026-1900+1, func(elem Vino) int {
		a := obtenerAnio(elem)
		return a - 1900
	})
	anio := 0
	if len(arr) > 0 {
		anio = obtenerAnio(arr[0])
	}
	return arr, anio
}

/*
El algoritmo de ordenamiento implementado para este ejercicio fue radix sort, utilizando como ordenamiento auxiliar
counting sort en cada una de sus subdivisiones.

Elegí radix sort ya que dentro de cada fecha, podemos subdividirla en tres partes con diferentes jerarquías, donde
en cada una de ellas podemos establecer un rango de claves limitado, lo que nos permite pensar en utilizar countingsort
como algoritmo auxiliar, que a su vez es compatible con radix debido a su estabilidad.

Pasando al análisis de complejidad, sabemos que radix sort tiene una complejidad de O(D * H), siendo H la complejidad del
algoritmo interno y D la cantidad de subdivisiones, donde esta última es 3.

Sabemos también que dado un N (cantidad de elementos) y un K (rango de valores posibles), counting sort tiene una complejidad
de O(N + K).

Por lo tanto, si H = O(N + K), entonces la complejidad final del algoritmo de ordenamiento es O(3(N + K)), que al ser K y 3 constantes,
tiende a tener complejidad lineal O(N).
*/
