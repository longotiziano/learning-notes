# Detección de Outliers

## ¿Qué es un outlier?
Un outlier es un valor que se aleja significativamente del comportamiento general de una variable. Puede deberse a errores de carga, medidas atípicas o eventos excepcionales.

## ¿Por qué se usan boxplots para detectar outliers?
- No requieren que la variable tenga distribución normal.
- Se basan en medidas robustas como cuartiles e IQR.
- Funcionan bien con datos reales y ruidosos.
- Son fáciles de interpretar visualmente.
- Método IQR (Interquartile Range)

---

### El boxplot utiliza el rango intercuartílico:
**IQR (Interquartile Range) = Rango Intercuartílico**

Que es simplemente la diferencia entre dos números importantes:
- Q1 -> el valor debajo del cual está el 25% de los datos
- Q3 -> el valor debajo del cual está el 75% de los datos

```
IQR = Q3 - Q1
```

Valores fuera de esos límites **se consideran outliers.**

---

### Ventajas del método IQR
- Robusto frente a valores extremos.
- No depende de la forma de la distribución.
- Es el estándar más utilizado en análisis exploratorio de datos.

**Cálculo de outliers en Python**
```python
Q1 = df[col].quantile(0.25)
Q3 = df[col].quantile(0.75)
IQR = Q3 - Q1

lim_inf = Q1 - 1.5 * IQR
lim_sup = Q3 + 1.5 * IQR

outliers = df[(df[col] < lim_inf) | (df[col] > lim_sup)]
```
---

**Visualización con boxplot**

```python
import matplotlib.pyplot as plt

plt.figure(figsize=(5, 6))
plt.boxplot(df[col])
plt.title(f"Boxplot de {col}")
plt.ylabel(col)
plt.show()
```