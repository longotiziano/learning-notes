package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const _ERROR_CLAVE_NO_PERTENECE = "La clave no pertenece al diccionario"

type campo[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[campo[K, V]]
	tam      int
	cantidad int
}

// FNV-1a, implementado mediante el paquete estándar hash/fnv de Go.
// Fuente del algoritmo: Fowler, Noll y Vo, "FNV Hash"
// https://www.isthe.com/chongo/tech/comp/fnv/
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hash[K comparable](clave K) uint64 {
	h := fnv.New64a()
	h.Write(convertirABytes(clave))
	return h.Sum64()
}

func (h *hashAbierto[K, V]) obtenerIndice(clave K) int {
	return int(hash(clave) % uint64(h.tam))
}

func crearCampo[K comparable, V any](clave K, valor V) campo[K, V] {
	return campo[K, V]{clave, valor}
}

func (h *hashAbierto[K, V]) obtenerClaveLista(clave K) TDALista.Lista[campo[K, V]] {
	return h.tabla[h.obtenerIndice(clave)]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		make([]TDALista.Lista[campo[K, V]], 100),
		100,
		0,
	}
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(clave K, valor V) {

}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return iterExterno{}
}

func (h *hashAbierto[K, V]) obtenerElementoAux(clave K) (TDALista.IteradorLista[campo[K, V]], bool) {
	lista := h.obtenerClaveLista(clave)
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			return iter, true
		}
	}
	return nil, false
}

func (h *hashAbierto[K, V]) Guardar(clave K, valor V) {
	lista := h.obtenerClaveLista(clave)

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			// Acá habría que reemplazar el valor existente
			// según las operaciones que tenga tu iterador.
			return
		}
	}

	lista.InsertarUltimo(crearCampo(clave, valor))
	h.cantidad++
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, encontrado := h.obtenerElementoAux(clave)
	return encontrado
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	iter, encontrado := h.obtenerElementoAux(clave)
	if !encontrado {
		panic(_ERROR_CLAVE_NO_PERTENECE)
	}

	return iter.VerActual().dato
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	iter, encontrado := h.obtenerElementoAux(clave)
	if !encontrado {
		panic(_ERROR_CLAVE_NO_PERTENECE)
	}

	borrado := iter.Borrar()
	h.cantidad--
	return borrado.dato
}
