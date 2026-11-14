"""
(★★★) Implementar un algoritmo que reciba un grafo dirigido, un vértice V y un número N, 
y devuelva una lista con todos los vértices que se encuentren a exactamente N aristas de 
distancia del vértice V. Indicar el tipo de recorrido utilizado y el orden del algoritmo. 
Justificar.
"""

from collections import deque

def a_n_aristas(grafo, origen, n):
    'Devolver una lista con los vértices que cumplen la propiedad'
    if n == 0:
        return [origen]
    vertices_n = []
    visitados = set()
    distancias = {origen: 0} # todo esto es 01
    q = deque()
    q.append(origen)
    visitados.add(origen)
    while q: # O(V)
        v = q.popleft()
        for w in grafo.adyacentes(v): # O(E)
            if w not in visitados:
                distancias[w] = distancias[v] + 1
                if distancias[w] == n:
                    vertices_n.append(w) # todo esto es O1
                else:
                    q.append(w)
                visitados.add(w)
                
    return vertices_n

# Tipo de recorrido utilizado: BFS

# La complejidad de este algoritmo es O(V+E), ya que se realizan V veces las aristas de 
# cada vértice (que son en total todas las aristas), y dentro de esos bucles se hacen
# operaciones de tiempo constante