# Machine Learning Models with Apache Spark.

## Introducción al Machine Learning.

### Machine Learning para todos.
- `Machine learning` (ML): Subcampo de la informática que permite a las computadoras la 
habilidad de aprender sin ser explicitamente programada.
- Funcionan a través del reconocimiento de algoritmos luego de los datos recopilados. Mientras
más tengan, más precisos serán.
* Confusiones comunes:
    - `Inteligencia artificial`: Visión, creatividad, mucho más general.
    - `Machine learning`: Mucho más centrado en la estadística con redes neuronales.
    - `Deep learning`: Como ML, pero mucho más profundo.
* Categorias de aprendizaje de ML: 
    - `Aprendizaje supervisado`: Predicción de valores continuos o categorías luego de una 
    observación.
    - `Aprendizaje no supervisado`: Buscando patrones o estructura en los datos.
- `Ramas del ML`: Deep learning, Natural language processing, Computer vision, y 
reinforcement learning (manera de enseñar que recompensa el triunfo y castiga el error).

### Roles de Ingeniería de datos en Machine Learning.
* `Responsabilidades principales` de los ingenieros de datos:
    1. `Recolección de datos`
        - Definir requisitos, evaluar fuentes, asegurar calidad...
    2. `Preparación de datos`
        - Exploración de datos, limpieza, transformación, integración y formateo.
    3. `Almacenamiento y gestión de datos`
        - Elegir almacenamiento, organizarlo, asegurarlo, optimizarlo y guardarlo.
    4. `Transformación y extracción de características`
        - Elegir las relevantes para análisis, ajustar variables a una misma escala,
        crear nuevas características, reducir la dimensionalidad, convertir datos
        categóricos en numéricos y rellenar valores faltantes.

### Ciclo de vida de un modelo Machine Learning.
- `Ciclo de vida`: Definición del problema, colección de datos, preparación de datos,
desarrollo de modelo y evaluación, y el despliegue del modelo.

### Aprendizaje supervisado contra no supervisado.
- ¿Cómo se les enseña a los modelos de ML? A través de _labeled data_. Estos datos tienen
clases o categorías explícitas. Las columnas numéricas son llamadas _features_. En ML
suelen haber únicamente 2 tipos de datos, los numéricos y las palabras o "varchars".
* Hay 2 tipos de aprendizaje supervisado:
    - `Clasificación`: Predecir la categoría
    - `Regresión`: Predecir los valores continuos.
- El `aprendizaje no supervisado` consta del uso de modelos de ML para reconocer patrones que
son irreconocibles para el humano.  
- `Diversas técnicas de aprendizaje no supervisado`: Dimension reduction, density estimation,
market basket analysis y clustering.
- `Clustering`: Agrupado dee puntos de datos o objetos que son, de alguna manera, similares.

### Regresión.
- `Regresión`: Relación entre una variable independiente y otra dependiente. La variable
dependiente varía de manera continua. Como por ejemplo, una predicción del clima.
- La fórmula de regresión es la misma que una función lineal -> `y=mx+b`
- A veces la cantidad de variables va a causar que la regresión no logre ser lineal, por
lo que en estos casos se utiliza la forma `polinómica de regresión`, que tiene la misma forma
que una cuadrática -> `y=ax1^2+bx2+c`
- `Regresión regularizada`: Se les agrega un término de penalización en la función del costo,
obligando a los coeficientes más grandes del modelo a hacer más pequeños o inclusive 0.
* `Tipos de regresión regularizada`:
    - `Ridge`: Reduce el tamaño de los coeficientes, pero no los lleva a 0.
    - `Lasso`: Penaliza los valores absolutos de los coeficientes, dando la posibilidad de
    que queden en 0, eliminando esas variables del modelo (útil para selección de variables).
    - `ElasticNet`: Combinación óptima de Ridge y Lasso.
* Algoritmos de regresión:
    - `Random forest`: Grupo de arboles de decisión combinados en un único modelo.
    - `Support vector regression` (SVR): Crea una lina o hiperplano que separa los datos en 
    clases.
    - `Gradient boosting`: Si bien también utiliza árboles de decisión...
* `Diferencias entre Random Forest (RF) y Gradient Boosting (GB)`:
    - `Construcción`: RF en paralelo de manera independiente y GB en secuencia aprendiendo
    de errores anteriores.
    - `Combinación`: Promediada en RF contra suma de errores en GB.
    - `Objetivo`: Evitar _overfitting_ en RF, y reducir sesgo en GB.
    - `Tolerancia`: RF se la banca más, a diferencia de GB que es una mami.
