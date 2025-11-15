# Introducción al análisis de datos

## Tipos de análisis de datos
1. **Análisis descriptivo**: Responde a la pregunta "¿qué pasó?" y resume eventos observados.
2. **Análisis diagnóstico**: Investiga por qué ocurrió algo (causas y relaciones) usando
    datos históricos y técnicas de contraste.
3. **Análisis predictivo**: Estima lo que podría ocurrir en el futuro mediante modelos
    estadísticos o de machine learning.
4. **Análisis prescriptivo**: Recomienda acciones concretas (reglas, optimización, simulaciones)
    para alcanzar objetivos, basándose en resultados predictivos y restricciones del negocio.

## Proceso del análisis de datos
1. Definir el problema y el objetivo de negocio o investigación.
2. Establecer métricas y criterios de éxito (qué vamos a medir y cómo).
3. Identificar y evaluar fuentes de datos, requisitos de calidad y herramientas necesarias.
    (El orden puede variar: en proyectos iterativos a veces se prototipa antes de fijar todas
    las métricas.)
4. Obtener y versionar los datos (captura, extracción, y registro de versiones para
    reproducibilidad).
5. Limpieza y preprocesamiento: tratar valores faltantes, duplicados, tipos, formato y errores.
6. Análisis exploratorio: visualizar, resumir y buscar patrones, correlaciones y anomalías.
7. Modelado y validación: construir modelos, evaluar su rendimiento y comprobar supuestos.
8. Interpretación y validación de resultados: evaluar limitaciones, robustez y contraejemplos.
9. Comunicación y visualización: presentar hallazgos y recomendaciones claras y reproducibles.

## Responsabilidades del analista de datos
- Adquirir datos, diseñar y ejecutar consultas, filtrar, limpiar y estandarizar, transformar,
    analizar patrones, preparar reportes y visualizaciones, y documentar el proceso.
- Combinación de habilidades:
    - **Técnicas / Herramientas**: hojas de cálculo, herramientas de visualización, programación
        (Python, R), SQL, APIs, plataformas de procesamiento (Spark, pandas), y conceptos de Big Data.
    - **Analíticas / Funcionales**: estadística, modelado, diseño de experimentos, definición de
        métricas y evaluación de resultados.
    - **Habilidades blandas**: comunicación efectiva, storytelling con datos, colaboración,
        curiosidad e intuición para plantear hipótesis.

## Día de un analista de datos
- Al momento de resolver una problemática, es importante y útil plantear las hipótesis iniciales.
- Siempre documentar el proceso para poder mostrarlo de manera efectiva y convincente.

## Identificación de datos para análisis
1. Determinar la información que se necesita
   - Variables o indicadores clave.
   - Posibles fuentes (internas, externas, APIs, sensores, encuestas, etc.).
2. Definir el plan de obtención
   - Cronograma y prioridades.
   - Criterio de cantidad suficiente (tamaño de muestra) y calidad mínima.
   - Dependencias, riesgos y plan de mitigación.
3. Definir métodos y requisitos
   - Tipos de datos (estructurados, semiestructurados, no estructurados).
   - Frecuencia y ventana temporal.
   - Estimación de volumen y costes.

Tras ejecutar el plan se iterará: los requisitos suelen cambiar a medida que se obtienen
insights.

Los datos identificados implican consideraciones sobre:
1. **Calidad**: precisión, completitud, consistencia y relevancia.
2. **Seguridad**: cumplimiento normativo, controles de acceso y cifrado.
3. **Privacidad**: consentimiento, anonimización, licencias y auditoría.

## Fuentes de datos
- **Datos primarios**: Información obtenida directamente de la fuente (entrevistas, encuestas,
    observaciones, mediciones). Son los datos originales recogidos para el propósito del estudio.
- **Datos secundarios**: Información ya existente (bases de datos internas, informes, artículos,
    datasets públicos). Se reutilizan datos generados por terceros.
- **Datos de terceros**: Datasets comercializados o agregados por proveedores externos. Revisar
    licencias, calidad y posibles sesgos.

## Data wrangling
El data wrangling (preprocesamiento) suele incluir cuatro fases principales:
1. **Descubrimiento**: explorar y perfilar los datos, entender su estructura y calidad.
2. **Transformación**: limpieza, normalización, conversión de tipos y enriquecimiento.
3. **Validación**: comprobación de reglas, consistencia, y control de calidad.
4. **Publicación**: exportar datos preparados (APIs, datasets, informes) con control de acceso.

