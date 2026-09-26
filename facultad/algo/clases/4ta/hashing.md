# Hashing

## Iteradores

### Externo
```go
iter.HayAlgoMas() bool
// fallan si no hay algo más:
iter.VerActual() (k, v)
iter.Avanzar()

for iter := dic.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
    k, v := iter.VerActual()
}
```
Los iteradores externos, a diferencia de por ejemplo TDA de lista, no contienen elementos de inserción o borrado, ya que no existe un concepto de orden dentro de estas estructuras.

El iterador siempre inicializa en la primera posición válida.

### Complejidad
Cuando hablamos de las complejidades de algoritmos iterativos de hahes, tienden a ser `O(M)`, siendo M el largo de la tabla. No me importa el largo de cada lista (en el caso del hash abierto), ya que no me cambia generalmente tanto.