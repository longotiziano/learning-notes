# Pipeline Development

## Casos de uso para un pipeline ETL

* Digitalización y conversión a formato estructurado de archivos multimedia (videos, audios, dibujos, fotos).
* Generación de dashboards e informes empresariales.
* Migración de datos desde sistemas OLTP a OLAP (para análisis).
* Procesamiento de datos previos a modelos de machine learning.

> **Nota:** ETL se usa donde el control de calidad y la validación previa a la carga es más importante que la velocidad o volumen.

## Casos de uso para un pipeline ELT

* Procesamiento y análisis de Big Data.
* Análisis en tiempo real o casi real (con tecnologías como Kafka, Spark).
* Integración de grandes volúmenes y variedad de datos (estructurados, semi y no estructurados).
* Ideal para sistemas cloud-based que permiten alta escalabilidad y flexibilidad.
* Casos donde el almacenamiento "crudo" de datos antes de transformarlos es una ventaja (por auditoría o trazabilidad).

> **Importante:** ELT es común en arquitecturas tipo Data Lake.

## Diferencias entre ETL y ELT

## Razones por las que ELT es emergente

* Demoras en procesos ETL tradicionales.
* Necesidad de trazabilidad completa (mantener datos sin transformar).
* El crecimiento del Cloud Computing y del Data Lake.
* Mejora en storage y procesamiento distribuido (como Hadoop/Spark).

## Técnicas de extracción de datos

* **OCR (Optical Character Recognition):** escaneo y reconocimiento de texto en imágenes/documentos.
* **ADC/CCD:** digitalización de audio o señales.
* **Cookies y logs:** comportamiento de usuarios.
* **Web Scraping:** extracción automatizada desde sitios web.
* **APIs:** interfaces para extracción estructurada desde servicios externos.
* **Querying a bases de datos:** extracción directa mediante SQL u otros lenguajes.
* **Edge computing:** captura y preprocesamiento de datos desde dispositivos periféricos.
* **Aparatos biomédicos:** captura de datos clínicos desde sensores o monitores.

## Tipos de transformación de datos

* **Data typing:** casteo entre tipos (string → int, etc.).
* **Data structuring:** reestructuración (ej. JSON ↔ CSV ↔ tablas).
* **Anonymizing / Encrypting:** asegurar la privacidad.
* **Cleaning:** limpieza de duplicados, nulos, errores.
* **Normalizing:** conversión a una escala o unidad común.
* **Filtering / Sorting / Aggregating / Grouping:** transformaciones analíticas básicas.

## Schema on write vs schema on read

* **Schema on write:**

  * Estructura definida antes de cargar los datos.
  * Más eficiente para consultas estructuradas.
  * Asociado al enfoque ETL y data warehouses.
* **Schema on read:**

  * La estructura se define al leer los datos.
  * Ideal para datos no estructurados o variados.
  * Asociado al enfoque ELT y data lakes.

## Tipos de carga de datos

* **Full loading:** carga completa desde cero.
* **Incremental loading:**

  * Solo carga datos nuevos o modificados.
  * **Stream loading:** en tiempo real (push).
  * **Batch loading:** programado (pull).

## Planes de carga de datos

* **Serial/secuencial:**

  * Uno a uno.
  * Más simple, pero más lento.
* **Paralelo (parallel loading):**

  * Múltiples fuentes a la vez.
  * Los datos se dividen en *chunks* para acelerar el proceso.
  * Muy usado en arquitecturas distribuidas.

## Herramientas populares en procesos ETL

* **Talend Open Studio:**

  * Soporta big data, data warehousing y profiling.
  * Incluye colaboración, monitoreo y scheduling.
  * Drag-and-drop GUI para creación de pipelines ETL.
  * Genera código Java automáticamente.
  * Integración con muchos data warehouses.
  * Open-source.
* **AWS Glue:**

  * Servicio ETL que simplifica la preparación de datos.
  * Sugiere esquemas para almacenamiento.
  * Permite crear ETL jobs desde la consola de AWS.
* **IBM InfoSphere DataStage:**

  * Herramienta de integración de datos para diseñar, desarrollar y ejecutar ETL y ELT.
  * Interfaz gráfica drag-and-drop.
  * Procesamiento paralelo y conectividad empresarial.
