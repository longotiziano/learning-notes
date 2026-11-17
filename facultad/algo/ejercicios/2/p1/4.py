def grados_entrada_salida(grafo): # O(V + E)
    entrada = {}
    salida = {}
    for v in grafo:
        entrada[v] = 0
        salida[v] = 0

    for v in grafo:
        for w in grafo.adyacentes(v):
            entrada[w] += 1
            salida[v] += 1
    return entrada, salida

def es_scatter_gatther(grafo, n): # Cada vértice lo visito hasta 3 vecesy q tiene
    entrada, salida = grados_entrada_salida(grafo) # O(V+E)
    pares = []
    for v, grado_salida in salida.items(): # O(V)
        if grado_salida >= n:
            contador_sumidero = {}
            for w in grafo.adyacentes(v): # Por cada vértice, veo las aristas de sus adyacentes
                if w == v:
                    continue
                for x in grafo.adyacentes(w):
                    if x != v and x != w and entrada[x] >= n:
                        contador_sumidero[x] = contador_sumidero.get(x, 0) + 1
            for sumidero, apariciones in contador_sumidero.items():
                if apariciones >= n:
                    pares.append((v, sumidero))
    return pares

