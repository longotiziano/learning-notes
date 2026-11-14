package ejercicios

/*
4. Implementar una primitiva func (pila *Pila[T]) InsertarEnPos(elemento T, n int) que inserte el elemento en la
posición n de la pila. Es decir, la pila debe quedar con sus elementos originales, más el nuevo elemento, el cual debe salir
como n-ésimo elemento al desapilar.

En caso de que n sea mayor a la cantidad de elementos de la pila, finalizar con panic.

Ejemplo: pila = [1, 2, 3, 4, 5, 6] (tope 6), elemento = 70, n = 2. El estado final debe ser pila = [1, 2, 3, 4,
70, 5, 6]. Si n = 0 es la operación resulta equivalente a, simplemente, apilar.
Indicar y justificar la complejidad de la función implementada. Ah, Hollander nos envía sus saludos.
*/