- `Overfitting` (sobreajustado): Cuando el modelo memoriza errores y aprende demasiado bien los 
datos del entrenamiento, causando así que memorice en lugar de aprender.
* `Diferencias entre Clasificación y regresión`:
    - Mientras que en uno los valores estan atados a una clase predefinida, en la
    regresión se trata de una variable continua.
    - En clasificación los valores no están ordenados porque no hay necesidad de ello,
    a diferencia de la regresión que si lo que requiere.
    - El rendimiento esta medido por precisión, y en la regresión esta medido por el error
    y qué tan lejos estuvieron del éxito.

### Clasificación.
- En la clasificación, únicamente hay una única respuesta, no hay 2 "categorías" a las
que el algoritmo pueda asignar simultáneamente.
- `Clasificación`: Proceso de predecir una clase basado en unos resultados recibidos.
- `Ejemplo`: ¿Voy a aprobar mi examen de biologia? -> La respuesta se puede basar en diversas
variables, como el promedio de puntuación obtenido en examenes de biología pasados, asistencias
número de horas de estudio... Lo que hace que el resultado dependa de varias variables.
* `Terminología de la clasificación`: 
    - `Clasiffier`: Algoritmo.
    - `Feature`: Variable independiente.
    - `Evaluation`: Los medios de validación.
* `Se pueden dividir los algoritmos de clasificación en dos`:
    - `Lazy learner`: No tienen una etapa de entrenamiento. Funcionan con la idea del
    _K-NN_, que basicamente asigna un valor en un plano cartesiano a cada variable, y
    en base a posiciones de cercanía realiza los calculos y por consiguiente conclusiones.
    Sus desventajas es su lentitud y la fuerte sensibilidad por aquel valor "K" (distancia entre 
    neighbors), pero su ventaja es su facilidad de implementación.
    - `Eager learner`: Pasa mucho tiempo siendo entrenado y generando el modelo. Funciona con
    la regresión logística, que es intentar deducir una función matemática en base a las
    variables de entrada y luego aplica una función sigmoide para convertir el resultado
    en una probabilidad entre 0 y 1.
    - `Decision trees`: Arboles de decisiones que utilizan reglas de if-then.

### Evaluando modelos de Machine Learning.
- Es importante no utilizar todos los datos para entrenar un modelo, si no mejor guardar
una parte de ellos para intentar dividir etapas de entrenamiento con etapas de testeo.
- `Clasificando la precisión`: Suma de buenos resultados dividido la cantidad total. 
También se puede clasificar mediante a las _confusion matrix_.
- También se debe calcular la precisión de cada uno de los casos, no únicamente la general
del total de resultados.
- Otra medida fundamental es el `Recall`, que mide que tan bien encuentra todos los casos
positivos reales. Se debe de utilizar y tener en cuenta cuando los falsos negativos son
muy costosos.
- La combinación de estas dos medidas es el `F1-score`, que es el promedio de ambas.
Se calcula a través de (2xPrecisiónxRecall) / (Precision+Recall).
* La evaluación de modelos de regresión se puede hacer de diversas formas:
    - `Mean squared error` (MSE): Promedio de diferencia entre la predicción y el verdadero 
    output. 
    - `Root mean squared error`: Raiz cuadrada del MSE
    - `Mean Absolute Error`: Promedio del valor absoluto del error.
    - `R-squared o coefficient of determination`: Cantidad de variación en la variable 
    dependiente que puede ser explicada por la variable independiente. Mide que tan bueno es el 
    modelo, de 0 a 1.

### Clustering.
- El `clustering` agrupa datos mediante el reconocimiento de patrones o conexiones, machine
learning no supervisado e identifica grupos basados en similaridad o distancia.
- `Ejemplos de clustering`: segmentación de imagenes (basados en color y contenido), clientes 
(basado en un historial de compras o revisiones), detección de anomalías (puntos inusuales),
sistemas de recomendación...
* `Tipos de clustering`:
    - `Partitioning clustering`: El dataset está dividido en _k_ particiones o clusters,
    donde _k_ es especificado por el usuario e incluye el _k-mean_ y el _k-medoids_.
    - `Hierarchical clustering`: Los datasets son divididos en clusters por una jerarquía.
    Cada jerarquía representa un nivel diferente de granularidad, incluyendo agglomerative
    clustering y divisive clustering. 
    - `Density-based clustering`: Define los clusters como regiones densas de datos, 
    separados por regiones de menor densidad. Incluye DBSCAN y OPTICS.