Documentar cada paso es esencial para la reproducibilidad y el trazo de decisiones.

## Normalización y desnormalización
- **Normalización**: Diseño relacional que divide información en tablas coherentes para reducir
    redundancia y evitar anomalías de inserción, borrado o actualización. Se aplica mediante formas
    normales (1NF, 2NF, 3NF, BCNF) cuando corresponde.
- **Desnormalización**: Duplicar o combinar estructuras para mejorar rendimiento de lectura,
    consultas analíticas o simplicidad en entornos OLAP/BI. Es una decisión de rendimiento y
    costes de almacenamiento.

## Proceso de limpieza de datos
En la limpieza de datos suelen reconocerse tres etapas principales:
1. **Inspección / Perfilado**: detectar errores, valores atípicos y patrones de calidad;
    validar contra reglas y restricciones.
2. **Limpieza**: aplicar técnicas según el problema:
    - Tratar valores nulos: rellenar desde la fuente, imputar (media, modelos), o marcarlos.
    - Duplicados: normalmente se deduplican, salvo que los duplicados sean relevantes para el
      dominio (por ejemplo, versiones históricas).
    - Conversión y estandarización de tipos y formatos.
    - Corrección de errores sintácticos o de codificación.
    - Tratamiento de outliers: investigar antes de eliminar.
3. **Validación**: comprobar que las transformaciones respetan reglas de negocio y mantienen
    la integridad esperada.

Documentar cada cambio y conservar copias/versiones de los datos originales.

## Análisis estadístico
- **Estadística**: rama de las matemáticas que estudia la recolección, análisis,
    interpretación y presentación de datos.
- **Análisis estadístico**: aplicación práctica de la estadística para resumir y extraer
    conclusiones sobre datos.

**Tipos de estadística**:
- **Estadística descriptiva**: resume y describe las características de un conjunto de datos
    (medidas de tendencia central, dispersión y visualizaciones). Sirve para describir los
    datos observados, no necesariamente para generalizar a poblaciones sin metodología adicional.
- **Estadística inferencial**: utiliza muestras para hacer generalizaciones o predicciones
    sobre una población (tests de hipótesis, intervalos de confianza, modelos de regresión).

`Mediana`: valor que divide los datos ordenados en dos mitades.

`Dispersión`: medida de la variabilidad en un dataset:
- Varianza: promedio de las desviaciones al cuadrado respecto a la media.
- Desviación estándar: raíz cuadrada de la varianza (unidad igual a la del dato).
- Rango: diferencia entre el máximo y el mínimo.
- Asimetría (skewness): indica si la distribución está sesgada hacia un lado.

## Data Mining
Data mining es la extracción de conocimiento a partir de grandes volúmenes de datos.
Es interdisciplinario y usa técnicas de estadística, machine learning y reconocimiento de
patrones para identificar correlaciones, tendencias y estructuras útiles.

**Técnicas comunes de data mining**:
1. Clasificación: asignar etiquetas o categorías a ejemplos (ej. árboles, SVM, redes).
2. Clustering: agrupar observaciones similares sin etiquetas previas (ej. k-means, DBSCAN).
3. Detección de anomalías: identificar comportamientos inusuales o outliers.
4. Reglas de asociación: descubrir relaciones frecuentes entre ítems (ej. market basket).
5. Patrones secuenciales: encontrar secuencias frecuentes en eventos temporales.
6. Modelado predictivo / Regresión: predecir valores continuos o probabilidades.
7. _Affinitty grouping_: encontrar patrones de asociación entre elementos.
8. _Decision trees_: Modelos de clasificación ya vistos.

**Herramientas del data mining**
- Spreadsheets: Son fáciles de leer, y con agregados como Data Mining Client o XMiner que 
permiten realizar tareas simples como regresión o clasificación.
- R: Uno de los lenguajes más usados al momento de realizar modelos y computaciones estadísticos.
- Python
- Jupyter Notebooks
- IBM SPSS (Statistical Package for Social Sciences) Statistics: Closed-source
- IBM Watson Studio
- SAS: SAS Enterprise Miner es espacio de trabajo completo y gráfico de minería de datos

Buenas prácticas: validar modelos (cross-validation), usar métricas adecuadas, vigilar
sesgos y evitar overfitting.