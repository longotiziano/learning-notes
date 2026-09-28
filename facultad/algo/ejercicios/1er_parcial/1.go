package ejercicios

import (
	"strconv"
	"strings"
)

/*
ORDEN DE IMPORTANCIA:
- Blue | Yellow: Ordenados por InicioSubscripcion (AAAA-MM-DD) ascendiente
- Black: Lo mismo pero descediente
- Turista: Ordenado por CoefApurado descendiente
*/

// Justificar elección de algoritmo y complejidad.

type Categoria int

const (
	BlueYellow Categoria = iota
	Black
	Turista
)

type Pasajero struct {
	Nombre             string
	InicioSubscripcion string
	CoefApurado        int8
	Categoria          Categoria
}

/*
Dentro de este ordenamiento podemos distinguir varios ordenamientos internos, del cual únicamente
podemos globalizar el de mayor importancia: por categorías.

La estrategia elegida será subdividir el arreglo por categorías, de manera tal de poder aplicar el
ordenamiento independiente en cada sector, sin tocar el resto de registros. Esto para finalmente
mergearlos de manera adecuada.

Para el ordenamiento de las categorías Blue/Yellow y Black de fechas, se decidió utilizar Radix Sort,
subdividiendo el formato YYYY-MM-DD y ordenando de menor relevancia a mayor (día -> mes -> año),
siendo Counting sort el algoritmo auxiliar debido a que podemos acotar el rango de claves K para
el día hasta 31, mes hasta 12 y año (dado que se trata a un inicio de subscripción) podemos dar 100 años
atrás (2026 - 1926 + 1). Agrego también que se refuerza la elección de counting sort por su estabilidad,
ya que de no tenerla, no podría ser utilizado como algoritmo auxiliar de radix.

En el caso de los turistas, dado a que se trata de int8 está acotado de 0 a 127, por lo tanto también
podemos utilizar Counting sort.
*/

func obtenerCategorias(arr []Pasajero) ([]Pasajero, []Pasajero, []Pasajero) {
	var blueYellowArr, blackArr, turistaArr []Pasajero
	// O(N), ya que son operaciones ctes N veces (lineal)
	for _, p := range arr {
		if p.Categoria == BlueYellow {
			blueYellowArr = append(blueYellowArr, p)
		} else if p.Categoria == Black {
			blackArr = append(blackArr, p)
		} else {
			turistaArr = append(turistaArr, p)
		}
	}
	return blueYellowArr, blackArr, turistaArr
}

func counting[T any](arr []T, rango int, selecDigito func(elem T) int, asc bool) {
	// O(1)
	resultado := make([]T, len(arr))
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	// O(N)
	for _, elem := range arr {
		// O(1)
		d := selecDigito(elem)
		frecuencias[d]++
	}

	if asc {
		// O(K) (cualquiera de las 2 alternativas)
		for i := 1; i < rango; i++ {
			// o(1)
			sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
		}
	} else {
		for i := rango - 2; i >= 0; i-- {
			// o(1)
			sumasAcumuladas[i] = sumasAcumuladas[i+1] + frecuencias[i+1]
		}
	}
	// O(N)
	for _, elem := range arr {
		// o(1)
		d := selecDigito(elem)
		pos := sumasAcumuladas[d]
		resultado[pos] = elem
		sumasAcumuladas[d]++
	}
	// O(N)
	for i := range arr {
		// o(1)
		arr[i] = resultado[i]
	}
} // O(N) + O(K) + O(N) + O(N) + operaciones de complejidad cte = O(N + K)

func ordenarPorFecha(arr []Pasajero, asc bool) {
	// O(N + K) con K = 31
	counting(arr, 31, func(p Pasajero) int {
		partes := strings.Split(p.InicioSubscripcion, "-")
		dia, _ := strconv.Atoi(partes[2]) // op ctes
		return dia
	}, asc)
	// O(N + K) con K = 12
	counting(arr, 12, func(p Pasajero) int {
		partes := strings.Split(p.InicioSubscripcion, "-")
		mes, _ := strconv.Atoi(partes[1]) // op ctes
		return mes
	}, asc)
	// O(N + K) con K = 100
	counting(arr, 100, func(p Pasajero) int {
		partes := strings.Split(p.InicioSubscripcion, "-")
		anio, _ := strconv.Atoi(partes[0]) // op ctes
		return anio - 1927
	}, asc)
} // T(N) = O(3(N + K)) que tiende a O(N + K)

func OrdenarPasajeros(arr []Pasajero) []Pasajero {
	resultado := make([]Pasajero, 0, len(arr))                    // o(1)
	blueYellowArr, blackArr, turistaArr := obtenerCategorias(arr) // o(N)
	ordenarPorFecha(blueYellowArr, true)                          // O(3(N1 + K)
	ordenarPorFecha(blackArr, false)                              // O(3(N2 + K))
	counting(turistaArr, 128, func(p Pasajero) int {
		return int(p.CoefApurado)
	}, true) // O(N + K)
	resultado = append(resultado, blueYellowArr...) // o(n1)
	resultado = append(resultado, blackArr...)      // o(n2)
	resultado = append(resultado, turistaArr...)    // o(n3)
	return resultado
} // O(1) + O(N) + O(3(N1 + (31 + 12 + 100)) + O(3(N2 + (31 + 12 + 100)) + O(N3 + 127) + O(N1) + O(N2) + O(N3)) TIENDE A O(N)
