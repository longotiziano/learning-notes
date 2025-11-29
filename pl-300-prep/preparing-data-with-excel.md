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

### Shortcuts / Calidad de vida
- Ctrl + Home/End: Te lleva a la primera o última celda con contenido de la sheet. Si durante esta acción se mantiene el Shift previamente se seleccionará todo aquello que tenga contenido.
- F4: Asigna el valor absoluto de una celda (`$`)
- Para utilizar el shortcut de doble clickear una fórmula para replicarla debajo, la celda se guía en la columna que tiene a su izquierda, donde ello sería lo que le pone el límite a la escritura.
- Para eliminar las fórmulas dejando únicamente los valores, se pueden copiar todas las celdas y pegar únicamente los valores en una misma acción.

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

## Preparando datos en Excel para el análisis
En esta sección se abarca la preparación de datos en Excel.

### Errores comunes
Podemos identificar errores comunes a la hora de limpiar datos, tales como:
- Errores de tipeo
- Carácteres y espacios inneccesarios
- Ubicación de información errónea
- Incosistencias en el esquema de la worksheet
- Abreviaciones y acrónimos (no se recomiendan)
- Cuidado con el manejo de tiempo y fecha
- Información duplicada 

### Maneras interesantes de lidiar con esos problemas
En las funciones que colaboran en la limpieza de datos, encontramos:
- **Extracción de texto**: LEFT, RIGHT y MID/EXTRAE. En LEFT y RIGHT son requeridos 2 argumentos, especificamente la celda de donde comienza a extraer texto y donde finaliza (comienza en el 1). Y en MID, se requiere celda, posición y finalización.
- **Eliminación de espacios**: TRIM/ESPACIOS elimina los espacios innecesarios (no más de 2 entre palabras)
- **Mayúsculas y minúsculas**: LOWER, UPPER y PROPER/NOMPROPIO (capitalización de letras)
- **Concatenación**: CONCAT() (funciona exactamente igual que un `print()`)

### Manejo de tiempo y fechas
El manejo de tiempo y fechas en Excel es fundamental cuando del negocio se habla por una variedad de factores.

**Composición de fechas en Excel**

Las fechas en Excel de manera interna son compuestas por un número entero, llamado **Serial Number**. El número aumenta 1 cada 24 horas.\
Ej: `05/11/23 -> 45057` 

Algo positivo de este enfoque, es que se pueden realizar operaciones entre los números.\
Ej: `05/11/23 - 03/11/23 = 2`

> Dato curioso: Microsoft creyó que 1900 era bisiesto, por lo que las fechas están erróneas técnicamente. Apple, queriendo evitar esta falla informática, comenzó la cuenta de los Serial Numbers desde 1904, por lo que cuando se pase un archivo de Excel de Windows a Mac puede suceder que las fechas estén atrasadas 4 años y un día.

**Funciones de fechas**

- **HOY/TODAY**: Para obtener la fecha actual del sistema.
- **NOW**: Te da la fecha y la hora.
- **DAY, MONTH, YEAR**: Extraen el valor con respecto a una fecha.
- **DATE**: Con ello formás una fecha.
- **NETWORKDAYS**: Cantidad de días hábiles desde una fecha a otra (sin fines de semana). El tercer argumento permite agregar días feriados.
- **DATEDIF**: Diferencia entre fechas. El primer y segundo argumento son las fechas, y el tercero es lo que querés diferenciar, por ejemplo `"y"`.

### Funciones y operadores lógicos
Donde se contienen las típicas bases de todo sistema informático.

**Operadores lógicos**

- Mayor (>) o menor (<) que.
- Distinto de (<>).
- Igual a (=).
- Mayores/menores iguales.

**Funciones** 

- **IF/SI**: Esta función asigna una comparación lógica en el primer argumento, y en los próximos 2 se asignarán los valores a devolver si `True` o `False` respectivamente.
- **IFS**: Permite tener varios comparadores, donde sigue la secuencia de `Comparador -> Resultado; Comparador -> Resultado...`
- **AND/OR**: Estas funciones permiten realizar comparaciones lógicas simultáneas. Devuelven únicamente `True` o `False`.  
- **SUMIF**: Suma basado en un criterio. Primero se coloca el rango del criterio, el criterio en si mismo y finalmente el rango de valores a sumar.\
Ej: `=SUMIF(B2:B24,”seattle”,G2:G24)`.
- **AVERAGEIF, COUNTIF...**: Funcionan exactamente igual a SUMIF().

