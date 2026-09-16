package maze

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	TDACola "tdas/cola"
)

type Laberinto struct {
	filas    int
	cols     int
	casillas [][]string
}

const _START = "S"
const _PARED = "#"
const _CAMINO = "."
const _EXIT = "E"

const _ERROR = "ERROR"

func CrearLaberinto(filas int, cols int, casillas [][]string) Laberinto {
	return Laberinto{
		filas:    filas,
		cols:     cols,
		casillas: casillas,
	}
}

// verifica que el punto se encuentre dentro de la matriz
// NO CUENTA PAREDES
func (l Laberinto) estaEnRango(p punto) bool {
	return p.fila >= 0 && p.fila < l.filas && p.col >= 0 && p.col < l.cols
}

// verifico que el punto sea "." o "EXIT"
func (l Laberinto) esCaminoValido(p punto) bool {
	c := l.casillas[p.fila][p.col]
	return (c == _CAMINO || c == _EXIT)
}

// devuelve el objeto Punto con el índice de la ubicación inicial
func (l Laberinto) buscarInicial() punto {
	for i := range l.filas {
		for j := range l.cols {
			if l.casillas[i][j] == _START {
				return punto{i, j}
			}
		}
	}
	// este caso nunca pasa, ya que el enunciado garantiza una S y una E
	return punto{-1, -1}
}

// devuelve un slice de direcciones a puntos que tienen caracteres válidos, están en el rango del laberinto y no fueron recorridos
func (l Laberinto) encolarPuntosValidos(p punto, visitados [][]bool, anteriores [][]punto, cola TDACola.Cola[punto]) {
	for _, dir := range obtenerDirecciones() {
		movimiento := dir.obtenerDireccionAPunto()
		pNuevo := p.sumar(movimiento)
		if l.estaEnRango(pNuevo) {
			if l.esCaminoValido(pNuevo) && !visitados[pNuevo.fila][pNuevo.col] {
				visitados[pNuevo.fila][pNuevo.col] = true
				cola.Encolar(pNuevo)
				// guardo el anterior
				anteriores[pNuevo.fila][pNuevo.col] = p
			}
		}
	}
}

func contarPasos(anteriores [][]punto, inicial punto, actual punto) int {
	distancia := 0
	for actual != inicial {
		actual = anteriores[actual.fila][actual.col]
		distancia++
	}
	return distancia
}

func reconstruirPasos(anteriores [][]punto, inicial punto, final punto) ([]direccion, int) {
	distancia := contarPasos(anteriores, inicial, final)
	pasos := make([]direccion, distancia)
	actual := final
	idx := distancia - 1

	for actual != inicial {
		anterior := anteriores[actual.fila][actual.col]
		pasos[idx] = direccionEntrePuntos(anterior, actual)
		actual = anterior
		idx--
	}

	return pasos, distancia
}

// recorre en olas concéntricas el laberinto
// en caso de no haber solución, entonces paniquea con error
func ResolverLaberinto(l Laberinto) ([]direccion, int, error) {
	// al ser matriz booleana, defaultea a false
	visitados := initMatriz[bool](l.filas, l.cols)
	// creada para poder reconstruir el camino una vez encontrada la salida
	anteriores := initMatriz[punto](l.filas, l.cols)

	cola := TDACola.CrearColaEnlazada[punto]()
	inicial := l.buscarInicial()
	visitados[inicial.fila][inicial.col] = true
	cola.Encolar(inicial)

	for !cola.EstaVacia() {
		actual := cola.Desencolar()

		if l.casillas[actual.fila][actual.col] == "E" {
			pasos, distancia := reconstruirPasos(anteriores, inicial, actual)
			return pasos, distancia, nil
		}
		l.encolarPuntosValidos(actual, visitados, anteriores, cola)
	}
	return nil, 0, fmt.Errorf(_ERROR)
}

// lee la primera línea del stdin y devuelve una tupla representando (filas, columnas)
// si no está la información esperada, devuelve (0, 0)
func ObtenerDimension(text string) (int, int) {
	numsDim := strings.Split(text, " ")
	if len(numsDim) == 2 {
		filas, err1 := strconv.Atoi(numsDim[0])
		cols, err2 := strconv.Atoi(numsDim[1])

		if err1 == nil && err2 == nil {
			return filas, cols
		}
	}

	return 0, 0
}

// helper: inicializa una matriz vacía
func initMatriz[T any](filas int, cols int) [][]T {
	matriz := make([][]T, filas)
	for i := range matriz {
		matriz[i] = make([]T, cols)
	}
	return matriz
}

// recibe el scanner, las filas y columnas del laberinto. Devuelve la matriz del laberinto
// no hace falta validar casillas, ya que el enunciado garantiza los caracteres y F líneas de C caracteres
// complejidad: O(F * C)
func AlmacenarCasillas(scanner *bufio.Scanner, filas int, cols int) [][]string {
	matriz := initMatriz[string](filas, cols)
	for i := range filas {
		// no debería ser false
		if scanner.Scan() {
			casillas := scanner.Text()
			for j := range cols {
				matriz[i][j] = string(casillas[j])
			}
		}
	}
	return matriz
}
