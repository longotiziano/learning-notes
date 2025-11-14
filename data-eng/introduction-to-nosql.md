## introduction-to-nosql

# ¿Qué es NoSQL?
- NoSQL hace referencia a “no únicamente SQL”.
- Engloba bases de datos no relacionales que no siguen el modelo relacional tradicional.
- Ofrecen esquema flexible y escalable.
- Son sistemas distribuidos que aumentan la productividad de los desarrolladores.
- DBaaS (Database as a Service) permite soluciones más flexibles para desarrolladores.

# Aplicaciones de NoSQL:
- Red social:
- Perfiles de usuario: bases de datos de documentos
- Feed de usuario: bases de datos orientadas a columnas
- Administración de sesiones: bases de datos key-value
- Conexiones entre usuarios: bases de datos gráficas

# Características de bases de datos NoSQL:
- No son relacionales.
- Clasificación: key-value, document, column, graph.
- Comunes: escalabilidad horizontal, uso de clave global única, facilidad para desarrolladores.
- Muchas tienen versiones open-source y comerciales.

# ¿Por qué usar NoSQL?
- Escalabilidad horizontal y alto rendimiento
- Disponibilidad y bajo costo
- Esquema flexible y soporte para estructuras variadas
- Capacidades especializadas: indexado, consultas avanzadas, replicación, APIs modernas
- Operaciones CRUD eficientes

# Bases de datos Key-Value:
- Cada clave se asocia a un valor.
- Consultas atómicas, no para relaciones complejas.
- Uso recomendado: datos básicos y no interconectados.
- Ejemplos: DynamoDB, Oracle NoSQL, Redis, Aerospike, Riak KV.

# Bases de datos Document-Based:
- Cada información es un documento (JSON/XML).
- Permiten consultar contenido y operaciones atómicas a nivel de documento.
- Casos de uso: registros de eventos, blogs, APIs REST, manejo de JSON.
- Ejemplos: MongoDB, CouchDB, Terrastore.
- Sharding: particionamiento horizontal de datos distribuidos en nodos.

# Bases de datos Column-Based:
- Surgieron de Bigtable (Google).
- Almacenan datos por columnas y grupos de columnas.
- Cada fila tiene clave única; columnas pueden variar por fila.
- Ventajas: eficiencia de espacio, despliegue en clusters, counters, TTL para datos caducados.
- Desventajas: no soportan ACID, inserción/actualización más débil.
- Casos de uso: logging, análisis comportamiento, e-commerce, IoT.
- Ejemplos: Cassandra, HBase, Hypertable.

# Bases de datos Graph:
- Almacenan nodos (entidades) y edges (relaciones).
- Útiles para datos representables como grafos.
- Limitaciones: difícil escalabilidad horizontal, mal sharding.
- Casos de uso: redes sociales, mapas, sistemas de recomendación.
- Ejemplos: Neo4j, OrientDB, Amazon Neptune.

# Modelos de consistencia:
- ACID: atomic, consistency, isolation, durability (bases relacionales)
- BASE: basically available, soft state, eventual consistency (NoSQL)

# Bases de datos distribuidas:
- Colección de múltiples bases interconectadas y fragmentadas.
- Réplica de datos para protección ante fallas.
- Mejoran disponibilidad, rendimiento y escalabilidad.
- Desafío: consistencia entre nodos.
- El desarrollador maneja la consistencia.

# Teorema CAP:
- Un sistema distribuido no puede garantizar simultáneamente:
- Consistencia (C): todos los nodos ven los mismos datos
- Disponibilidad (A): cada solicitud recibe respuesta
- Tolerancia a particiones (P): funciona aunque haya fallas de red

# MongoDB

## MongoDB:
- En lugar de almacenar los datos en registros o de manera más convencional se especifica en el tipo de base de datos NoSQL document-based. 
- Los documentos son arrays asociados como objectos JSON o diccionarios de Python.
- En MongoDB encontramos colecciones, que son los grupos de documentos almacenados.
- En la base de datos de MongoDB se almacenan las distintas colecciones con una gran variedad de data types, incluyendo sub-documentos.
- MongoDB es una buena elección cuando se tienen datos grandes y complejos y se busca flexibilidad, escalamiento, etcétera. (búsqueda, servicios web...)

## Ventajas de MongoDB:
- Flexibilidad del esquema: No hay necesidad de tener una estructura igual en cuanto a "columnas" (entre comillas)
- "Code-First approach": El código define a la base de datos.
- Esquema evolutivo: Podés agregar datos y "columnas" cuando quieras.
- Datos sin estructura.
- Consultado y análitica: Lenguaje llamado MQL (MongoDB Query Language), el cual es muy amplio.
- Alta disponibilidad: Generalmente, en MongoDB se tiene un replicamiento de datos en trés nodos, 
donde uno de ellos es el primario, y el resto es secundario. 
- Lo anterior mencionado también colabora al mantenimiento, ya que si un nodo se cae, otro lo reemplazará.

