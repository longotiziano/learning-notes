package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	func TestPosicionPico(t *testing.T) {
		arreglo := []int{1, 3, 20, 4, 1, 0}

		require.Equal(t, 2, posicionPico(arreglo))
	}
*/
func TestOrdenarAnios(t *testing.T) {
	eventos := []Evento{
		{anio: 2000, evento: "Zorro"},
		{anio: 1990, evento: "Historia"},
		{anio: 2000, evento: "Arbol"},
	}

	resultado := OrdenarAnios(eventos)

	require.Equal(t, int64(1990), resultado[0].anio)
	require.Equal(t, int64(2000), resultado[1].anio)
	require.Equal(t, int64(2000), resultado[2].anio)
}

func TestOrdenarEventos(t *testing.T) {
	eventos := []Evento{
		{anio: 2000, evento: "Zorro"},
		{anio: 1990, evento: "Historia"},
		{anio: 2000, evento: "Arbol"},
	}

	resultado := OrdenarEventos(eventos)

	require.Equal(t, "Arbol", resultado[0].evento)
	require.Equal(t, "Historia", resultado[1].evento)
	require.Equal(t, "Zorro", resultado[2].evento)
}

func TestOrdenarAniosYEventos(t *testing.T) {
	eventos := []Evento{
		{anio: 2000, evento: "Zorro"},
		{anio: 1990, evento: "Historia"},
		{anio: 2000, evento: "Arbol"},
	}

	resultado := OrdenarAniosYEventos(eventos)

	require.Equal(t, Evento{1990, "Historia"}, resultado[0])
	require.Equal(t, Evento{2000, "Arbol"}, resultado[1])
	require.Equal(t, Evento{2000, "Zorro"}, resultado[2])
}
