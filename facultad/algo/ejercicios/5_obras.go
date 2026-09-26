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
del arreglo N e identificar el rango de claves K, que en este caso se trata de un rango de 18 claves (1987 a 2005).

Como la complejidad de counting sort es de T(N) = O(N + K), con N = 8 y K = 19

En este algoritmo, lo primero que tenemos que hacer es el arreglo de frecuencias, que tendrá un largo K,
y contará para cada K cuantos elementos hay (de manera lineal).

Luego, en otro arreglo, se introducirá cada elemento haciendo "saltos" por cada K, y el tamaño de los saltos
dependerá de la frecuencia de cada K. Por ejemplo, si tenemos que 1995 aparece 3 veces en las frecuencias,
entonces antes de pasar a la próxima clave, en este nuevo arreglo se dejarán 3 posiciones vacías. Este
proceso tiene complejidad O(K).

Esto permite realizar
*/

type Obra struct {
	anio   int
	titulo string
}

func OrdenarObras(obras []Obra) []Obra {
	rango := 2005 - 1987 + 1
	return countingSortDos(obras, rango, func(o Obra) int {
		return o.anio - 1987
	})
}

func countingSortDos[T any](arr []T, rango int, selecDigito func(T) int) []T {
	return nil
}
