# Introduction to Big Data with Spark and Hadoop.

## ¿Qué es el Big Data?

### Definición de Big Data.
- `Small Data`: es suficientemente pequeñea para que los humanos puedan leerla,
se puede acumular lentamente, consistente y mayormente localizado en sistemas de almacenamiento
en la empresa o en centros de datos.
- `Big data`: Enormes volúmenes con, semi y sin estructura que son ilegibles únicamente por humanos. 
Están distribuidos en la nube o en servidores específicos para este tipo de datos.
- `Ciclo de vida de Big Data`: Bussiness case - Data Collection - Data Modeling - Data Processing - Data Visualization.
- `Las cuatro versus del Big Data`: Velocidad, Volumen, Variedad y Veracidad. También está la
quinta V, que sería Valor (haciendo referencia al valor que se crea con los datos).

### Procesamiento paralelo, escalamiento y paralelismo de datos.
- `Procesamiento linear`: Al tener un problema, se sigue una serie de instrucciones hasta
llegar al output, donde si nada salió bien, entonces se repite el proceso desde el principio.
- `Procesamiento paralelo`: Al tener un problema, se ejecutan las instrucciones de manera
paralela y permitiendo así localizar y arreglar el error de manera más rápida.
- El procesamiento paralelo tiene beneficios como la velocidad, menor consumo de memoria y hardware, y flexibilidad.
- En el Big Data, debido a las velocidades en las que los datos crecen, se opta más por 
el escalamiento horizontal (aumento de nodos en lugar de ampliamento de uno). Algo a 
tener en cuenta es que todos estos nodos tienen que estar coordinados, por lo que no es 
tan fácil. Este último desafio se arregló usando un nodo central, del cual se comunica con
el resto de sistemas.

### Open Source y Big Data.
- Un software open source es gratis, la fuente del código es pública y puede ser 
libremente modificada por quien guste.
- El modelo open source es empleado en el Big Data ya que tienden a ser proyectos largos
 y complejos, los cuales sirven las necesidades de varias organizaciones. 
- Ejemplos de plataformas open source: The Apache Software Foundation, Linux Foundation 
y Cloud Native Computing Foundation. Aunque, el que más impacto tuvo en el Big Data, es 
Hadoop (por lejos).
- Ejemplos de contribuciones de Hadoop: MapReduce, HDFS (Hadoop Distributed File Storage), YARN (manejo de recursos)...

## Introducción al ecosistema de Hadoop.

### Introducción a Hadoop.
- ¿Qué es Hadoop? Un framework utilizado para procesar enormes data sets.
- Los servidores corren aplicaciones en clusters.
- Un Hadoop Cluster es una colección ded computadoras trabajando juntas para performar tareas.
- Esta optimizado para procesar enormes cantidades de datos, tanto estructurados, semi/no estructurados.
- Hadoop Common: Parte esencial del Apache Hadoop Framework. 
- Partes principales del framework de Hadoop: HDFS, MapReduce y YARN.
- `Hadoop Distributed File System (HDFS)`: Maneja enormes cantidades de datos de manera escalable.
- `MapReduce`: Unidad de procesamiento de Hadoop que se basa en la división de datos en grupos
más pequeños.
- `Yet Another Resource Negotiator (YARN)`: Prepara el sistema para correr los datos tanto en
batch, stream, graph processing (administrador de recursos).
- Algunos de los puntos negativos de Hadoop son la débil disponibilidad, cuando la tarea
no puede ser paralelizada, cuando hay dependencias en los datos, las altas latencias en los 
accesos a datos y el procesamiento de muchos datos pequeños.

