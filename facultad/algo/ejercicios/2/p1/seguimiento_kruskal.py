"""
Seguimiento algoritmo de Kruskal.

Con V cantidad de vértices y E cantidad de aristas:
1. Busco, por cada vertice, sus aristas, lo cual tiene una complejidad de O(V+E). Ya que estoy recorriendo todas 
las aristas del grafo y sus vertices.
2. Aplico un ordenamiento a las aristas obtenidas de menor a mayor. Esta complejidad depende del algoritmo elegido, y suponiendo
que no tenemos información acerca de las mismas, elegimos Merge Sort. Esto tendría complejidad O(E log E), y como
en el peor de los casos E = V(V-1)/2 (grafo completo), O(E log E) = O(E log (V^2)) = O(E log V).
3. Inicializamos la estructura del TDA UnionFind con cada vértice en su propia componente conexa, el cual permite revisar y unir 
conjuntos con complejidad O(a(V)) (a inversa de la función de Ackermann) para ambas operaciones. La inicialización tiene coste O(V).
4. Comenzamos a recorrer el arreglo de aristas ordenado, uniendo los vertices entre ellos y asegurándonos que antes de unirlos,
pertenezcan a componentes conexas diferentes.

Esto nos dejaría una complejidad total O(V+E) + O(E log V) + O(V) + O(E*a(V)), lo que se resume en O(E log V).
"""