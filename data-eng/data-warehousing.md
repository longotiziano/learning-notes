# Data Warehousing

## 1. Introducción
Un **repositorio de datos** (data warehouse) es un sistema de almacenamiento optimizado para análisis, minería de datos, machine learning y reportes.  
Soporta OLAP (procesamiento analítico en línea) y consultas complejas sobre grandes volúmenes de datos.

### Evolución histórica
- **Mainframes:** centros de datos centralizados.  
- **Appliances:** hardware + software prearmados, optimizados.  
- **Cloud:** repositorios escalables, más económicos y disponibles globalmente.

### Importancia en la industria
- Centraliza datos de múltiples fuentes → una sola fuente de verdad.  
- Mejora velocidad de consultas y calidad de datos.  
- Facilita decisiones de negocio más rápidas y precisas.

## 2. Tipos de Repositorios
### Appliances
- Hardware + software prearmado.  
- Ejemplo: `IBM Netezza` (soporte ML, despliegue cloud).  

### Cloud
- Escalables, bajo costo, alta disponibilidad.  
- Ejemplos: `Amazon Redshift`, `Snowflake`, `Google BigQuery`.  

### On-Premises / Híbridos
- Servidores propios o combinados con cloud.  
- Ejemplos: `Oracle Exadata`, `Azure Synapse`, `Teradata Vantage`, `IBM Db2 Warehouse`, `Vertica`, `Oracle Autonomous DW`.

### Consideraciones para elegir un repositorio
- Localidad: on-premises, cloud, híbrido.  
- Seguridad y privacidad.  
- Tipos de datos: estructurados, semiestructurados, no estructurados.  
- Escalabilidad y rendimiento.  
- Compatibilidad, migración y gobernanza de datos.  
- Costos de infraestructura, licencias y mantenimiento.  

## 3. Tipos de Repositorios Analíticos
- **Data Marts:** subconjunto temático (ventas, finanzas), optimizado para OLAP.  
- **Data Lakes:** almacenan datos en formato nativo, muy usados en Big Data.  
- **Data Lakehouses:** combinan ventajas de Data Lakes y Data Warehouses.  

## 4. Arquitectura de un Data Warehouse
### Componentes
1. **Fuentes de datos**  
2. **ETL/ELT pipelines**  
3. **Enterprise Data Warehouse (EDW)**  
4. **Data Marts**  
5. **Capa Analítica:** BI, ML, reporting  

### Componentes analíticos
- **Data Cubes:** análisis multidimensional OLAP (`slicing`, `dicing`, `drill-down/up`, `pivoting`, `roll-up`).  
- **Materialized Views:** snapshots precomputados para mejorar performance de consultas pesadas.

## 5. Modelado de Datos
### Facts y Dimensions
- **Fact tables:** métricas o cantidades, con claves a dimensiones.  
- **Dimension tables:** atributos que dan contexto a los hechos.  

### Esquemas
- **Star Schema:** tabla de hechos central con dimensiones conectadas.  
- **Snowflake Schema:** versión normalizada del star schema.  

### Slowly Changing Dimensions (SCD)
Permite manejar cambios históricos de dimensiones:
- `Retain Original Value`  
- `Overwrite Existing Data`  
- `Preserve Historical Data`  
- `Add New Attribute`  
- `Historical Table`  
- `Hybrid Approach`

## 6. Staging Areas.
- Área intermedia entre fuentes de datos y repositorios.  

## 7. Calidad de Datos.
* `Los datos deben ser`:
    - `Preciso`: Sin duplicados ni errores
    - `Completo`: Sin valores nulos importantes
    - `Consistente`: Formatos compatibles entre sistemas
    - `Actual`: Refleja el estado más reciente
* `Procesos`:
    - Detectar
    - Capturar
    - Reportar
    - Investigar
    - Diagnosticar
    - Corregir
    - Automatizar

## 8. Poblando un Data Warehouse.
* `Tipos de cargas`:
    - Carga inicial + incrementales periódicas
    - Full refresh: cambios de esquema o errores críticos
* `Consideraciones`:
    - Fact tables cambian más rápido que dimension tables
    - Validar staging y calidad de datos antes de poblar