* El algoritmo de k-means sigue los pasos de:  
    1. Seleccionar k centroides iniciales aleatoriamente.
    2. Asignar cada punto al centroide más cercano (usando distancia euclidiana).
    3. Recalcular los centroides como el promedio de los puntos asignados.
    4. Repetir los pasos 2 y 3 hasta que las asignaciones no cambien.
- El resultado son k clústeres, con cada punto asignado al más cercano.

### Inteligencia artificial generativa y casos de uso.
- `Inteligencia artificial` (AI): Inteligencia aumentada que permite a los expertos escalar sus 
capacidades.
- `Inteligencia artificial generativa`: Es una técnica de IA que puede generar únicos y nuevos
datos. Utiliza técnicas de deep learning basándose en enormes datasets.
- `IA generativa y LLM`: El Large Language Model (LLM) es un tipo de IA que utiliza técnicas
de deep learning para procesar y generar lenguaje natural.
- `Beneficios de la IA generativa`: Creatividad e innovación, ahorro de tiempo y costos,
personalización, escalabilidad, robustez y exploración de nuevas posibilidades.
- Cosas a saber interesantes: GPT significa Generative Pre-trained Transformers, que es
un LLM desarrollado por OpenAI, y ChatGPT es el chatbot; Bard, que genera texto utilizando 
LaMBDA y se ajusta al estilo y tono del usuario; DeepDream, que es una IA para generar
imagenes de la vida real; StyleGAN, para crear imagenes ficticias...

## Machine Learning con Spark.

### Regresión con SparkML.
- Spark cuenta con una librería llamada `Spark MLlib`, la cual sirve para contruir y entrenar
modelos de regresión. Es escalable, distribuida y se integra bien al ecosistema de Spark.
* Pasos para la regresión con Spark ML:
    1. Importar las librerias necesarias:
    ```python
    from pyspark.sql import SparkSession
    from pyspark.ml.feature import VectorAssembler
    from pyspark.ml.regression import LinearRegression
    from pyspark.ml.evaluation import RegressionEvaluator
    ```
    2. Crear la sesión de Spark.
    ```python
    spark = SparkSession.builder.appName("LinearRegressionExample").getOrCreate()
    ```
    3. Leer los datos de un CSV.
    ```python
    data = spark.read.csv("path", header=True, inferSchema=True)
    ```
    4. Seleccionar las columnas relevantes.
    ```python
    selected_data = data.select("col1", "col2", "target_col")
    ```
    5. Crear un `Vector Assembler` (sirve para transformar múltiples columnas de características
    en un solo vector numérico).
    ```python
    assembler = VectorAssembler(inputCols=[], outputCol="output")
    ```
    6. Transformar los datos seleccionados utilizando el Vector Assembler.
    ```python
    transformed_data = assembler.transform(selected_data)
    ```
    7. Dividir los datos transformados en un sets de entrenamiento y testeo.
    ```python
    train_data, test_data = transformed_data.randomSplit([0.7, 0.3], seed=123)
    ```
    8. Crear una instancia del modelo de regresión linear.
    ```python
    lr = LinearRegression(featuresCol="features", labelCol="target_col")
    ```
    9. Entrená el modelo con los datos.
    ```python
    model = lr.fit(train_data)
    ```
    10. Haz predicciones en los datos de testeo.
    ```python
    predictions - model.transform(test_data)
    ```
    11. Evaluá el modelo.
    ```python
    evaluator = RegressionEvaluator(labelCol="target_col", predictionCol="prediction",metricName="rmse")
    rmse = evaluator.evaluate(predictions)
    ```
    12. Mostrá el RMSE.
    13. Mostrá los coeficientes e interceptá el modelo.
    14. Frená la sesión de Spark.

### Clasificación con SparkML.
* Los pasos para realizarla son los mismos que en regresión, con algunas modificaciones:
    1. En el primer paso, agregamos librerías:
    ```python
    from pyspark.ml.classification import LogisticRegression
    from pyspark.ml.classification import MulticlassClassificationEvaluator
    ```
    7. Crear el modelo de regresión logística:
    ```python
    lr = LogisticRegression(labelCol="label", featuresCol='features')
    ```
    10. Evaluar el rendimiento del modelo de clasificación:
    ```python
    evaluator = MulticlassClassificationEvaluator(labelCol=label, metricName='accuracy')
    accuracy = evaluator.evaluate(predictions)
    ```

