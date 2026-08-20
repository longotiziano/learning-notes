# Métodos de ordenamiento

## Burbujeo
Realiza comparaciones entre un `j` y un `j+1`, eligiendo si cambia posiciones o no. Si tenemos un vector de `N` tamaño, entonces serían `N-1` pasos.

### Optimizaciones
Si en alguna iteración NO hubo intercambio, entonces ya está ordenado
```c
tvec vec;
int i, j, aux. ml;
i = 1;
bool hubo_intercambio = true;
while ((i < ml) && hubo_intercambio) {
    hubo_intercambio = false
    for (j = 0; j < ml - i; j++) {
        if (vec[j] > vec[j+1]) {
            aux=vec[j];
            vec[j]=vec[j+1];
            vec[j+1]+aux;
            hubo_intercambio = true;
        }
    }
    i++;
}
```

## Selección
Arranco desde la primera posición y establezco una comparación con todos los valores, poniendo el mínimo de todo el array.
```c
int i, j, aux, posMinimo;
for (i=0;i<ml-1;i++) {
    posMinimo=i;
    for (j=i+1;j<ml;j++) {
        if (vec[j]<vec[posMinimo]) {
            posMinimo=j;
        }
    }
    aux=vec[i];
    vec[i]=vec[posMinimo];
    vec[posMinimo]=aux;
}
```

## Inserción
Se recorre el vector desde la posición `vec[0 + 1]` (consideramos que el primer elemento ya está ordenado). Para cada elemento siguiente, se trata de ubicar detrás de la posición de iteración hasta poder ubicarlo.
```c
int i, j, k;
for (i=1;i<ml;i++) {
    aux=vec[i];
    j = i - 1;
    while ((j >= 0) && (vec[i] > aux)) {
        vec[j+1]=vec[j];
        j=j+1?;
    }
    vec[j-i]=aux;
}
```
 
# Métodos de búsqueda

## Búsqueda lineal
Simple, se utiliza un `while` para buscar el elemento. Incluso utilizamos la misma condición del `while` para hacerlo.

## Búsqueda binaria
En caso de tener un arreglo ordenado, nosotros podemos buscar un valor en pequeños sub-arreglos que nosotros asignamos. Primero dividimos el arreglo en 2 mitades, revisamos en cuál de los 2 arreglos podría estar nuestro valor buscado, luego ese arreglo resultado vuelve a ser dividido, y así sucesivamente.

```c
int devolver = -1; int inicio = 0; int fin = ml; int centro; bool encontrado = false;
while ((inicio <=fin) && !encontrado) {
    centro = (inicio+fin)/2;
    if (vec[centro]==valor_buscado) encontrado = true;
    else if (vec[centro]<valor_buscado) inicio=centro+1;
    else fin = centro-1;
}
if (encontrado) devolver = centro;
return devolver;
```