### Introducción a MapReduce.
- Programa utilizado en Hadoop para el procesamiento de Big Data. 
- Técnica de procesamiento para computación distribuida.
- Computación distribuida: Sistema con múltiples componentes y máquinas que comunican
acciones en una sola vista para el usuario.
- Consiste en dos tareas principales: Map y Reduce.
- `Map`: Procesa los datos en pares clave-valor.
- `Reduce`: Toma los resultados y los agrupa, creando el output final.
- Utilizamos MapReduce dado al Parallel Computing, la velocidad, la flexibilidad, soporte 
para múltiples lenguajes y una plataforma de análisis y almacenamiento de datos.
- Casos comunes de MapReduce: Redes sociales, recomendaciones, instituciones financieras y publicidad.

### Ecosistema de Hadoop.
- El ecosistema de Hadoop esta conformado por componentes que se apoyan entre ellos.
- En ingestión de datos podemos encontramos `Flume` (tiene una arquitectura simple con un 
modelo extensible que permite la aplicación análitica online.) y `Sqoop` (enfocado a la 
transferencia de datos entre Hadoop y una RDBMS).
- En almacenamiento: `HDFS` (ya lo explicamos), `Cassandra` y `HBase` (base de datos
no-relacional que corre en el tope de HDFS y permite acceso aleatorio).
- Procesamiento y análisis: `Pig` (procesamiento ordenado de datos que opera en el
lado cliente del cluster) y `Hive` (más enfocado a reportes, con su propio lenguaje de
programación declarativo y con operación en el cluster del lado del server).
- Acceso a datos: `Impala` (fácil de utilizar, sin necesidad de saber programación) y 
`Hue` (Hadoop User Experience, utiliza SQL y permite un manejo de datos más extenso).

### Hadoop Distributed File Systeme.
- `HDFS` Divide los archivos en bloques, creando réplicas de los bloques y 
almacenandolas en diferentes máquinas.
- Provee acceso a la streaming data.
- Utiliza una linea de comandos para interactuar con Hadoop.
- `Características clave de HDFS`: Costo-eficiente, enormes cantidades de datos,
replicación, tolerancia a fallos, escalabilidad y portabilidad.
- `Bloques`: Unidades de almacenamiento de HDFS, que una vez llegada la cantidad
máxima de almacenamiento se divide en otro bloque, por lo que en caso de que
el espacio que ocupa un data set no sea divisible por la cantidad asignada de
almacenamiento de los bloques siempre quedará un último bloque "incompleto".
- `Nodos`: Un sistema singular que es responsable de almacenar y procesar datos.
- Hay dos tipos de nodos: El `nodo primario` (NameNode), que regula el acceso a los archivos
para los clientes y se encarga de la administración del nodo secundario, y el
`nodo secundario` (DataNode), en las que se performan operaciones de lectura y escritura 
luego de las instrucciones del nodo primario. 
- Un `rack` es una colección de 40 o 50 nodos de datos que usan el mismo network switch.
- Un `network switch` es un dispositivo físico que dirige el tráfico dentro de un
cluster o centro de datos para mejorar su eficiencia.
- `Rack awareness`: Es una caracterísitca de HDFS que sirve para optimizar dónde se guardan
y leen los bloques de HDFS teniendo en cuenta la topología de red. Hadoop "sabe" en qué
rack se encuentra cada DataNode, gracias a una configuración que le da el NameNode.
- Analogía de network switch y rack awareness: El switch de red sería como el cableado
físico, mientras que el rack awareness el GPS.
- HDFS hace muy buenas replicaciones gracias al rack awareness.
- `Replicación`: Crea una copia de un bloque de datos. Las copias son creadas para backups.
- El factor de replicación es la cantidad de veces que un bloque fue copiado.
- `Operaciones de escritura`: No se puede modificar datos que ya están almacenados en el
bloque, pero si se puede agregar. El NameNode se asegura que el archivo no exista (en 
caso de que si lo haga larga una excepción).
- `Operaciones de lectura`: El cliente envía una solicitud al nodo primario para
encontrar la ubicación de los DataNodes que contienen los bloques. El cliente podrá
leer los archivos más cercanos al DataNode.

