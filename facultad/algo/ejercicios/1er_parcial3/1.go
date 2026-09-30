package ejercicios

import (
	TDALista "tdas/lista"
)

type campo[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[campo[K, V]]
	tam      int
	cantidad int
}

type Diccionario[K comparable, V any] interface {
	Guardar(K, V)
	Pertenece(K) bool
	Obtener(K) V
	Borrar(K) V
	Cantidad() int
	Iterar(func(K, V)) bool
	Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {
	HayAlgoMas() bool
	Avanzar()
	VerActual() (K, V)
}

/*
Después del fenómeno Tim Payne, se quiere analizar a los jugadores con menor cantidad de seguidores de cada selección, y
vaya que nuestro conocimiento es Escarso. Así que el el primer paso será, obviamente, saber cuáles son dichos jugadores.

Se pide implementar una función con la siguiente firma:

func ObtenerMenosConocidos(
	selecciones Diccionario[string, Diccionario[string, int]]
) Diccionario[string, string]

El diccionario de entrada tiene como clave nombre del país y como valor un diccionario.

Los diccionarios de valor tienen como clave nombre de jugador y como valor la cantidad de seguidores que posee.

La función debe devuelvor un nuevo diccionario donde, para cada país (clave), se almacene como valor: el nombre del jugador
con menos seguidores de dicho país. En caso de que haya más de un jugador con la menor cantidad de seguidores, se puede
devolver cualquiera de ellos. Indicar y justificar, cuidadosamente, la complejidad del algoritmo implementado.
*/