* **Alteryx:**

  * Plataforma de analítica de datos self-service.
  * Interfaz drag-and-drop sin necesidad de SQL o código.
* **Apache Airflow y Python:**

  * Plataforma de pipelines como código.
  * Open source, escalable y dinámico.
  * Permite autoría, programación y monitoreo de workflows.
* **Pandas (Python):**

  * Excelente para ETL, exploración y prototipado.
  * Basado en *dataframes*.
  * No escala bien para Big Data → considerar alternativas como Vaex, Dask o Spark.
* **Panoply:**

  * Específica para ELT en la nube.
  * Integración sin código, interfaz SQL.
  * Integración directa con herramientas como Tableau y Power BI.
* **Herramientas para data streams:**

  * Apache Storm, SQLstream, Apache Samza, Apache Spark.

## Rendimiento de un pipeline

* Latencia y latencias individuales.
* **Throughput:** cantidad de datos transferidos por unidad de tiempo.
* Alertas, errores y fallos.
* Ratio de utilización de recursos.
* Sistemas de logging y alertas.

## Etapas de un pipeline

* Extracción
* Ingestión
* Transformación
* Carga
* Cronograma
* Monitoreo
* Sincronización (buffers de I/O)
* Mantenimiento y optimización

> **Clave:** Mantener throughput uniforme y evitar cuellos de botella.

## Load Balanced Pipeline

* Busca que todas las etapas trabajen a ritmos similares.
* Etapas lentas se pueden dividir en sub-etapas en paralelo.
* Introduce pipelines dinámicos o no lineales.
* Objetivo: minimizar y gestionar cuellos de botella.

## Batch vs Streaming Data Pipelines

* **Batch:**

  * Precisión crítica.
  * Programados por cron o triggers.
  * Casos: backups, procesamiento de órdenes, predicciones.
* **Streaming:**

  * Ingesta en tiempo real.
  * Baja latencia.
  * Casos: OLTP, trading, detección de fraude.

## Arquitectura Lambda

* Diseñada para Big Data.
* Dos capas: **Batch layer** + **Speed layer**.
* Se unifican en **Serving layer**.
* Ventaja: combina velocidad y precisión.
* Desventaja: complejidad de mantenimiento.

## Características comunes de herramientas de pipelines

* Automatización de tareas repetitivas.
* Interfaces visuales drag-and-drop.
* Soporte para transformaciones complejas.
* Seguridad y cumplimiento de regulaciones (HIPAA, GDPR).

## Ventajas de representar pipelines en DAGs

* Nodos = tareas, aristas = dependencias.
* Evitan loops y repeticiones.
* Permiten visualización clara y control de dependencias.
* Reintentos aislados y paralelización.

## Concepto de Evento

* Representa algo que ocurrió en un sistema en un momento específico.
* Formatos: simples, key-value, key-value con timestamp.
* Transporte de eventos → *event streaming*.

## Event Streaming Platforms (ESP)

* Desacoplan productores y consumidores.
* Mejoran escalabilidad y organización.

### Componentes de una ESP

* **Event Broker:** núcleo del sistema.

  * Ingester, Processor, Consumer layer.
* **Event Storage:** almacena eventos.
* **Analytics & Query Engine:** consultas en tiempo real.

### Ejemplos de ESP

* Apache Kafka
* Amazon Kinesis
* Apache Flink
* IBM Event Streams
* Azure Event Hubs

## Apache Kafka

* **Usos:** actividad de usuarios, métricas, logs, transacciones, IoT.
* **Arquitectura:**

  * Brokers con particiones.
  * Producers publican mensajes.
  * Consumers leen mensajes (mantienen offset).
  * Topics = canales de datos.
* **Kafka Streams:**

  * Librería para procesamiento en tiempo real.
  * Permite *map, filter, aggregate* sobre streams.
  * Genera *topologías de procesamiento*.

> **Clave:**
>
> * **Topic:** bandeja de entrada/salida.
> * **Source Processor:** consume de un topic.
> * **Sink Processor:** publica a un topic.
> * **Map, Filter, Aggregate:** funciones comunes para procesar datos.
