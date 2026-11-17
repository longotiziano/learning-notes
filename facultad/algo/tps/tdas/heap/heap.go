package cola_prioridad

type cmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      cmpHeap[T]
}

const _TAMANIO_INICIAL = 10
const _ERROR_COLA_VACIA = "La cola esta vacia"

// Crea una ColaPrioridad vacía
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{make([]T, _TAMANIO_INICIAL), 0, funcion_cmp}
}

// Ingresado un arreglo y una función comparativa, devuelve una ColaPrioridad con los datos ordenados
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	tamanio := len(arreglo)
	heapify(arreglo, tamanio, funcion_cmp)
	return &heap[T]{arreglo, tamanio, funcion_cmp}
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func calcularHijos(pos int) (int, int) {
	return 2*pos + 1, 2*pos + 2
}

func downheap[T any](datos []T, pos int, cantidad int, cmp cmpHeap[T]) {
	for {
		izq, der := calcularHijos(pos)
		mayor := pos
		if izq < cantidad && cmp(datos[izq], datos[mayor]) > 0 {
			mayor = izq
		}
		if der < cantidad && cmp(datos[der], datos[mayor]) > 0 {
			mayor = der
		}
		if mayor == pos {
			break
		}
		swap(datos, pos, mayor)
		pos = mayor
	}
}

func heapify[T any](datos []T, cantidad int, cmp cmpHeap[T]) {
	for i := cantidad/2 - 1; i >= 0; i-- {
		downheap(datos, i, cantidad, cmp)
	}
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(dato T) {

}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic(_ERROR_COLA_VACIA)
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic(_ERROR_COLA_VACIA)
	}
	return h.datos[0]
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T any](elementos []T, funcion_cmp cmpHeap[T]) {
	tamanio := len(elementos)
	heapify(elementos, tamanio, funcion_cmp)
	for tamanio > 1 {
		swap(elementos, 0, tamanio-1)
		tamanio--
		downheap(elementos, 0, tamanio, funcion_cmp)
	}
}

func cmpMin(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	}
	return 0
}

func ordenarLogK(arr []int, k int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	limite := min(k+1, n)
	inicial := make([]int, limite)
	copy(inicial, arr[:limite])

	heap := CrearHeapArr(inicial, cmpMin) // O(k)

	pos := 0
	for i := limite; i < n; i++ { // O(n - k) vueltas
		arr[pos] = heap.Desencolar() // O(log k)
		heap.Encolar(arr[i])         // O(log k)
		pos++
	}

	for !heap.EstaVacia() { // O(k) vueltas
		arr[pos] = heap.Desencolar() // O(log k)
		pos++
	}
}

type cancion struct {
	titulo   string
	artista  string
	anio     int
	duracion int
}
type itemCancion struct {
	c          cancion
	grupoIdx   int
	cancionIdx int
}

func cmpItem(a, b itemCancion) int {
	return b.c.anio - a.c.anio // Min-Heap por año
}

func ordenarKCanciones(canciones [][]cancion, k, n int) []cancion {
	if k == 0 || n == 0 {
		return []cancion{}
	}

	iniciales := make([]itemCancion, k)
	for i := 0; i < k; i++ {
		iniciales[i] = itemCancion{c: canciones[i][0], grupoIdx: i, cancionIdx: 0}
	}

	heap := CrearHeapArr(iniciales, cmpItem) // O(k)
	arrOrdenado := make([]cancion, 0, k*n)

	for !heap.EstaVacia() { // O(k * n * log k)
		minimo := heap.Desencolar()
		arrOrdenado = append(arrOrdenado, minimo.c)

		sigIdx := minimo.cancionIdx + 1
		if sigIdx < n {
			heap.Encolar(itemCancion{
				c:          canciones[minimo.grupoIdx][sigIdx],
				grupoIdx:   minimo.grupoIdx,
				cancionIdx: sigIdx,
			})
		}
	}

	return arrOrdenado
}
