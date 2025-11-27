# Preparing Data for Analysis with Microsoft Excel

### Course introduction
- El proceso del análisis de datos se basa en
```
Defining -> Gathering -> Cleaning -> Analyzing -> Reporting
```
 
---

## Referente a Generative AI
Encontramos diversos términos, entre ellos:

### Artifical Intelligence (AI)
Campo de la informática que se concentra en crear sistemas que son capaces de performar tareas que tienden a requerir inteligencia humana.

### Machine Learning (ML) 
Concetra los modelos que se basan en el desarrollo de modelos estadísticos y algorítmicos. Estos mismos modelos requieren de una preparación de datos robusta y buena, ya que la calidad de los mismos afectará de manera directa a los resultados del modelo.

**Tenemos variedad de tipos de aprendizaje**:
- Supervizado: Tratan datos _labeleados_
- No-supervizado: No requieren las etiquetas (clustering y asosiación)
- Reinforcement: Aprenden a crear secuencias basadas en feedback

**Técnicas y algorítmos**:
- Lineas de regresión, árboles de decisión, neural networks...
- Deep learning

### Deep Learning (DL)
El Deep Learning es una técnica avanzada del ML que utiliza redes neuronales artificiales con múltiples capas (deep neural networks). Ejemplos: Generative adversarial networks (GANs)

**Tipos de aprendizaje**:
- Supervizado: _Labeled_
- Semi-supervizado: Mix de _labeled_ y _unlabeled_
- No-supervizado: Únicamente datos _unlabeled_

### Neural networks
Estructura inspirada en el cerebro humano, donde son múltiples capas de nodos interconectados de manera compleja por su medida, que es acorde a su importancia. Mientras que los datos van pasando, esas mismas medidas van ajustandosé, y ahí es donde el modelo "aprende".

Por supuesto que las _neural networks_ son parte del Deep Learning.

### GANs
En este tipo de modelos encontramos 2 redes neuronales compitiendo entre ellas, donde una genera datos indistinguibles y la otra los discrimina, evaluando consistencias.

### Natural language processing (NLP)
Modelo que permite el reconocimiento del lenguaje natural humano.

### Transformers
Avance significativo en el DL, donde los mismos utilizan el mecanismo conocido como "_self-attention_", el cual le permite al modelo tomar decisiones en base a un contexto.

### Generative pre-trained transformers (GPT)
Modelos que son pre-entrenados con un diverso rango de texto de internet, que utilizan tanto arquitectura de _transformers_ como también NLPs.

### Tokenization, Word2vec & BERT
Convierten palabras y diversos datos de manera tal que sean digeribles para los modelos.
- Tokenization: Convierte en tokens
- Word2vec: El nombre lo dice, palabras a vectores

### Bidirectional Encoder Representations from Transformers (BERT) 
Es un modelo que entiende texto como lo entiende una persona, teniendo en cuenta el contexto de antes y después de cada palabra.

---

## Introducción al manejo de Excel
En esta sección voy a ir tildando cosas que no sabía o a lo mejor no tenía tan claras.
- Para configuración de fechas, tipado, idioma y demás se debe ir a `Archivo -> Más -> Opciones -> Avanzado`.
- Los archivos de Excel son llamados **workbooks**, compuestos por **worksheets/sheets**.
- Hasta un millón de filas tiran.
- Se pueden ocultar filas/columnas que no van a ser requeridas para el momento (pestaña `Home`). 
- Para volver a mostrarlas, debes seleccionar el área que la contiene y hacer clic derecho en `Unhide Columns/Rows`. Ej: Oculto G, y selecciono desde F hasta H.
- Los formatos de datos no únicamente son útiles para quien escribe información o analiza profundamente, si no también para quien de afuera los ve, sirviendo así como guía.
- Se pueden referenciar datos de otros archivos Excel en una celda
- Para agregar una referencia absoluta, agregar `$` al prinicpio de la columna o letra.

### Features de lectura: 
- Freeze: Se queda un bloque de información estático mientras scrolleás.
- Window: División del worksheet para trabajar con 2 pantallas simultáneas.

### Shortcuts
- Ctrl + Home/End: Te lleva a la primera o última celda con contenido de la sheet. Si durante esta acción se mantiene el Shift previamente se seleccionará todo aquello que tenga contenido.
- F4: Asigna el valor absoluto de una celda (`$`)
- Para utilizar el shortcut de doble clickear una fórmula para replicarla debajo, la celda se guía en la columna que tiene a su izquierda, donde ello sería lo que le pone el límite a la escritura.

### Filtering & Sorting
Estas funcionalidades se aplican sobre toda la columna en la cual el cursor esté seleccionado una única celda.

**-> Multi-sorting**

Para realizar multi-sorting se debe entrar a la sección de `Sort`, marcar la casilla "My data has headers". Ahí el proceso es bastante intuitivo, vas agregando columnas y configurando los órdenes.

**-> Filtrado**

Simplemente presionas el botón de `Filter` y se desplegarán flechas que permiten el filtrado inteligente de las columnas.

## Manejo de fórmulas en Excel

- **Operadores**: Símbolos que realizan cálculos (+, -, /, *), con el order jerárquico común de la matemática (división, multiplicación, suma y resta)
- Se pueden calcular **porcentajes** utilizando nomás que `X%`. Ej: `A2*10%` (y daría el 10% del valor)
- En las _type hints_ de las funciones de Excel, aquellos argumentos que aparecen con corchetes "[]", entonces significa que son **opcionales**, en el caso contrario son obligatorios.
- En cuanto a las funciones, podemos ir a la sección de `Fórmulas -> Insertar función` o a la `Barra de fórmulas -> FX` y allí podrás encontrar todas las funciones disponibles, con posibilidad de filtrar por categorías.\
Esta manera de usar las funciones es super intuitiva y fácil.
- Según el curso, si en AVERAGE seleccionamos rango de X a Y, si en Y hay texto, entonces toma de X a Y-1 automáticamente.

**Funciones COUNT**
- COUNT: Cuenta si hay valores numéricos.
- COUNTA: Cuenta si hay cualquier tipo de contenido
- COUNTBLANK: Cuenta las celdas vacías  