### Clustering en SparkML.
* Los pasos para realizarla son los mismos que en regresión, con algunas modificaciones:
    1. En el primer paso, agregamos librerías:
    ```python
    from pyspark.ml.clustering import KMeans
    ```
    5. Entrenar el modelo _K-means_
    ```python
    kmeans = KMeans(k=3)
    model = kmeans.fit(data)
    ```

### GraphFrames en Apache Spark.
- `Graph Theory`: Estudio matemático de modelado _pairwise_ en sus relaciones entre
objetos, con vértices y aristas.
- Los graphs pueden ser `direccionados`, donde las vértices y las aristas tienen una única
dirección; o `no direccionados`, donde las aristas no tienen una dirección definida.
- Spark tiene una extensión llamada `GraphFrames`, la cual está hecha para procesamiento
de gráficos. Está basada en los DataFrames de Spark y contiene algoritmos pre-construidos.
Te permiten hacer consultas SQL sobre los datos que representan los vértices y las aristas
del grafo. 
- Para el correcto funcionamiento de los grafos, es necesario definir un directorio para los
checkpoints.
- Los GraphFrames pueden hacer "motif finding", que es buscar patrones estructurales dentro
del grafo.
* Algoritmos de los grafos soportados:
    - Breadth-first search
    - Connected components
    - Label Propagation Algorithm
    - PageRank
    - Shortest paths
    - Triangle count

## Ingenieria de datos para Machine Learning con Spark.

### Spark SQL.
* Pasos para correr SparkSQL:
    1. Importar librerias necesarias
    ```python
    from pyspark.sql import SparkSession
    ```
    2. Crear la SparkSession (ya sabemos).
    3. Cargar datos en un DataFrame (ya sabemos).
    4. Registrar los datos en una vista temporal, esto para poder consultar los datos
    utilizando SQL.
    ```python
    data.createOrReplaceTempView("my_table")
    ```
    5. Hacer consultas.
    ```python
    result = spark.sql("CONSULTA SQL")
    ```
    6. Mostrar resultados.
    ```python
    result.show()
    ```
    7. Frenar la sesión de Spark.

### Extracción de características y transformación.
- `Data preprocessing`: Incluye limpieza, transformación y preparación de datos crudos que
serán utilizados para tareas analíticas y de machine learning. Este proceso incluye 
extracción de características y transformación.
- `Extracción de características`: Seleccionar características relevantes que asisten al
análisis y modelado de tareas. Incluye técnicas como PCA, factor analysis o feature engineering.
* Para este último proceso mencionado, Spark MLlib ofrece: 
    - `Tokenzation`: Divide textos en palabras individuales o tokens y utiliza la clase
    Tokenizer.
    - `TF-IDF`: Identifica palabras importantes en un documento de texto.
    - `Word2Vec`: Representa palabras en un documento de texto como vectores, utilizando la
    clase Word2Vec.
    - `Principal Component Analysis` (PCA): Identifica un pequeño conjunto de características 
    para explicar la varianza, reduciendo la dimensión de los datos. Utiliza la clase PCA.
- `Transformación de características`: Convierte las características extraídas en un formato
como la gente, incluyendo modificación y manipulación de las características originales.
Basicamente una transformación en el mundo de los datos cualquiera.
* Para la transformación, encontramos las siguientes técnicas y funciones:
    - `Escalamiento y normalización`: Con funciones como StandardScaler, MinMaxScaler, 
    MaxAbsScaler...
    - `One-hot encoding`: Convierte características categóricas en numéricas. Utiliza
    la función OneHotEncoder.
    - `StringIndexer`: Convierte características categóricas en numéricas y asigna un índice
    numérico a cada categoría distinta. Utiliza la función StringIndexer.
    - `PCA`: Ya explicado.

### Pipelines de Machine Learning utilizando Spark.
* Los pasos para llevar a cabo un pipeline de Machine Learning son:
    1. Colección de datos e ingestión.
    2. Limpieza y preprocesado.
    3. Extracción de características.
    4. Selección de modelo y entrenamiento.
    5. Evaluación del modelo.
    6. Despliegue del modelo.

### Persistencia del modelo.
- `Persistencia del modelo`: Se basa en almacenar el modelo en disco para uso futuro, 
permitiendo así su reutilización. Esto tiene muchas ventajas, como reducir costos computacionales
y tiempo y la necesidad de entrenar repetidas veces un modelo. También permite la portabilidad,
aumenta en gran medida la eficiencia, flexibilidad y la escalabilidad.










