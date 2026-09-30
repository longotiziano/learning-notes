package hash

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
