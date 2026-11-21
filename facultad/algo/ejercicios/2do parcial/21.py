"""
(★★★) 
♠
♠ Contamos con un grafo dirigido que modela un ecosistema. En dicho grafo, cada vértice 
es una especie, y cada arista (v, w) indica que v es depredador natural de w. 
Considerando la horrible tendencia del ser humano por llevar a la extinción especies, 
algo que nos puede interesar es saber si existe alguna especie que, si llegara a desaparecer, 
rompería todo el ecosistema: quienes la depredan no tienen un sustituto (y, por ende, 
pueden desaparecer también) y/o quienes eran depredados por esta ya no tienen amenazas, 
por lo que crecerán descontroladamente. 

Implementar un algoritmo que reciba un grafo de dichas características y devuelva una lista 
de todas las especies que cumplan lo antes mencionado. 

Indicar y justificar la complejidad del algoritmo implementado.
"""

# Lo que quiero hacer es encontrar aquellos puntos (x, v, w) tal que:
# x tiene grado salida 1
# w tiene grado entrada 1

def grados_entrada_salida(grafo): # O(V+E)
    grados_entrada = {}
    grados_salida = {}
    for v in grafo: # O(V)
        grados_entrada[v] = 0
        grados_salida[v] = 0
    for v in grafo: # total O(V+E)
        for w in grafo.adyacentes(v): 
            grados_entrada[w] += 1
            grados_salida[v] += 1
    return grados_entrada, grados_salida

def amenazados(grafo): # O(V+E)
    especies = []
    vistos = set()
    g_entrada, g_salida = grados_entrada_salida(grafo)
    for x, gs in g_salida.items():
        if gs == 1:
            v = grafo.adyacentes(x)[0]
            if v not in vistos:
                especies.append(v)
                vistos.add(v)

    for v in grafo: # total O(V+E)
        if v not in vistos: # O(1) promedio
            for w in grafo.adyacentes(v): 
                if g_entrada[w] == 1:
                    especies.append(v)
                    vistos.add(v)
                    break

    return especies