### Hive.
- `Hive` es un repositorio de datos de Hadoop que está diseñado para la lectura, 
escritura y manejo de bases de datos de tipo tabular.
- Bases de datos de tipo tabular: De tipo columnas y filas, como una base de datos
relacional, pero con la diferencia de que no necesariamente son transaccionales.
- Utiliza Hive Query Language (HiveQL), que está inspirado en SQL.
- Escalable, rápido y fácil de usar.
- Soporta limpieza y filtrado de datos dependiendo de las necesidades del
usuario
- `Diferencias entre RDBMS y Hive`: Utilización de SQL contra HiveQL; diferencia de
objetivos (bases de datos - repositorio de datos); capacidad de análisis en tiempo real contra una 
mejor capacidad para el análisis de datos estáticos; capacidad máxima en 
terabytes, mientras que en Hive de petabytes; forzado de cumplimiento de esquema contra
flexibilidad; y a diferencia de las tradicionales RDBMS, Hive soporta particionamiento.
- `Estructura de Hive`: Hive Clients, Hive Services y Hive Storage and Computing.
- `Hive Clients`: Drivers JDBC (permite la conexión a Hive desde aplicaciones de Java) y 
ODBC (permite la conexión a Hive desde aplicaciones basadas en ODBC). 
- `Hive Services`: Permite las consultas, recibiendolás y enviandolas al compilador de manera
optimizada y eficiente. Las tareas se ejecutan luego de ser optimizadas, y también acá se
haya el almacenamiento de los metadatos.

### HBase.
- `HBase` es una base de datos no relacional orientada a columnas, como Cassandra. Corre en
HDFS, anda muy bien con data streaming y provee tolerancia a fallos.
- Es escalar de manera linear.
- Provee lectura y escritura consistente.
- No tiene un esquema fijo.
- Es bueno para conectar con API de Java y provee replicación de datos sobre todos los 
clusters.
- Se pueden hacer familias de columnas, como si fuera una combinación de celdas en Excel.
- `Diferencias entre HBase y HDFS`: Uno almacena datos en forma de columnas y filas, mientras
que el otro almacena de manera distribuida en los diferentes nodos de una red; HBase permite
cambios dinámicos, mientras que HDFS tiene una arquitectura rígida; HBase banca escrituras y
lecturas aleatorias, mientras que Hive es una escritura y varias lecturas; y finalmente 
mientras que uno permite almacenamiento y procesamiento de Big Data, el otro se especializa 
en almacenar.
-  `Conceptos de HBase`: HMaster, Region Servers, Region y Zookeeper.
- `HMaster`: Monitorea las instancias de servidores de región, asigna regiones a los 
servidores y administra todos los cambios que se realizan en el esquema.
- `Region Servers`: Recibe y asigna las peticiones a los servidores, es responsable de
administrar las regiones y se comunica de manera directa con el cliente.
- `Region`: La unidad más pequeña que tiene el cluster de HBase. Contiene múltiples stores y
está formado por dos componentes (HFile y Memstore).
- `Zookeeper`: Mantiene conexiones sanas entre los distintos nodos, provee sincronización 
distribuida y traquea los fallos del servidor.

## Apache Spark.

### ¿Por qué usar Apache Spark?
- `¿Qué es Spark?` Un framework open source in-memory para procesamiento de datos distribuidos
y análisis de volumenes masivos de datos iterativos.
- `In-memory`: Manera de procesar datos cargándolos en RAM en lugar de leyendo y escribiendo
en disco constantemente. Trae beneficios como velocidad, posible iteración y eficiencia, pero 
a consta de el uso de bastante RAM.
- Spark está escrito mayormente en Scala y corre en Java virtual machines (JVMs).
- `Parallel versus Distributed Computing`: Mientras que los procesadores de la computación paralela 
acceden a memoria compartida, los procesadores de la computación distribuida tienen normalmente
su propia memoria distribuida.  
- `Beneficios de la computación distribuida`: Escalabilidad y crecimiento modular, y tolerancia a 
fallos y redundancia.
- `Beneficios de Spark`: Procesamiento a gran escala de datos y análisis; provee tanto 
procesamiento paralelo como también distribuido, escalabilidad y tolerancia a fallos; velocidad
gracias a el procesamiento in-memory; permite flexibilidad en programación con Python, Scala, 
y otras APIs de Java.
- `Comparación de Apache Spark con otras herramientas del BigData`: Un problema que por ejemplo
tiene MapReduce es que al momento de procesar los datos en iteraciones, tiene que leer y 
escribir los datos en cada una de ellas, lo que signifca menor velocidad a comparación de Spark,
que lee una única vez al comienzo de la iteración.

