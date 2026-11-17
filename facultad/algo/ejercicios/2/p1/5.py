"""
a. Esto es verdadero para todo grafo sin pesos negativos (o ciclos con pesos resultantes negativos, en caso de grafos dirigidos)
b. Falso, ya que el MST lo que busca es unificar todos los vertices de la manera más óptima (en cuanto a pesos), dejándolos 
en una misma componente conexa.

En el contraejemplo de la imagen, si queremos el camino más corto de (a, d) tenemos que el peso es 6, pero a través del MST 
sería 9 (ya que intenta conectar todos los vértices de la manera más barata).
c. Considerando "Sumar C a todas las aristas" como "aumentar en C el peso de las aristas" y mientras C sea positiva, si: nunca
disminuye la cantida de caminos mínimos posibles.
"""