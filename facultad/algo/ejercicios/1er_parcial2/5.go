package ejercicios

/*
Hacer el seguimiento de counting sort para ordenar por año los siguiente sucesos del siglo XX.
¿Cuál es la complejidad del algoritmo?
¿Qué sucede con el orden de los elementos de un mismo año, respecto al orden inicial, luego de finalizado el algoritmo?
Justificar brevemente.

1955 - Pacto de Varsovia
1956 - Fundación de Flat Earth Soc.
1955 - Bombardeo a Plaza de Mayo
1953 - Revolución Cubana
1952 - Coronación de Elizabeth II
1961 - Yuri Gagarin viaja al espacio
1958 - Creación de la NASA
1961 - Construcción del Muro de Berlín
1950 - Guerra de Corea
1955 - Creación Vacuna Poliomielítis
*/

/*
Para determinar la complejidad del algoritmo de ordenamiento counting sort, debemos determinar un rango de claves
y analizar las operaciones del algoritmo en si.

Entonces, para un arreglo de N elementos y un rango de claves K:

Lo primero que se realiza es el armado del arreglo de frecuencias, el cual tendrá un largo K. El costo resultante
de esta operación es O(N), ya que se hacen procesamientos de tiempo constante N veces.

Luego se arma el arreglo de posiciones de largo K, que indica dónde va cada uno de los elementos. Al tener largo K,
se ejecutan operaciones O(1) K veces, por lo tanto tiene una complejidad de O(K).

Finalmente se recorre el arreglo original, donde se hacen operaciones constantes para ubicar cada elemento en su
posición correspondiente acudiendo al arreglo de posiciones previamente hecho. Esto, la igual que la primera operación
(y por las mismas razones), es O(N)

Esto nos deja en O(N) + O(K) + O(N) = O(N + K), y si ordenamos estos elementos por año, nos queda el rango desde 1955 a
1961, por lo tanto K = 1961 - 1950 + 1 = 12 (+1 para incluir ambos extremos), que nos deja O(N + 12), y sabemos que los
valores constantes son redundantes en comparación a N, por lo que el resultado es O(N).

Luego de finalizado el algoritmo, el orden que tienen los elementos con mismo año es el que
*/
