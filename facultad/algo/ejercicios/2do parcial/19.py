from grafo import Grafo # type: ignore
from collections import deque

def crear_vertice_inex(grafo, ingresados, c):
    if c not in ingresados:
        grafo.agregar_vertice(c)
        ingresados.add(c)

def crear_grafo_abc(palabras):
    grafo = Grafo(es_dirigido=True, vertices_init=[])
    ingresados = set()

    for palabra in palabras:
        for letra in palabra:
            crear_vertice_inex(grafo, ingresados, letra)

    for i in range(1, len(palabras)):
        p1 = palabras[i-1]
        p2 = palabras[i]

        min_len = min(len(p1), len(p2))

        for j in range(min_len):
            c1 = p1[j]
            c2 = p2[j]

            if c1 != c2:
                if not grafo.estan_unidos(c1, c2):
                    grafo.agregar_arista(c1, c2)
                break

    return grafo

def grados_entrada(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados

def abecedario_alien(palabras):
    grafo = crear_grafo_abc(palabras)
    grados = grados_entrada(grafo)
    q = deque()
    orden = []
    for v in grafo:
        if grados[v] == 0:
            q.append(v)
    while q:
        v = q.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0:
                q.append(w)
    if len(orden) == len(grafo):
        return orden

    