## Casos de uso de MongoDB:
- Para una única vista.
- Internet of Things (IoT): Debido a la masiva cantidad de datos que genera.
- E-commerce: Puede almacenar productos con una variada cantidad de atributos, y MongoDB lo soportará de manera excelente.
- Analítica en tiempo real: Respuestas rápidas, debido a su sistema simplificado ETL.
- Gaming: Debido a que se requiere la velocidad.
- Finanzas: Por la velocidad, seguridad (encriptado mientras están en transmisión).

## Operaciones CRUD:
- MongoDB tiene una shell llamada "Mongo shell" para comunicarse con la base de datos.
- use <db>: Te lleva a la base de datos y la crea si no existe.
- show collections: Lista las colecciones.
- <db>.<collection>.insertOne(<document>): Para crear documentos. El outut será el estado de la operación y 
su ID (la crea automáticamente).
- Con el mismo procedimiento de recién pero cambiando .insertMany podemos agregar una lista de documentos.
- <db>.<collection>.findOne/find/count(<condition>): Devuelve registros para lectura.
- <db>.<collection>.replaceOne/updateOne/updateMany(<condition>): Reemplaza y devuelve el estado de la operación 
y la cantidad de afectados.
- <db>.<collection>.deleteMany/deleteOne(<condition>).
- Operador "$set": Es un operador de actualización que asigna o reemplaza un valor de uno o varios campos según una 
condición asignada. Si el campo no existe, MongoDB lo crea automáticamente, y si existe, lo reemplaza.

## Índices:
- Creados para las consultas más frecuentes.
- <db>.<collection>.createIndex({ <column>: 1/-1 }, { unique: True/False }): Crea el indice. En la parte de columnas se pueden crear índices compuestos agregando más columnas.

Framework/pipeline de agregado:
- Serie de operaciones que se aplican sobre los datos para conseguir un resultado específico.
- <db>.<collection>.aggregate([ { $operation: { document } }, ...]): Almacena procesos y crea una especie de pipeline simple de operaciones.
- Ejemplos de etapas: $project, $sort, $count, $merge, $limit, $group, $sum, $min, $max, $avg...

## Replicación y sharding:
- Como se dijo antes, los Replica Set de MongoDB tienden a dividirse en tres nodos.
- Esto último mencionado colabora a reducir la redundancia, mejorar la disponibilidad, 
la tolerancia a fallos (aunque la replicación no es lo mismo que recuperación de desastres!).
- Proceso de replicación: La aplicación actualiza el nodo primario, para que luego los 
nodos secundarios también se actualicen.
- Sharding: Escalamiento horizontal, que se refiere al particionamiento de almacenamiento 
en distintos almacenes de datos, lo que aumenta la capacidad de manera fácil y más barata.
- El nodo primario es el único que puede recibir operaciones.

## Acceso a MongoDB desde Python:
 pip install pymongo
 from pymongo import MongoClient
 uri = "mongodb://USER:PASS@uri/test"
 client = MongoClient(uri)
 campusDB = client.campusManagmentDB
 students = campusDB.students
- Las funciones son las mismas que en pongo MongoDB, pero en lugar de utilizar camel 
case se usa un snake_case.

# Apache Cassandra

## Introducción a Apache Cassandra
- Es una base de datos open source, distribuida, desencentralizada, escalable elásticamente, 
altamente disponible, tolerante a fallos, personalizable y consistente.
- Su diseño se basa en Amazon Dynamo y su modelo en Google Big Table.
- Sus casos de uso tienden a priorizar la disponibiidad constante, la escritura de datos rápida, la escalabilidad...
- La arquitectura de Apache Cassandra es peer-to-peer.
- Es una de las 10 mejores opciones en lo que elección de bases de datos respecta.
- No es reemplazo de RDBMS, no soporta `JOINS`, no está pensado para hacer consultas complejas, 
está cantidad de transacciones. En estos casos se puede combinar con Spark.
- Cassandra es una buena opción cuando la escritura excede a la lectura y este tipo de 
escritura es de agregado (no de actualización o eliminación), cuando las consultas (simples) son 
predefinidas y el acceso a tus datos en via una PK.    
- Casos comunes de uso para Cassandra pueden ser servicios online, para traquear la actividad 
de los usuarios; sitios web eCommerce, para almacenar transacciones y demás datos de interés... 

