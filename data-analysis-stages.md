# Data Analysis Workflow – Pipeline Realista

Este documento resume el flujo técnico correcto de trabajo en análisis de datos, diferenciando claramente cada etapa del proceso para evitar confusiones comunes en cursos introductorios.

---

## 1. EDA Inicial (Exploratory Data Analysis)

Objetivo: entender el dataset y diagnosticar problemas antes de limpiar.

Tareas comunes:
- Inspección de estructura de tablas
- Análisis de tipos de datos
- Detección de valores faltantes
- Identificación de outliers
- Revisión de distribuciones
- Visualizaciones rápidas
- Detección de incoherencias
- Evaluación de calidad general de los datos

Resultado: lista clara de problemas y tareas futuras de limpieza.

---

## 2. Data Wrangling

Objetivo: ejecutar la limpieza y transformación de datos detectadas en EDA.

Tareas comunes:
- Conversión de tipos de datos
- Eliminación de duplicados
- Manejo de valores nulos
- Normalización de formatos
- Corrección de errores
- Filtrado de registros
- Unión de tablas
- Estandarización de categorías

Resultado: dataset usable, consistente y confiable.

---

## 3. EDA Post-Wrangling

Objetivo: validar que los cambios realizados no introdujeron errores.

Tareas comunes:
- Reanálisis de distribuciones
- Chequeo de consistencia
- Verificación de duplicados
- Revisión de nuevos nulos
- Validación de relaciones entre campos

Resultado: dataset validado para análisis real.

---

## 4. Feature Engineering

Objetivo: generar valor analítico a partir de datos crudos.

Tareas comunes:
- Creación de columnas derivadas
- Cálculo de métricas
- Agregaciones
- Encoding de variables categóricas
- Discretización de variables
- Normalización
- Transformaciones lógicas

Resultado: dataset enriquecido para análisis.

---

## 5. Análisis

Objetivo: extraer insights reales de los datos.

Tareas comunes:
- Detección de patrones
- Análisis de correlaciones
- Segmentación
- Análisis exploratorio profundo
- Validación de hipótesis

Resultado: comprensión profunda del comportamiento de los datos.

---

## 6. Reporte y Comunicación

Objetivo: transformar análisis en decisiones.

Tareas comunes:
- Construcción de dashboards
- Visualización clara de datos
- Storytelling
- Generación de conclusiones
- Recomendaciones accionables

Resultado: impacto real en decisiones.

---

## Notas importantes

- El análisis NO empieza limpiando, empieza entendiendo los datos.
- Data Wrangling NO incluye análisis.
- Feature Engineering no es limpiar: es crear valor.
- Si los resultados no se comunican, el análisis no existe.

---

## Frase profesional de referencia

> I perform initial EDA to diagnose data issues, execute data wrangling, apply feature engineering, and then focus on analysis and reporting.

---

Documento creado como referencia personal para formación en Data Analytics.
