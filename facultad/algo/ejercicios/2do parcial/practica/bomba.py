from collections import deque

def bomba(grafo, obj, n):
    afectados = []
    vistos = set(obj)
    q = deque()
    q.append((obj, 0))
    while q:
        v, dist = q.popleft()
        afectados.append(v)
        if dist == n:
            continue
        for w in grafo.adyacentes(v):
            if w not in vistos:
                vistos.add(w)
                q.append((w, 1 + dist))
    return afectados