package ejercicios

/*

const BUCKETS = 100

type elem struct {
	cadena  string
	hashing int64
}

func Ordenar(cadenas []string, valoresHash []int64, K int64) []string {
	baldecitos := make([][]elem, BUCKETS)
	for i := range baldecitos {
		baldecitos[i] = make([]elem, 0)
	}
	distancia := K / BUCKETS
	for i := range cadenas {
		nuevo := elem{cadenas[i], valoresHash[i]}
		cual_balde := valoresHash[i] / distancia
		baldecitos[cual_balde] = append(baldecitos[cual_balde], nuevo)
	}

	for i, balde := range baldecitos {
		baldecitos[i] = mergesort(balde)
	}

	resultado := make([]string, len(cadenas))
	indice := 0
	for _, balde := range baldecitos {
		for _, elemento := range balde {
			resultado[indice] = elemento.cadena
			indice += 1
		}
	}
	return resultado
}

func raiz(f func(int) int, a int, b int) int {
	medio := (b + a) / 2
	aplicado := f(medio)
	if aplicado == 0 {
		return medio
	}
	if f(a) == 0 {
		return a
	}
	if aplicado*f(a) > 0 {
		return raiz(f, medio+1, b)
	} else {
		return raiz(f, a, medio-1)
	}
}

func parteEnteraRaiz(n int) int {
	if n == 0 {func TestOrdenarAnios(t *testing.T) {
	eventos := []Evento{
		{anio: 2000, evento: "Zorro"},
		{anio: 1990, evento: "Historia"},
		{anio: 2000, evento: "Arbol"},
	}

	resultado := OrdenarAnios(eventos)

	require.Equal(t, int64(1990), resultado[0].anio)
	require.Equal(t, "Historia", resultado[0].evento)

	require.Equal(t, int64(2000), resultado[1].anio)
	require.Equal(t, "Arbol", resultado[1].evento)

	require.Equal(t, int64(2000), resultado[2].anio)
	require.Equal(t, "Zorro", resultado[2].evento)
}
		return 0
	}
	return _parteEnteraRaizRec(n, 0, n)
}

func cuadrado(n int) int {
	return n * n
}

func _parteEnteraRaizRec(n int, inicio int, fin int) int {
	if inicio == fin {
		return inicio
	}
	medio := (inicio + fin) / 2
	if cuadrado(medio) <= n && cuadrado(medio+1) > n {
		return medio
	} else if cuadrado(medio) > n {
		return _parteEnteraRaizRec(n, inicio, medio-1)
	} else {
		return _parteEnteraRaizRec(n, medio+1, fin)
	}
}
*/
