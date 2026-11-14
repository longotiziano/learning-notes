package ejercicios

/*
El equipo de desarrollo de AlgoAPP está en problemas. Acaban de lanzar una nueva versión de su aplicación y todo explotó.
Tienen un historial de versiones numeradas correlativamente (1, 2, . . . , N) y saben que hubo una versión a partir de la cual
todo empezó a andar mal (incluyéndola). Es decir, todas las versiones desde esa versión rota en adelante están fallando.
De cada versión se puede consultar:
version.autor(): devuelve, en O(1), el nombre de la persona que introdujo la versión.
version.estaOk(): devuelve, también en O(1), true si la versión funciona correctamente, false en caso contrario.
Implementar una función eficiente que reciba un arreglo con todas las versiones y devuelva el nombre de la persona que
introdujo la primera versión rota. Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

type Version struct {
	autorNombre string
	estado      bool
}

func (v *Version) autor() string {
	return v.autorNombre
}
func (v *Version) estaOk() bool {
	return v.estado
}

func ObtenerCulpableVersion(versiones []Version) string {
	return helperObtenerCulpable(versiones, 0, len(versiones))
}

func helperObtenerCulpable(versiones []Version, ini int, fin int) string {
	if fin-ini <= 1 {
		return versiones[ini].autor()
	}
	medio := (ini + fin) / 2
	if versiones[medio].estaOk() {
		return helperObtenerCulpable(versiones, medio+1, fin)
	} else {
		return helperObtenerCulpable(versiones, ini, medio)
	}
}

/*
Dado a que se trata de un algoritmo de División y Conquista (por dividir el problema en partes mas pequeñas que vamos resolviendo) y además
las divisiones creadas tienen el mismo tamaño, estamos en las condiciones para aplicar el Teorema Maestro para justificar la complejidad.

Siendo A la cantidad de llamados recursivos, B el tamaño de esas subdivisiones y O(n^C) la complejidad del resto de operaciones:
- A = 1 (único llamado recursivo)
- B = 2 (en cada iteración dividimos en 2 el arreglo)
- n^C = 1 (ya que el resto de operaciones son de tiempo constante), por lo tanto C = 0

Entonces si log_B(A) = C, entonces la complejidad es O(n^C * log n)
-> log_2(1) = 0 = C, por lo tanto O(n^0 * log n) = O(log n)
*/

// testeo:

// t t t t t f f f
// 0, 8
// 4 -> t
// 5, 8
// 13 / 2 = 6 -> f
// 5, 6
// 6 - 5 = 1
// return versiones[5]

// f f f
// 0, 3
// 3 / 2 = 1 -> no ok -> f(v, 0, 1)
// 1 - 0 <= 1 -> return versiones[0]

// t t t f
// f(0, 4)
// 4 / 2 = 2 -> v[2] ok -> f(2+1. 4)
// 4 - 3 <= 1 return v[3]
