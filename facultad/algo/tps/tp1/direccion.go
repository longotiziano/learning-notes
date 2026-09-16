package main

// Creación de tipo de dato interno del paquete, representando las direcciones
type direccion int

const (
	ARRIBA direccion = iota
	ABAJO
	IZQUIERDA
	DERECHA
)

func direccionEntrePuntos(inicio punto, destino punto) direccion {
	delta := destino.restar(inicio)

	for _, dir := range obtenerDirecciones() {
		if dir.obtenerDireccionAPunto() == delta {
			return dir
		}
	}
	panic(_ERROR)
}

func obtenerDirecciones() []direccion {
	return []direccion{ARRIBA, ABAJO, IZQUIERDA, DERECHA}
}

func (d direccion) String() string {
	switch d {
	case ARRIBA:
		return "ARRIBA"
	case ABAJO:
		return "ABAJO"
	case IZQUIERDA:
		return "IZQUIERDA"
	case DERECHA:
		return "DERECHA"
	default:
		return "ERROR"
	}
}

// método de Direccion. convierte una dirección en un punto
func (d direccion) obtenerDireccionAPunto() punto {
	switch d {
	case ARRIBA:
		return punto{fila: -1, col: 0}
	case ABAJO:
		return punto{fila: 1, col: 0}
	case IZQUIERDA:
		return punto{fila: 0, col: -1}
	case DERECHA:
		return punto{fila: 0, col: 1}
	default:
		return punto{fila: 0, col: 0}
	}
}