## Características clave de Apache Cassandra
- Es distribuida (corre en múltiples máquinas) y descentralizada. Todos los nodos son idénticos.
- La distribución de datos inicia con una consulta y una `Partition Key`, que determina el 
nodo en el cual es almacenada la partición y a dónde van los datos. Cabe recalcar que si el 
nodo que recibió la petición (como una consulta) no tiene los datos, se le consultará a aquel 
que si los tenga.
- Arquitectura `peer-to-peer` -> Todos sus nodos son idénticos en cuanto lo que jerarquía 
respecta, por eso se dice que es descentralizada. 
- Réplica: Copia de tus datos, para tolerancia a fallos - La tolerancia a fallos se logra ya que en caso de que un nodo no funcione, se le retornará un error al usuario
informándole.
- Cassandra al momento de la repartición de datos tiene en cuenta Racks (agrupaciones físicas de nodos en un mismo servidor o gabinete)
y Datacenters (grupos lógicos de nodos en distintas ubicaciones).
- Tiene constante disponibilidad, que a veces es sacrificada por consistencia (administrable por los distintos
ingenieros entre altamente disponible con eventual consistencia, o fuerte consistencia con indisponibilidades momentáneas).
- Escalar horizontalmente mediante la adición de nodos en el cluster (conjunto de nodos). La adición y eliminación de nodos no afecta al funcionamiento constante de Cassandra.
- No hay lectura antes de la escritura (predeterminadamente).
- CQL (Cassandra Query Language). Es practicamente lo mismo que SQL.  

## Modelado de datos en Apache Cassandra I.
- Almacena los datos en tablas, las cuales son almacenadas en KeySpaces.
- Los centros de datos y la replicación se definen en el nivel del KeySpace.
- Se recomienda un KeySpace por aplicación.
- En las tablas se contienen filas de columnas. Las tablas pueden ser creadas, dropeadas, eliminadas, 
modificadas... sin afectar el funcionamiento constante de las actualizaciones y consultas.
- Para crear una tabla se deebe definir una PK y otras columnas de datos. Una vez definida la 
PK esta no puede ser cambiada. Las PKs tienen dos componentes -> (Partition Key, Clustering Key/s (esta última es opcional)). 
- La Clustering Key define el orden de final dentro de la partición que asigna la Partition Key.
- En caso de que la PK tenga ambos componentes (como se comentó recién) se la define como una 
tabla dinámica. Caso contrario será una estática. 

## Modelado de datos en Apache Cassandra II.
- Construir una PK que optimice el tiempo de execución de la consulta. 
- Elegir una Partition Key que comience a responder tu consulta y distribuya los datos de manera uniforme en el cluster. 
- Intentar reducir la cantidad de particiones para responder la consulta. 
- Intentar reducir aún más la lectura de datos con una Clustering Key.  

## Introducción a CQL Shell (cqlsh)
- CQL es bastante intuitivo y similar a SQL.

## Tipos de datos de CQL.
- Built-in: Ya están construidos de forma predeterminada en Cassandra (Boolean, Blob (te guarda algo de 1MB, como una imagen), Bigint...)
- Collection types: Asigna a los datos especificados una categoría o grupo. Estas colecciones no deberian de tener enormes cantidades
de datos. Los tipos de collection types son: `Lists`, que es utilizado cuando el orden
de los elementos tiene que ser mantenido; `Maps`, que son pares {Clave : Valor}; y 
`Sets`, que son utilizados cuando los elementos son únicos y no requieren de un orden específico.
- - ALTER TABLE <tabla> ADD <nombre> <coll_type><data_type>;
- User-defined data types (UDTs): Definidos por el usuario. Pueden relacionar múltiples campos de
datos a una columna.

## Operaciones de Keyspaces.
- Debe ser creado antes de crear tablas. Los keyspaces pueden tener un montón de tablas, pero
las tablas pueden tener únicamente un keyspace.
- CREATE KEYSPACE <nombre> WITH REPLICATION = { <replic_strat>, <replic_factor> };
- Para mayor entendimiento, <replic_strat> hace referencia a qué nodos van a almacenar las
réplicas, mientras que <replic_factor> señala la cantidad de réplicas que se van a
poner en los diferenets nodos en el cluster.
- El flujo de datos en tablas es:
- - Partition Key → (define la partición/nodo) → Clustering Key → (define el orden de las filas en esa partición).
- A nivel distribución en el cluster:
- - Replication Factor (RF) → cuántas copias de cada partición existen en el cluster.
- - Replication Strategy → cómo se reparten esas copias (por datacenter, racks, etc.).
- `Lo que te queda`: Partition Key -> Replication Factor & Strategy -> Clustering Key.
- En Cassandra no hay mejores réplicas que otras.

## Operaciones en tablas.
- Tabla: Entidad lógica que guarda datos en un nivel de cluster y nodo. La creación es
prácticamente igual que en SQL. Se tiende a crear la PK al final del todo (si es compuesta).
- La especificación `STATIC` en significa que una columna tendrá un único valor por partición.
- CREATE TABLE <creacion_de_tabla> WITH CLUSTERING ORDER BY (<columna>); -> Esto define el orden
de almacenamiento de las filas dentro de la partición. 

## Operaciones CRUD.
- Con respecto a las operaciones en general, cuando el nodo recibe una consulta, ese mismo se convierte en el coordinador
para esa operación.
- En cada operación, los datos van primero a la memoria (Memtable, es para mayor velocidad de escritura),
luego se registran en el Commit Log (tolerancia a fallos) y finalmente se guardan en disco en forma de
SSTables.
- Los INSERTs siempre requieren una PK completa.
- 
