# Introduction to Data Engineering

## Introducción al Mundo de los Datos
- `Ingeniero de datos`: Obtiene, transforma, crea estructura y presenta datos.
- `Analistas de datos`: Analizan datos y generan soluciones o recomendaciones.
- `Científicos de datos`: Diseñan modelos predictivos y de machine learning.
- `Analistas de negocio`: Traducen datos a decisiones comerciales.

### Roles en la Ingeniería de Datos
- `Ingeniero de almacenes de datos`: Pipelines, extracción, transformación (Python), carga.
- `Arquitecto de datos`: Diseña y organiza bases de datos.
- `Data Manager`: Administración global de datos, seguridad, cumplimiento.
- `Administrador de datos`: Seguridad y privacidad de datos.

### Tipos de Datos
- `Estructurados`: Tablas, columnas. Ej: SQL, spreadsheets.
- `Semi-estructurados`: Emails, JSON, XML, binarios.
- `No estructurados`: Web, redes sociales, imágenes, videos.

### Tipos de Archivos
- `CSV`, `TSV`, `XML`, `PDF`, `JSON`.

### Fuentes de Datos
- Bases de datos relacionales.
- Archivos planos.
- `APIs / Web Services`: JSON, XML, CSV.
- `Web Scraping`: BeautifulSoup, Pandas.
- `Data Streams`: Kafka, datos en tiempo real.

### Lenguajes para Ingeniería de Datos
- `SQL`, `Python`, `R`, `Java`, `Bash/Shell`, `PowerShell`.

### Metadatos
- `Technical metadata`: columnas, tipos de datos.
- `Process metadata`: logs, auditorías.
- `Business metadata`: relevancia, contexto.

### Repositorios de Datos
- `Bases de datos relacionales`: PostgreSQL, MySQL, SQL Server.
- `NoSQL`: Redis, MongoDB, Cassandra, Neo4J.
- `Data Warehouses`: BigQuery, Snowflake, Redshift.
- `Data Marts`: subconjuntos de data warehouses.
- `Data Lakes`: almacenamiento barato y flexible.

### Elección de Repositorio
Depende de tipo, volumen, frecuencia de acceso, escalabilidad.

### ETL, ELT y Data Pipelines
- `ETL (Extract – Transform – Load)`:
  - Extract: batch/streaming.
  - Transform: limpieza, validación.
  - Load: inicial, incremental, full refresh.
- `ELT (Extract – Load – Transform)`: transformación luego de cargar.
- `Data Pipelines`: orquestación automática del flujo.  
  Herramientas: Apache Airflow, Apache Beam.

### Data Integration
Selección, conexión, enriquecimiento, filtrado, formateo y carga.

### Big Data
- `5 Vs`: Volumen, Velocidad, Variedad, Veracidad, Valor.
- Herramientas: Hadoop, Spark, bases NoSQL.

### Arquitectura de la Plataforma de Datos
- Extracción batch + streaming.
- Almacenamiento rápido (bases de datos o iPaaS).
- Procesamiento con `Spark` + `HDFS`.
- Orquestación con data pipelines.

### Seguridad en Arquitecturas de Datos
- `Infraestructura física`: autenticación, vigilancia.
- `Seguridad de red`: firewalls, segmentación.
- `Aplicación`: desarrollo seguro.
- `Datos`: cifrado en reposo y tránsito (SSL/TLS).

### Triada CIA
- `Confidencialidad`, `Integridad`, `Disponibilidad`.

### Extracción de Datos
- `SQL`, `CQL`, `APIs`, `Web scraping`, `Sensor data`, `Data Exchanges`.

### Data Wrangling
- Transformación y preparación de datos.
- `Normalización` y `Desnormalización`.
- Manejo de nulos, duplicados, tipos de datos.

### Herramientas de Data Wrangling
- `Spreadsheets`, `OpenRefine`, `Google Dataprep`, `Watson Studio`, `Trifacta Wrangler`.
- `Python`: Jupyter, NumPy, Matplotlib, Pandas, PySpark.
- `R`: dplyr, data.table, jsonlite.

### Evaluación de Performance y Buenas Prácticas
- Métricas: latencia, fallos, uso de recursos.
- Logs, recreación de incidentes.
- Optimización: indexación, normalización.
- Rutinas de mantenimiento: time-based y condition-based.

### Data Governance y Compliance
- `Data Governance`: prácticas de privacidad, seguridad e integridad.
- `Compliance`: GDPR, CCPA, HIPAA, PCI DSS, SOX.