### Básicos de la programación funcional
- `Definición de programación funcional`: Estilo de programación de funciones matemáticas,
que siguen un modelo declarativo. En lugar de enfocarse en el "Qué" en lugar del "Cómo".
- ¿Qué son las funciones lambda? Funciones o operadores anónimos que permiten la 
programación funcional.
- Todos los programas de Spark son hereditariamente paralelos.

### Programación paralela utilizando Resilient Distributed Datasets.
- `Resilient Distributed Datasets (RDDs)`: Abstracción de datos primaria de Spark, que
tiene una colección de elementos tolerantes a fallos, está particionado en los nodos del cluster,
es capaz de aceptar operaciones paralelas y es inmutable.
- RDD soporta (sobre tipos de archivos) texto, SequenceFiles, Avro, Parquet, Hadoop input 
formats. En cuanto a los formatos se encuentran Local, Cassandra, HBase, HDFS, Amazon S3, entre 
otros.
- `Programación paralela`: Uso simultáneo de múltiples recursos para resolver un problema
computacional, corriendo instrucciones simultáneas en múltiples procesadores.
- Spark corre una tarea por cada partición del cluster.
- `Resilient en RDDs`: Siempre son recuperables porque son inmutables, pueden persistir datasets
en memoria sobre las operaciones, lo que aumenta su velocidad.

### Pararlelismo de datos & escalamiento horizontal en Apache Spark.
- `La arquitectura de Spark consiste en 3 componentes`: Almacenamiento, interfaz y administración
del cluster.
- `Almacenamiento`: HDFS y otros formatos
- `Interfaz`: APIs como Scala, Java y Python.
- `Administración del cluster`: Computación distribuida, Standalone, Mesos, YARN y Kubernetes.
- `Núcleo de Spark`: Tolerante a fallos, sistema base, maneja memoria, performa procesamientos de
datos a gran escala de manera paralela y distribuida, cronograma tareas...
- `Arquitectura de una aplicación de Spark`: Driver, Cluster Manager y Executors.
- `Driver`: Aquel que crea el SparkContext, define RDDs, DataFrames y operaciones, divide el 
trabajo en tasks y las envía al cluster, y recibe los resultados y los combina si es necesario.
Tiende a correr en la máquina del usuario o un nodo maestro dentro del cluster.
- `Cluster Manager`: Asigna nodos y recursos para ejecutar las tareas enviadas por el Driver
(ej: Standalone, YARN, Mesos o Kubernetes).
- `Executors`: Procesos que corren en los nodos del cluster. Ejecutan las tareas enviadas por
el driver, mantienen las particiones en memoria y devuelven resultados al driver.

### DataFrames & SparkSQL.
- `SparkSQL` es un módulo de Spark para procesar datos estructurados, utilizado para consultar
a través de SQL o una API de DataFrames familiar. Utilizable en Java, Scala, Python y R.
- `Beneficios de SparkSQL`: Incluye optimización en función del costo, almacenamiento en columnas y
generación de código para hacer consultas más rápidas y finalmente provee una abstracción llamada
DataFrames.
- `DataFrames`: Colección de datos organizada en columnas nombradas. Utilizan RDDs para
performar consultas avanzadas.
- `Beneficios de DataFrames`: Altamente escalables, optimizados, soporte para un amplio rango de
formatos de datos e integración con la mayoría de herramientas del Big Data.

## DataFrames & Spark SQL.

