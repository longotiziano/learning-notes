"""
encontrar el camino más corto entre v y w, pasando si o si por z

mi idea es reutilizar la funcion camino_forzado primero teniendo como origen a v, y despues cuando consiga la minima distancia a z,
busco la distancia de z a w. Aunque esto dependería que está más cerca, por lo tanto, si w es más cercano a v que z, entonces tendría
que hacer el recorrido v -> w -> z. Por lo tanto, cuando consiga el diccionario de distancias respecto a z, Tambien habría que evaluar
las distancias de z a w.

hay un camino minimo de V -> W
hay un camino minimo de W -> Z
hay un camino minimo de V -> Z

primero busco V -> W < V -> Z
si es así, entonces construyo el camino minimo de v a w y dsp de w a z
"""
import heapq

def dijkstra(grafo, origen):
    padres = {origen: None}
    distancias = {v: float("inf") for v in grafo}
    distancias[origen] = 0
    heap = [(0, origen)]  # Distancia PRIMERO para que ordene bien
    
    while heap:
        dis_v, v = heapq.heappop(heap)
        if dis_v > distancias[v]:
            continue
        for w in grafo.adyacentes(v):
            dis_w = dis_v + grafo.peso_arista(v, w)
            if dis_w < distancias[w]:
                padres[w] = v
                distancias[w] = dis_w
                heapq.heappush(heap, (dis_w, w))
                
    return padres, distancias

def camino_forzado(grafo, v, w, z):
    # 1. Camino mínimo de v a z
    padres_v, distancias_v = dijkstra(grafo, v)
    if z not in padres_v:
        return None  # No se puede llegar a z desde v
        
    # 2. Camino mínimo de z a w
    padres_z, distancias_z = dijkstra(grafo, z)
    if w not in padres_z:
        return None  # No se puede llegar a w desde z
        
    camino1 = reconstruir_camino(padres_v, v, z)
    camino2 = reconstruir_camino(padres_z, z, w)
    
    # Unimos evitando duplicar z
    camino1.extend(camino2[1:])
    return camino1

def reconstruir_camino(padres, origen, destino):
    if origen == destino:
        return [origen]
    camino = []
    actual = destino
    while actual is not None:
        camino.append(actual)
        actual = padres[actual]
    return camino[::-1]