package ejercicios

/*
Hacer el seguimiento de counting sort para ordenar por año las siguientes obras:

    1988 - Crónicas del Ángel Gris
    2000 - Los Días del Venado
    1995 - Alta Fidelidad
    1987 - Tokio Blues
    2005 - En Picada
    1995 - Crónica del Pájaro que Da Cuerda al Mundo
    1995 - Ensayo Sobre la Ceguera
    2005 - Los Hombres que No Amaban a las Mujeres

¿Cuál es el orden del algoritmo?
¿Qué sucede con el orden de los elementos de un mismo año, respecto al orden inicial,
luego de finalizado el algoritmo? Justificar brevemente.
*/

/*
Si se quiere ordenar por año mediante counting sort, entonces tenemos que tener en cuenta el largo
del arreglo N e identificar el rango de claves K, que en este caso se trata de un rango de 9 claves (1987 a 2005).

Como la complejidad de counting sort es de T(N) = O(N + K), con N = 8 y K = 19

En este algoritmo, lo primero que tenemos que hacer es el arreglo de frecuencias, que tendrá un largo K,
y contará para cada K cuantos elementos hay (de manera lineal).

Luego, en otro arreglo, se introducirán las posiciones respectivas de cada elemento, dejando "saltos" con
tamaño dependiente de la cantidad de cada clave. Este proceso, al recorrer el arreglo de frecuencias, es O(K).

Posteriormente se insertan los elementos en sus respectivas posiciones, actualizando cada una de ellas por
elemento insertado. Esto es O(N).

Finalmente se copia el arreglo resultante al introducido (O(N)).
*/

type Obra struct {
	anio   int
	titulo string
}

func OrdenarObras(obras []Obra) {
	rango := 2005 - 1987 + 1
	countingSortDos(obras, rango, func(o Obra) int {
		return o.anio - 1987
	})
}

func countingSortDos[T any](arr []T, rango int, selecDigito func(T) int) {
	lenArr := len(arr)
	frecuencia := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]T, lenArr)

	for _, elem := range arr {
		d := selecDigito(elem)
		frecuencia[d]++
	}
	for i := 1; i < rango; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencia[i-1]
	}
	for _, elem := range arr {
		d := selecDigito(elem)
		pos := sumasAcumuladas[d]
		resultado[pos] = elem
		sumasAcumuladas[d]++
	}
	for i := range arr {
		arr[i] = resultado[i]
	}
}