### RDDs en programación paralela & Spark.
- Entre las operaciones de RDD podemos encontrar las `transformaciones` (aunque realmente
crea una nueva RDD, ya que estas mismas son inmutables), que se suelen denominar "lazy"
ya que únicamente se ejecuta cuando es evaluada.
- Para evaluar una transformación se utilizan las `acciones`, que devuelven un valor al
programa del driver luego de correrlas. Mientras no se ejecuten las acciones, las 
transformaciones quedarán pendientes.
- `Spark utiliza DAGs con sus operaciones`. Es como si se fueran agrupando en un DAG interno
que luego de una acción es analizado, optimizado, dividido en etapas y enviado a los 
executors.
- Ejemplos de transformaciones: map(), filter(), distinct(), flatmap()...
- Ejemplos de acciones: reduce(<func>), take(<n>), collect(), takeOrdered(<n>, key=<func>)

### DataFrames & Datasets.
- `Datasets`: Son la abstracción de datos más nueva de Spark. Proveen una APIs para acceder
a una colección de datos distribuidos (como las RDDs). Son objetos JVM creados de manera
segura y proveen los beneficios combinados de las RDDs y de Spark SQL.
- `Características de los datasets`: Inmutables, encodificado que convierte objetos JVM a
una representación tabular, tiene parecido a un DataFrame seguro orientado a objetos y
buena compatibilidad con APIs de Scala y Java (lenguajes estáticos).
- `Beneficios de utilizar datasets`: Realizan operaciones más rápido que las RDDs,
ofrecen los beneficios de Spark SQL y DataFrames, optimizan consultas y mejoran el
uso de memoria.

### Catalyst & Tungsten.
- `Catalyst`: Optimizador de consultas de Spark SQL, basado en programación funcional y 
construido en Scala. Es "personalizable" a través de reglas pre-definidas que definen cómo
se va a correr la consulta.
- En las reglas de Catalyst ya mencionadas se utiliza una estructura de árbol, que define
las reglas en el siguiente orden: -> `Análisis` -> `Optimización logística` -> `Planeo físico` -> 
`Generación de código`.
- `Tungsten`: Optimiza la utilización de recursos como el rendimiento de la memoria y el CPU.
- `Características de Tungsten`: Maneja la memoria de manera explícita, organiza los datos en 
memoria de manera secuencial, puede generar bytecod para ciertas operaciones, no genera 
funciones 
virtuales y cuando se hace un cálculo se almacena en el CPU en lugar de en la RAM.

### Utilización real de Spark SQL.
- `Spark SQL` es un módulo para procesamiento de datos estructurados en Java, Scala y Python.
- Para correr las consultas SQL es necesario crear una table view en Spark SQL. Estas tablas
pueden dividirse en 2 tipos: las vistas temporales, que únicamente existen dentro de la session 
de Spark, y las vistas temporales globales, que existen dentro de toda la aplicación de Spark.
- `Fuentes de datos de SparkSQL`: Parquet Files, JSON Datasets y Hive Tables.

## Development & Runtime Environment Options. 

### Apache Spark Architecture
- `Una aplicación Spark tiene dos procesos principales`: El Driver Program y los Executors, que
se encuentran dentro del cluster.
- `Driver Program`: Un proceso singular que crea el trabajo para el cluster.
- `Executors`: Múltiples procesos dentro del cluster que trabajan en paralelo. Contienen
núcleos, los cuales cada uno de ellos les permite hacer 1 tarea
- `Spark Context`: Realiza la comunicación con el Cluster Manager. Está definido en el Driver,
con un Spark Context por aplicación de Spark. Divide también los distintos trabajos que envía
el Driver en tareas que pueden ser ejecutadas en el cluster, las cuales se encuentran en 
subsets de datos llamados `Partitions`, y estas mismas pueden ser ejeecutadas en paralelo.
- `Worker`: Son los nodos del cluster que lanzan Executors para ejecutar tareas. 
- Incrementar el número de Executors y de sus núcleos aumentaría el cluster parallelism.
- `Spark Stages y Shuffles`: Una Stage es un conjunto de tareas que pueden ser completadas en 
paralelo en la partición de datos que no está dividida por el Shuffle. El Shuffle consiste en
el reordenamiento de datos que ocurre en un cluster cuando Spark necesita enviar datos de unas
particiones a otras a través de la red.
- El Shuffle es costoso pero necesario dependiendo la tarea.
- `Driver deployment modes`: Client Mode y Cluster Mode. La diferencia entre ellos es la
manera física en la que corre el Driver, ya que en el Client Mode corre en tu computadora,
pero en el Cluster Mode corre, valgame la redundancia, en el cluster.

