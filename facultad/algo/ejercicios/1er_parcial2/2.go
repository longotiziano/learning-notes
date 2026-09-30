package ejercicios

/*
Fulanito necesita automatizar la verificación de sus ejercicios de derivadas mediante una función implementada en Go que, dada una
función continua f(x), encuentre un mínimo local dentro de un rango discreto [a, b]. Se garantiza que f'(x) = 0 en algún punto
del rango (y sucede en algún x entero).

Además se asegura que f'(a) < 0 y f'(b) > 0. -> esto asegura el minimo

Para todos los casos se trabajará con valores enteros.

Implementar una función que reciba la función f y los valores a y b y devuelva un mínimo local dentro de la misma, en
tiempo O(log n), siendo n = b − a (la longitud del rango). Justificar la complejidad de la función.
*/

func minimo(f func(int) int, a, b int) int {
	if a == b {
		return a
	}

	m := (a + b) / 2

	if f(m) < f(m+1) {
		return minimo(f, a, m)
	}

	return minimo(f, m+1, b)
}

// -2, 2, x^2
// 4 != 2
// 0 / 2
// if 0 >= 1 false
// -2, 0
