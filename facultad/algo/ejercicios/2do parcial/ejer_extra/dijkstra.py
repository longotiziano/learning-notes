"""
Dijkstra con aristas con peso = (tiempo, nombre) y vértices las esquinas de CABA. 

Cada nombre es una clave en el informe pasado por parámetro: "habilitada", "congestionada", "cortada".
"""
import heapq

def algo_dijkstra(grafo, informe, origen):
    padres = {} # {H: F} -> {F: A}
    distancias = {} # {H: 2}
    for v in grafo:
        distancias[v] = float("inf")
    distancias[origen] = 0
    heap = [(0, origen)]
    while heap:
        dis_v, v = heapq.heappop(heap) # dis_v: 1, v: F 
        if dis_v > distancias[v]:
            continue
        for w in grafo.adyacentes(v):
            segs, nombre = grafo.peso_arista(v, w)

            if informe[nombre] == "CORTADA":
                continue
            elif informe[nombre] == "CONGESTIONADA":
                segs *= 2

            dis_w = dis_v + segs
            if dis_w < distancias[w]:
                distancias[w] = dis_w
                padres[w] = v
                heapq.heappush(heap, (dis_w, w))
    return padres, distancias

def reconstruir_camino(padres, origen, destino):
    camino = [destino]
    actual = padres[destino]
    while actual != origen:
        camino.append(actual)
        actual = padres[actual]
    camino.append(origen)
    return camino[::-1]

def resolver(grafo, informe, origen, destino):
    padres, distancias = algo_dijkstra(grafo, informe, origen)
    return reconstruir_camino(padres, origen, destino)