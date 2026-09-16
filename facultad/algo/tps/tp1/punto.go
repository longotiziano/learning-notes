package maze

// tipo de dato interno para representar los puntos de la matriz
type punto struct {
	fila int
	col  int
}

func (p punto) sumar(delta punto) punto {
	return punto{
		fila: p.fila + delta.fila,
		col:  p.col + delta.col,
	}
}

func (p punto) restar(delta punto) punto {
	return punto{
		fila: p.fila - delta.fila,
		col:  p.col - delta.col,
	}
}
