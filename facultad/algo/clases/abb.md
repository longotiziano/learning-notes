# Árboles Binarios de Búsquedas (ABB)

Siempre el elemento menor a la izquierda, y mayor a su derecha. Una ventaja de esta
estructura es que al recorrerlo `in-order` SIEMPRE queda ordenado, ya que primero recorre
los izquierdos y va creciendo. 

Los ABB, a comparación con hash, se optan cuando el orden es importante, ya que las operaciones
tienden a ser logarítmicas. Caso contrario, el promedio O(1) del hash es superior.

Otra ventaja es que no dependemos de una función de hashing, lo que nos da más control para la
realización de (por ejemplo) pruebas.

Otra es el espacio. Si bien se maneja con la redimensión, el árbol tiene complejidad espacial O(N).

Desventaja: A veces es difícil de implementar y BORRAR :v.

Para borrar un elemento en un ABB tenemos que primero detectar el orden del árbol. Luego, ver que elementos 
tengo a los costados en el orden `in-order`. Aquellos que estén inmediatamente después o antes son candidatos
válidos para reemplazar el nodo, ya que cumplen con los casos de que son el más grande del lado izquierdo, y
el más chico del derecho, respectivamente. 

Este elemento SIEMPRE va a tener 1 o 0 hijos, por lo tanto no genera conflictos :3.

En caso de que los elementos se inserten en orden o con cierto patrón, todas las primitivas serán lineales. 
Si los elementos se insertan de manera aleatoria, estará balanceado (o muy bien balanceado).

Para lidiar con los problemas de balance, podemos rotar el árbol. Para darnos cuenta del balance, podemos fijarnos
en la altura del arbol. Siempre queremos que la altura sea logarítmica.

Propiedad AVL (Adelson Velski y Lambdis, creadores de la prioridad): |h_izq - h_der| <= 1

Acerca de rotaciones:
https://drive.google.com/file/d/1-RgQ9bO0hIG_nPfwIhEhv2tsSnrN-Rrn/view?usp=sharing

No van a pedir en el examen que programemos esto.

Siempre se van viendo las alturas de los árboles y subiendo.

No nos van a tomar borrados.

Recomendación: realizar rotaciones intermedias y, además, determinar los X Y y Z.