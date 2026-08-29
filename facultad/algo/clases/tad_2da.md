# Implementación de estructuras
Muchas veces va a convenir implementar estructuras de una forma u otra, sea por simplicidad o porque mejora la complejidad.

### Por qué elegimos una sobre otra
- Complejidad (temporal, espacial).
- Cuán fácil de entender la solución, cuánto tiempo lleva implementarla.

## Políticas de redimensión
Supongamos que estamos creando nuestro propio TDA para una pila. Para manejar la máxima cantidad de elementos de manera dinámica, debemos redimensionar. Para esto, siempre queremos mantener la menor complejidad posible.

Una técnica muy efectiva es duplicar la capacidad de la pila cuando se llegue a su máximo, de esta manera conseguiriamos una complejidad de O(n log(n)), que resulta en O(n), donde la base del logaritmo depende si duplicadomos, triplicamos, etc...

Esto se llama **análisis (de complejidad) amortizado**.

Otra técnica de redimensión efectiva, pero aplicada al caso de reducir el tamaño de la pila, es que al momento de que se ocupe un cuarto del total de la capacidad, reducir la misma a la mitad. De esta manera, la pila aún quedaría con el doble de su capacidad.

### Append
En la cátedra, no utilizamos `append`, ya que este método únicamente agranda el vector. Esto va en contra de lo que la cátedra quiere enseñar, ya que busca que manejemos la redimensión, lo que también incluye el achicamiento de capacidad. Si usásemos `append`, de igual manera tendríamos que crear una lógica de achicamiento (que es la misma que la de agrandamiento), generando código inconsistente.

Esto no significa que no debamos usar el método fuera, simplemente es por la cuestión de la redimensión.

## Invariantes de representación
Características que internamente la representación siempre cumple, y que se cosideran un estado válido que todas las primitivas pueden asumir.

Ej: cantidad == cantidad de elementos en la pila, capacidad == cantidad de posiciones que son válidas en el arreglo de datos.

## Estructuras enlazadas
Manejadas por nodos (vagones, en criollo). Por cada elemento hay qeu crear un nodo. Para eliminar uno hay que deshacer el nodo. Hay que mantener correctamente las referencias.

En la materia hacemos pila en arreglo y cola enlazada, por motivos pedagógicos, eficiencia (por cómo se maneja una pila), menor complejidad.

Para hacer el TDA de cola sobre estructuras enlazadas es fácil. Simplemente agregas un puntero al principio y al final, y cambiás los valores según la operación.