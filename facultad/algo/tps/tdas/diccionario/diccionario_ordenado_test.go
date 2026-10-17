package diccionario_test

import (
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// revisar

func TestIteradorInternoSinCondicionDeCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	abb.Guardar(10, 100)
	abb.Guardar(2, 20)
	abb.Guardar(16, 160)
	abb.Guardar(4, 40)
	abb.Guardar(8, 80)

	sumaClaves := 0
	sumaValores := 0
	cantidadVisitados := 0

	abb.Iterar(func(clave int, valor int) bool {
		sumaClaves += clave
		sumaValores += valor
		cantidadVisitados++
		return true
	})

	require.EqualValues(t, 5, cantidadVisitados)
	require.EqualValues(t, 40, sumaClaves)
	require.EqualValues(t, 400, sumaValores)
}

func TestIteradorInternoConCondicionDeCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	elementos := []int{1, 2, 18, 3, 6, 13, 22}
	for _, elem := range elementos {
		abb.Guardar(elem, elem)
	}

	var visitados []int
	abb.Iterar(func(clave int, valor int) bool {
		if clave > 10 {
			return false
		}
		visitados = append(visitados, clave)
		return true
	})

	require.EqualValues(t, []int{1, 2, 3, 6}, visitados)
}

func TestIteradorInternoRangoNulo(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := []string{"prueba 1", "prueba 2", "prueba 3", "prueba 4"}
	for i, c := range claves {
		abb.Guardar(c, i)
	}

	var visitados []string
	abb.IterarRango(nil, nil, func(clave string, valor int) bool {
		visitados = append(visitados, clave)
		return true
	})
	require.EqualValues(t, []string{"prueba 1", "prueba 2", "prueba 3", "prueba 4"}, visitados)
}

func TestIteradorInternoDesde(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := []string{"prueba 1", "prueba 2", "prueba 3", "prueba 4"}
	for i, c := range claves {
		abb.Guardar(c, i)
	}

	desde := "prueba 2"
	var visitados []string
	abb.IterarRango(&desde, nil, func(clave string, valor int) bool {
		visitados = append(visitados, clave)
		return true
	})
	require.EqualValues(t, []string{"prueba 2", "prueba 3", "prueba 4"}, visitados)
}

func TestIteradorInternoHasta(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := []string{"prueba 1", "prueba 2", "prueba 3", "prueba 4"}
	for i, c := range claves {
		abb.Guardar(c, i)
	}

	hasta := "prueba 3"
	var visitados []string
	abb.IterarRango(nil, &hasta, func(clave string, valor int) bool {
		visitados = append(visitados, clave)
		return true
	})
	require.EqualValues(t, []string{"prueba 1", "prueba 2", "prueba 3"}, visitados)
}

func TestIteradorInternoDesdeYHasta(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	claves := []int{4, 6, 7, 9, 16, 23}
	for i, c := range claves {
		abb.Guardar(c, i)
	}

	desde, hasta := 6, 16
	var visitados []int
	abb.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		visitados = append(visitados, clave)
		return true
	})
	require.EqualValues(t, []int{6, 7, 9, 16}, visitados)
}

func TestIteradorExternoABBVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	iter := abb.Iterador()

	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestIteradorExternoRecorridoOrdenado(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	elementos := []int{2, 6, 8, 9, 10, 26, 97}
	for _, elem := range elementos {
		abb.Guardar(elem, elem*10)
	}

	iter := abb.Iterador()
	esperados := []int{2, 6, 8, 9, 10, 26, 97}

	for _, elem := range esperados {
		require.True(t, iter.HayAlgoMas())
		clave, dato := iter.VerActual()
		require.EqualValues(t, elem, clave)
		require.EqualValues(t, elem*10, dato)
		iter.Avanzar()
	}

	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestIteradorExternoRangosClavesInexistentes(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for _, elem := range []int{3, 5, 11, 15, 18, 24} {
		abb.Guardar(elem, elem)
	}

	desde, hasta := 11, 18
	iter := abb.IteradorRango(&desde, &hasta)

	esperados := []int{11, 15, 18}
	for _, elem := range esperados {
		require.True(t, iter.HayAlgoMas())
		clave, _ := iter.VerActual()
		require.EqualValues(t, elem, clave)
		iter.Avanzar()
	}
	require.False(t, iter.HayAlgoMas())
}

func TestMultiplesIteradores(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	abb.Guardar(1, 1)
	abb.Guardar(2, 2)
	abb.Guardar(3, 3)

	iter1 := abb.Iterador()
	iter2 := abb.Iterador()

	require.True(t, iter1.HayAlgoMas())
	require.True(t, iter2.HayAlgoMas())

	iter1.Avanzar()

	c1, _ := iter1.VerActual()
	c2, _ := iter2.VerActual()

	require.EqualValues(t, 2, c1)
	require.EqualValues(t, 1, c2)
}

func TestIteradorVolumen(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	n := 10000

	for i := 0; i < n; i++ {
		abb.Guardar(i, i)
	}

	iter := abb.Iterador()
	i := 0
	for iter.HayAlgoMas() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, i, clave)
		i++
		iter.Avanzar()
	}
	require.EqualValues(t, n, i)
}
