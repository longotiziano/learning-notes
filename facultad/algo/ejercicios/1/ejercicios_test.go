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

func TestOrdenarObras(t *testing.T) {
	obras := []Obra{
		{anio: 1988, titulo: "Crónicas del Ángel Gris"},
		{anio: 2000, titulo: "Los Días del Venado"},
		{anio: 1995, titulo: "Alta Fidelidad"},
		{anio: 1987, titulo: "Tokio Blues"},
		{anio: 2005, titulo: "En Picada"},
		{anio: 1995, titulo: "Crónica del Pájaro que Da Cuerda al Mundo"},
		{anio: 1995, titulo: "Ensayo Sobre la Ceguera"},
		{anio: 2005, titulo: "Los Hombres que No Amaban a las Mujeres"},
	}
	OrdenarObras(obras)
	resultado := obras

	require.Equal(t, Obra{1987, "Tokio Blues"}, resultado[0])
	require.Equal(t, Obra{1988, "Crónicas del Ángel Gris"}, resultado[1])
	require.Equal(t, Obra{1995, "Alta Fidelidad"}, resultado[2])
	require.Equal(t, Obra{1995, "Crónica del Pájaro que Da Cuerda al Mundo"}, resultado[3])
	require.Equal(t, Obra{1995, "Ensayo Sobre la Ceguera"}, resultado[4])
	require.Equal(t, Obra{2000, "Los Días del Venado"}, resultado[5])
	require.Equal(t, Obra{2005, "En Picada"}, resultado[6])
	require.Equal(t, Obra{2005, "Los Hombres que No Amaban a las Mujeres"}, resultado[7])
}