### Vista general de las modalidades de Cluster en Apache Spark.
- `Spark Cluster Manager`: Ya hemos hablado de él.
* Tipos de Cluster Managers:
    - `Spark Standalone`: Viene por defecto con Spark, es el mejor para configuraciones de 
    cluster simples. Este mismo tiene dos componentes principales: Los Workers, de los cuales
    ya se ha hablado; y el Master, que conecta y agrega Workers al cluster.
    - `YARN`: Para Hadoop (también se habló bastante de él).
    - `Apache Mesos`: Para uso general. Es escalable para varias instancias de Spark y
    dinámico entre Spark y otros frameworks del Big Data.
    - `Kubernetes`: Sistema open-source para correr aplicaciones en contenedores.
- Spark puede correr en local, lo que es útil en momentos de desarrollo o debugging.

### Como correr una aplicación de Spark.
- `Cluster Manager`: --master -> Indica a qué cluster manager conectarse. Ejemplos: local, yarn, 
mesos, k8s.
- `Clase principal` (Java/Scala): --class <full-class-name> -> Solo necesario si tu aplicación 
es en Java o Scala. No se usa en PySpark.
- `Modo de despliegue`:	--deploy-mode -> Define si el driver corre en el cliente (client) o en 
el cluster (cluster).
- `Recursos de ejecutores`:	--executor-cores y --executor-memory -> Ajusta CPU y memoria que 
usará cada executor. Ejemplo: --executor-cores 4 --executor-memory 8G.
- `Ayuda`: ./bin/spark-submit --help -> Lista todas las opciones disponibles.
- Cuando ejecutás una aplicación en Spark, tanto el driver como los executors necesitan acceso a 
las librerías y dependencias usadas en tu código.

### Utilizando Apache Spark en Cloud.
- `Beneficios de optar por correr Spark en Cloud` a diferencia de quizá un despliegue en una
máquina local o en un centro de datos: La configuración está optimizada y agilizada, y es simple
de escalar.

### Configurando Apache Spark.
* Spark se puede configurar de 3 maneras:
    - `Propiedades`: Donde se puede controlar y ajustar el propósito de la aplicación. Pueden
    ser hechas programáticamente cuando se crea la SparkSession o creando un objeto SparkConf, o
    bien en el archivo de configuración o bien corriendo un código en shell. El orden jerárquico
    de las propiedades tiene el mismo orden que el dictado (programáticamente > archivo > shell 
    o defaults).
    > Ubicación: <spark-defaykts.conf.template>
    > Ejemplo: <spark-submit --conf spark.executor.memory=4g>    
    - `Variables de entorno`: Configuran el entorno en el que corre Spark. Controlan cosas como 
    memoria, cores disponibles en el cluster manager, puertos, rutas, etc.
    > Ubicación: <spark-env.sh.template>
    > Ejemplo:  <export SPARK_WORKER_CORES=4>
    > Ubicación: <log4j.properties.template>
    - `Logging`: Permite controlar el nivel de logs: INFO, WARN, ERROR, DEBUG.
    > Ejemplo: <rootLogger.level = WARN>
- `Configuración estática`: Debe aplicarse únicamente a valores que no cambian en cada ejecución
y deben estar relacionados a la aplicación, no al despliegue.
- `Configuración dinámica`: Caso contrario que la estática

