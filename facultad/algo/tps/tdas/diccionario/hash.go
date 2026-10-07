package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const _ERROR_CLAVE_NO_PERTENECE = "La clave no pertenece al diccionario"
const _ERROR_EL_ITERADOR_TERMINO = "El iterador termino de iterar"
const _TAMANIO_INICIAL_TABLA_HASH = 11
const _TAMANIO_REDIMENSION = 2
const _CRITERIO_ACHICAMIENTO_REDIMENSION = 4

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

func auxObtenerIndice[K comparable](clave K, tam int) int {
	return int(hash(clave) % uint64(tam))
}

func crearTabla[K comparable, V any](tam int) []TDALista.Lista[campo[K, V]] {
	return make([]TDALista.Lista[campo[K, V]], tam)
}

func (h *hashAbierto[K, V]) obtenerIndice(clave K) int {
	return auxObtenerIndice(clave, h.tam)
}

func crearCampo[K comparable, V any](clave K, valor V) campo[K, V] {
	return campo[K, V]{clave, valor}
}

func (h *hashAbierto[K, V]) obtenerClaveLista(clave K) TDALista.Lista[campo[K, V]] {
	indice := h.obtenerIndice(clave)

	if h.tabla[indice] == nil {
		h.tabla[indice] = TDALista.CrearListaEnlazada[campo[K, V]]()
	}

	return h.tabla[indice]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		crearTabla[K, V](_TAMANIO_INICIAL_TABLA_HASH),
		_TAMANIO_INICIAL_TABLA_HASH,
		0,
	}
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func convertirPrimo(n int) int {
	if n < 2 {
		return 2
	}
	if n%2 == 0 {
		n++
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			n += 2
			i = 1
		}
	}
	return n
}

func (h *hashAbierto[K, V]) redimensionar(nuevaCapacidad int) {
	nuevaCapacidad = convertirPrimo(nuevaCapacidad)
	nuevosDatos := crearTabla[K, V](nuevaCapacidad)

	h.Iterar(func(clave K, dato V) bool {
		indice := auxObtenerIndice(clave, nuevaCapacidad)

		if nuevosDatos[indice] == nil {
			nuevosDatos[indice] = TDALista.CrearListaEnlazada[campo[K, V]]()
		}

		nuevosDatos[indice].InsertarUltimo(crearCampo(clave, dato))
		return true
	})

	h.tabla = nuevosDatos
	h.tam = nuevaCapacidad
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
	iter, encontrado := h.obtenerElementoAux(clave)

	if encontrado {
		iter.Borrar()
		iter.Insertar(crearCampo(clave, valor))
		return
	}

	if h.tam == h.cantidad {
		h.redimensionar(_TAMANIO_REDIMENSION * h.cantidad)
	}
	h.cantidad++
	h.obtenerClaveLista(clave).InsertarUltimo(crearCampo(clave, valor))
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

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	iter, encontrado := h.obtenerElementoAux(clave)
	if !encontrado {
		panic(_ERROR_CLAVE_NO_PERTENECE)
	}

	borrado := iter.Borrar()
	h.cantidad--

	if h.cantidad > 0 && _CRITERIO_ACHICAMIENTO_REDIMENSION*h.cantidad <= h.tam {
		h.redimensionar(h.cantidad * _TAMANIO_REDIMENSION)
	}

	return borrado.dato
}

func (h *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for _, lista := range h.tabla {
		if lista == nil {
			continue
		}

		for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
			actual := iter.VerActual()

			if !visitar(actual.clave, actual.dato) {
				return
			}
		}
	}
}

/*
Iterador externo
*/

type iterHashAbierto[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	posicion  int
	iterLista TDALista.IteradorLista[campo[K, V]]
}

func (iter *iterHashAbierto[K, V]) buscarSiguienteValido() {
	for iter.posicion < iter.hash.tam && (iter.hash.tabla[iter.posicion] == nil || !iter.iterLista.HayAlgoMas()) {
		iter.posicion++
		if iter.posicion < iter.hash.tam && iter.hash.tabla[iter.posicion] != nil {
			iter.iterLista = iter.hash.tabla[iter.posicion].Iterador()
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashAbierto[K, V]{
		hash:     h,
		posicion: 0,
	}

	if h.tam > 0 && h.tabla[0] != nil {
		iter.iterLista = h.tabla[0].Iterador()
	}

	if iter.iterLista == nil || !iter.iterLista.HayAlgoMas() {
		iter.buscarSiguienteValido()
	}

	return iter
}

func (iter *iterHashAbierto[K, V]) HayAlgoMas() bool {
	return iter.posicion < iter.hash.tam && iter.iterLista != nil && iter.iterLista.HayAlgoMas()
}

func (iter *iterHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_ERROR_EL_ITERADOR_TERMINO)
	}

	par := iter.iterLista.VerActual()
	return par.clave, par.dato
}

func (iter *iterHashAbierto[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_ERROR_EL_ITERADOR_TERMINO)
	}

	iter.iterLista.Avanzar()

	if !iter.iterLista.HayAlgoMas() {
		iter.posicion++
		if iter.posicion < iter.hash.tam && iter.hash.tabla[iter.posicion] != nil {
			iter.iterLista = iter.hash.tabla[iter.posicion].Iterador()
		}
		iter.buscarSiguienteValido()
	}
}
