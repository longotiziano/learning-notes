package ejercicios

type Nacionalidad int

const (
	Arg Nacionalidad = iota
	Eur
)

type Persona struct {
	nombre       string
	nacionalidad Nacionalidad
	edad         int
}

func OrdenarFila(personas []Persona) []Persona {
	niños := make([]Persona, 0)
	noNiños := make([]Persona, 0)
	for _, p := range personas {
		if p.edad <= 12 {
			niños = append(niños, p)
		} else {
			noNiños = append(noNiños, p)
		}
	}
	ordenarNiños(niños)
	ordenarNoNiños(noNiños)

	resultado := make([]Persona, len(personas))
	resultado = append(resultado, niños...)
	resultado = append(resultado, noNiños...)
	return resultado
}

func ordenarNiños(personas []Persona) {
	countingSort(personas, 13, func(p Persona) int {
		return p.edad
	})
}

func ordenarNoNiños(personas []Persona) {
	countingSort(personas, 120, func(p Persona) int {
		return p.edad - 12
	})
	countingSort(personas, 32, func(p Persona) int {
		return int(p.nacionalidad)
	})
}

func countingSort[T any](elementos []T, rango int, seleccDigito func(T) int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]T, len(elementos))

	// creo el arreglo de frecuencias
	for _, elem := range elementos {
		valor := seleccDigito(elem)
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	for _, elem := range elementos {
		valor := seleccDigito(elem)
		pos := sumasAcumuladas[valor]
		resultado[pos] = elem
		sumasAcumuladas[valor]++
	}

	for i := 0; i < len(elementos); i++ {
		elementos[i] = resultado[i]
	}
}