### Corriendo Spark en Kubernetes.
* Beneficios de utilizar Kubernetes para correr Spark:
    - Descubrimiento entre servicios en la network.
    - Balanceo de carga del cluster.
    - Facilidad de manejo de cluster gracias a un aumento o reducción automático de la
    cantidad de contenedores basado en la demanda.
    - Maneja cómo los contenedores acceden a almacenamiento persistente.
- Algo bueno es que puede correr de manera local previa a lanzarlo a producción mediante
un cluster de Kubernetes local, que al tratarse de contenedores sería lo mismo al momento
de tratarse de otras máquinas.
- Permite portabilidad, distribución de recursos, etcétera.

## Monitoreo y optimización

### Interfaz de usuario de Apache Spark.
- Podés conectarte a la UI mediante `http://<driver-node>:4040`.
- Muestra jobs, stages, tareas, almacenamiento de RDDs y DataFrames, configuración de 
propiedades y entorno, logs, información SQL... Es prácticamente igual que Airflow.

### Monitoreando el progreso de la aplicación.
- Una vez ue una de las tasks del DAG interno falla luego de varios intentos, Spark marca
la tarea, stage y job como fallidos y para la aplicación.
- Para poder acceder a la UI incluso cuando Spark completó todos los jobs y el SparkContext
finalizó, debe estar activado el event logging, al cual se conecta mediante el Spark History
Server.
* Para conectarse a lo recién dicho:
    - `spark.eventLog.enabled true`
    - `spark.eventLog.dir <path>`
    - `http://<host-url>:18080`
    - `.sbin/start-history-server.sh`

### Debuggeando la aplicación de Apache Spark.
* Areas de posibilidad: 
    - `Código de usuario`: Sintaxis, serialización, validación de datos,
    entre otros. Se reportan los errores al driver y aparecerán en sus logs.
    - `Configuración`
    - `Dependencias de aplicación`: Tanto archivos, scripts como también librerias.
    Las dependencias tienen que estar disponibles en todos los nodos del cluster.
    - `Recursos de aplicación`: Si los recursos no están disponibles (CPU y RAM insuficientes)
    luego de que Spark intente reiteradas veces que un worker se libere, largará un error de tipo
    TimeOutError.
    - `Comunicación con la network`
- Los logs se encuentran en el directorio work/
> work/<application-id>/<stdout|error>

### Entendiendo los recursos de memoria en Apache Spark.
- La memoria usable en Spark está limitada a la que permite a las aplicaciones hacer su trabajo,
evitando utilizar así toda la disponible en el cluster. Esto evita errores de falta de memoria.
- `Consideraciones para la configuración de memoria`: Para los Executors, considerar 
procesamiento y caching (caching excesivo tiende a errores) y para el Driver, carga de datos 
variables y manejo de resultados.
- `Persistencia de datos // Caching`: Almacenamiento de cálculos intermedios que persisten
en memoria.
```bash
./bin/spark-submit \
--class org.apache.spark.examples.SparkPi \
--master spark://<spark-master-URL>:7077 \ 
--executor-memory 10G \
/path/to/examples.jar \
1000
```
- Lo común suele ser poner toda la memoria disponible menos 1GB.

### Entendiendo los recursos de CPU en Apache Spark.
- Spark asigna núcleos del CPU a los Executors y Driver. 
- El paralelismo está limitado a la cantidad de núcleos disponibles.
- Luego del procesamiento, pueden ser vueltos a usar para otras tareas.
- Al igual que en la memoria RAM, se espera a que se liberen núcleos para poder ejecutar
las determinadas tareas reintentando y probando la disponibilidad de los recursos.
```bash
./bin/spark-submit \
--class org.apache.spark.examples.SparkPi \
--master spark://<spark-master-URL>:7077 \ 
--executor-cores 8 \ 
--total-executor-cores 50 \
/path/to/examples.jar \
1000
```
- Lo común es que te use todos los núcleos del procesador.
