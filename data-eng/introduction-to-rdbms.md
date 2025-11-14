# Introduction to RDBMS

## Modelos de Información vs. Modelos de Datos

### Principios Clave para Bases de Datos Eficientes y Escalables
- `Independencia lógica`: modificar la estructura lógica sin afectar aplicaciones.
- `Independencia física`: cambiar aspectos físicos (tipos de datos, almacenamiento) sin afectar el modelo lógico.
- `Independencia del almacenamiento`: usar diferentes motores (MySQL, PostgreSQL, etc.) sin alterar la lógica de la DB.

**Buenas prácticas:**
- Uso de `PRIMARY KEY`.
- Validación de datos.
- `DEFAULT VALUES` (valores predeterminados).
- Uso de `VIEWS` para simplificar y controlar accesos.
- Administración de acceso (usuarios, permisos).

## Entity-Relationship Diagram (ERD)
- **Entidades (rectángulos):** objetos, personas o conceptos.
- **Atributos (óvalos):** características de las entidades.
- **Relaciones (diamantes):** cómo se vinculan las entidades.
- **Crow's Foot Notation (`|<`, `||`, etc.):** define cardinalidad (1:1, 1:N, N:N).

## Conceptos Matemáticos detrás de RDBMS

### 1. Sets (Conjuntos)
- Lista desordenada de elementos únicos.
- Operaciones:
  - Unión, intersección, diferencia.
  - Subset, empty set, power set, universal set.
  - Disjoint (sin elementos en común).
- Visualización: Diagramas de Venn.

### 2. Relaciones
- Basadas en pares ordenados `(a, b)`.
- Tipos:
  - `Reflexiva`: a se relaciona consigo mismo.
  - `Simétrica`: si a → b entonces b → a.
  - `Antisimétrica`: si a ≥ b entonces b no puede ser mayor que a.
  - `Transitiva`: si a → b y b → c entonces a → c.

## Relational Schema (Esquema Relacional)
- Representa la estructura de una tabla: columnas y sus tipos de datos.
- `Degree`: número de columnas.
- `Cardinality`: número de filas.
- `Tuple`: valores individuales de una fila.

## Deployment Topology en RDBMS

### Factores a Considerar
- Escalabilidad.
- Performance.
- Naturaleza de la aplicación.

### Tipos de Deployment
- **Monolítico (Primer Grado):** todo en el mismo servidor (interfaz, lógica, DB).
- **Cliente-Servidor (Segundo Grado):**
  - Client Layer: interfaz de usuario.
  - Server Layer: lógica de negocio, base de datos en otro servidor.

## Capas de una Base de Datos
- `Data Access Layer`: interfaz entre aplicación y base (SQL, ORM, API).
- `Database Engine Layer`: motor que ejecuta SQL (PostgreSQL, MySQL).
- `Data Storage Layer`: almacenamiento físico (discos, SSD).

## Arquitectura de 3 Capas (Tier 3 Architecture)

### Propósitos
- Security: prevenir accesos no autorizados.
- Performance: reducir carga en DB.
- Maintainability: cambios sólo por administradores.

### Capas
1. **Presentation Layer:** interfaz de usuario.
2. **Application Layer:** lógica de negocio y conexión a la DB.
3. **Database Layer:** acceso a los datos.

## Cloud Deployment
- Todo vive en la nube, accesible con internet.
- Escalable y mantenido por el proveedor.  
Ejemplos: AWS RDS, Google Cloud SQL.

## Tipos de Arquitectura de Base de Datos

- **Shared Disk Architecture:** varios servidores acceden al mismo disco.  
  Ej: Oracle RAC.
- **Shared Nothing Architecture:** cada nodo tiene su propio almacenamiento, CPU y memoria → escalabilidad.  
  Ej: Cassandra, Google Bigtable.
- **Híbridas:** combinan ambas.

## Buenas Técnicas de Data Management
- **Database Replication:** copia de la DB en otros servidores (High Availability).
- **Partitioning / Sharding:**
  - Partitioning: dividir una tabla.
  - Sharding: dividir en varias bases.

## Roles Claves en el Manejo de Bases de Datos
1. **Ingenieros de datos y DBAs:** crean, administran y monitorean objetos de la DB.
2. **Data scientists y business analysts:** analizan datos, crean modelos predictivos.
3. **Application developers:** desarrollan apps que consultan la DB (SQL, APIs).

## Object Relational Mapping (ORM)
Traduce objetos de código a tablas de la base.  
Ejemplos: `SQLAlchemy` (Python), `Hibernate` (Java), `Entity Framework` (.NET).

## Functional Dependencies (Dependencias Funcionales)
Cuando un atributo determina otro.  
Ejemplo: `dni → nombre`.

## Candidate Keys (Claves Candidatas)
Columnas o grupos mínimos de columnas que pueden ser `PRIMARY KEY`.

## PostgreSQL
- Combina relacional + orientado a objetos.
- **Características:** herencia de tablas, tipos definidos por usuario, sobrecarga de funciones.

## Tipos de Sentencias SQL
- **DDL (Data Definition Language):** `CREATE`, `ALTER TABLE`, `DROP`, `TRUNCATE`.
- **DML (Data Manipulation Language):** `INSERT`, `UPDATE`, `SELECT`, `DELETE`.

## Razones para Mover Datos
- Poblar DB inicialmente.
- Crear entornos de desarrollo/testing.
- Snapshots de seguridad.
- Exportar datos.
- Agregar datos a tablas existentes.

## Backup y Restore
Permiten replicar la DB conservando su estado original.

## Importar y Exportar Datos
- Se puede usar GUI, APIs o CLI.
- Formatos comunes: `DEL`, `ASC`, `PC/IXF`, `JSON`.

## Database Hierarchy

1. **Instance:** entorno único de servidor.
2. **Database:** almacena datos y objetos.
3. **Schema:** organiza objetos y permisos.
4. **Database Objects:** tablas, índices, vistas, constraints, aliases.

## Indexes
Mejoran la performance de consultas (búsqueda directa).  
Pueden ocupar espacio extra y reducir rendimiento en bases pequeñas.

## Database Normalization

### Objetivos
- Reducir redundancia.
- Mejorar consistencia e integración.

### Formas Normales
- **1NF:** cada celda un solo valor, filas únicas.
- **2NF:** cumple 1NF, separa valores repetitivos en otras tablas.
- **3NF:** cumple 2NF, elimina columnas que no dependen de la `PRIMARY KEY`.

## OLTP vs OLAP

- **OLTP:** transacciones frecuentes, DB normalizada (3NF).
- **OLAP:** consultas complejas, analíticas, predominan lecturas.

## Constraints
- `DEFAULT`: valor por defecto.
- `PRIMARY KEY`: identificador único.
- `UNIQUE`: evita duplicados.
- `NOT NULL`: no permite nulos.

## Buenas Prácticas de Diseño de Base de Datos
- Comprender requerimientos del negocio.
- Normalizar (y denormalizar si es necesario).
- Establecer `FOREIGN KEYS`.
- Indexar correctamente.
- Aplicar `PARTITIONING`.
- Optimizar tipos de datos.
- Planificar crecimiento futuro.
