package ejercicios

// (★★) Se tiene un arreglo de estructuras de la forma type
// Evento struct {anio long, evento string}, que indica el año
// y evento de un hecho definido a lo largo de la historia de la Tierra.
// Indicar y justificar cuál sería un algoritmo de ordenamiento apropiado
// para utilizar para ordenar dicho arreglo por año. Indicar también,
// si en vez de ordenar por año se decide ordenar por evento (lexicográficamente).
// Si se quiere ordenar por año y dentro de cada año,
// por evento: ¿Deben utilizarse para ambos campos el mismo algoritmo de ordenamiento?
// ¿Que característica/s deben cumplir dicho o dichos algoritmos para que quede ordenado como se desea?
// ¿En qué orden deben aplicarse los ordenamientos?

type Evento struct {
	anio   int64
	evento string
}

/*
Utilizaría el método de ordenamiento de Merge Sort (o cualquier otro algoritmo con complejidad O(n log n))
para la realización de este ejercicio, ya que para ordenamientos como Radix, Counting o Bucket sort, requerimos
de información adicional que sea de valor para su elección:
- Counting: El rango de la clave K (años transcurridos en la Tierra) es enormemente superior a la cantidad de elementos N (K > N).
- Bucket: No se puede garantizar una distribución uniforme de los eventos a lo largo del tiempo histórico.
- Radix: La descomposición de claves es inviable, debido a su gran cantidad de dígitos.
- Merge Sort: Complejidad O(N log N) garantizada por comparación, independiente del rango de años o su distribución.

Aplica lo mismo para el caso de únicamente ordenar por evento, con la condición de que el K es enorme (todas las combinaciones
posibles de caracteres).

Para el ordenamiento por año y evento seguimos utilizando merge sort en ambos casos, por las mismas razones señaladas (además
de que ambos algoritmos son estables, por lo que).

La característica que debe cumplirse es que el segundo algoritmo sea estable, de manera tal de no modificar el primer
ordenamiento realizado.

El orden de aplicación deberia ir del de menor a mayor relevancia. En este caso, evento -> año.
*/

func mergesort[T any](arr []T, comparar func(T, T) bool) []T {
	largo := len(arr)

	if largo <= 1 {
		return arr
	}

	medio := largo / 2

	pMitad := mergesort(arr[:medio], comparar)
	sMitad := mergesort(arr[medio:], comparar)

	return merge(pMitad, sMitad, comparar)
}

func merge[T any](izq, der []T, comparar func(T, T) bool) []T {
	res := make([]T, 0, len(izq)+len(der))

	i, j := 0, 0

	for i < len(izq) && j < len(der) {
		if comparar(izq[i], der[j]) {
			res = append(res, izq[i])
			i++
		} else {
			res = append(res, der[j])
			j++
		}
	}

	res = append(res, izq[i:]...)
	res = append(res, der[j:]...)

	return res
}

func OrdenarAnios(eventos []Evento) []Evento {
	return mergesort(eventos, func(e1, e2 Evento) bool {
		return e1.anio <= e2.anio
	})
}

func OrdenarEventos(eventos []Evento) []Evento {
	return mergesort(eventos, func(e1, e2 Evento) bool {
		return e1.evento <= e2.evento
	})
}

func OrdenarAniosYEventos(eventos []Evento) []Evento {
	eventos = mergesort(eventos, func(e1, e2 Evento) bool {
		return e1.evento <= e2.evento
	})

	eventos = mergesort(eventos, func(e1, e2 Evento) bool {
		return e1.anio <= e2.anio
	})

	return eventos
}
