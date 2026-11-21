# determinar si hay al menos un v tal que sus adyacentes poseen una adyacencia entre sí
def es_algo_buy(grafo):
    vistos = set()
    es_buy = True
    for v in grafo: # O(V)
        if not es_buy or v in vistos:
            break
        ady_v = set(grafo.adyacentes(v))
        vistos.add(v)
        for w in ady_v: # O(V + E)
            if not es_buy or w in vistos:
                break
            for x in grafo.adyacentes(w):
                if x in ady_v:
                    break
    return es_buy

def algo(grafo1, grafo2):
    g1 = es_algo_buy(grafo1)
    g2 = es_algo_buy(grafo2)
    res_buy = "AlgoBuy"
    res_friends = "AlgoFriends"
    if g1 and not g2:
        return res_buy, res_friends
    if not g1 and g2:
        return res_friends, res_buy
    return None, None
