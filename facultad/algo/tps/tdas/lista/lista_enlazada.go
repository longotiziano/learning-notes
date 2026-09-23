package lista

type listaNodo[T any] struct {
	dato      T
	siguiente *listaNodo[T]
}

func nodoCrear[T any](dato T) *listaNodo[T] {
	return &listaNodo[T]{
		dato, nil,
	}
}
