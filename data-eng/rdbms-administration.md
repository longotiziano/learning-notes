# RDBMS Administration

## Ciclo de Vida de una Base de Datos Relacional

1. **Análisis de requerimientos:** definir propósito, datos, usuarios y productores de datos.
2. **Diseño y planificación:** definir instancias, bases, esquemas y tablas considerando escalabilidad, rendimiento y seguridad.
3. **Implementación:** crear y configurar objetos, autorizar accesos, automatizar backups y gestionar movimientos de datos.
4. **Monitoreo y mantenimiento:** reportes de rendimiento, parches de seguridad y mejoras continuas.

## Seguridad en Bases de Datos

* Protección contra ataques (anti-malware, ciberseguridad).
* Almacenamiento seguro frente a fallos y desastres.
* Control de acceso granular.
* Transferencias cifradas.
* Archivado confiable.
* Cumplimiento normativo (RGPD, HIPAA, etc.).

## Instancias de Bases de Datos

Definen entornos aislados para administrar cada base de forma independiente dentro del mismo servidor.
## Monitoreo y Metadata

* Uso de **system catalogs**, **schemas** y **directories** para rastrear origen, destino y responsables de los datos.
* En PostgreSQL, consultar metadatos mediante queries al system catalog.

## Configuración de Bases de Datos

* Archivos `.conf`, `.cfg` o `.ini` para:

  * Logs.
  * Puerto del servidor.
  * Connection timeout.
* En PostgreSQL: `postgresql.conf`.

## Plan de Almacenamiento

* Definir capacidad inicial y plan de escalabilidad.
* Separar diseño lógico del físico.
* Optimizar rendimiento.

### Tablespaces y Containers

* **Tablespace:** estructura lógica que guarda tablas/índices.
* **Container:** estructura física donde se almacena el tablespace.

**Niveles de acceso a datos:**

* **Hot:** acceso frecuente.
* **Warm:** acceso regular.
* **Cold:** acceso ocasional.

## Tipos de Backup

### 1. Backup Lógico

* Guarda comandos DDL/DML (`CREATE`, `INSERT`).
* Ideal para migraciones y auditorías.
* Más lento de restaurar.

### 2. Backup Físico

* Copia exacta de los archivos de la base.
* Recuperación rápida pero dependiente del mismo motor y versión.

### Estrategias de Backup

* **Full Backup:** copia completa, lenta pero confiable.
* **Point-in-Time Recovery (PITR):** recuperación hasta un momento exacto (usa logs de transacción).
* **Diferenciales:** cambios desde el último full backup.
* **Incrementales:** cambios desde el último backup (full o incremental).
* **Hot Backup:** con la base online.
* **Cold Backup:** con la base offline.

**Buenas prácticas:**

* Verificar consistencia.
* Guardar en múltiples ubicaciones.
* Comprimir y encriptar backups.


## Transaction Logs

* Registran todos los cambios en la base.
* Usados para **roll forward recovery**.
* En PostgreSQL: **Write-Ahead Logs (WAL)** en `pg_wal`.

## Seguridad del Servidor y la Base de Datos

* Control de acceso físico y lógico.
* Actualización del sistema operativo y monitoreo constante.
* Autenticación (IP, usuario/contraseña, certificados).
* Autorización granular siguiendo principio de menor privilegio.
* Auditorías y logging de actividad.
* Prevención de SQL injection (consultas parametrizadas).

## Sistema de Usuarios y Roles

* Acceso interno o externo.
* Roles predefinidos (owner, reader, writer, backup\_operator).
* Roles personalizados según necesidad.

### Autenticación y Autorización en PostgreSQL

```sql
GRANT CONNECT ON DATABASE db TO usuario;
REVOKE CONNECT ON DATABASE db FROM usuario;
GRANT SELECT ON tabla TO usuario;
GRANT CREATE ON DATABASE db TO usuario;
GRANT EXECUTE ON PROCEDURE proc TO usuario;
GRANT USAGE ON SCHEMA esquema TO usuario;
```
Nota: PostgreSQL no implementa `DENY`, a diferencia de SQL Server.

## Auditorías en Bases de Datos

* Detectan malas prácticas y errores de seguridad.
* Aseguran cumplimiento normativo (GDPR, HIPAA).
* En PostgreSQL se puede usar **pgAudit** para registrar eventos como SELECT, INSERT, UPDATE, DELETE.

## Encriptado de Datos

* **Simétrico:** misma clave para encriptar/desencriptar.
* **Asimétrico:** clave pública/privada.
* Usado para datos en tránsito y en reposo.
* BYOK (Bring Your Own Key) para control de claves.

## Monitoreo y Evaluación de Rendimiento

* Establecer **baseline** de performance.
* Medir QPS, uso de CPU/memoria/disco, tiempos de respuesta.
* Monitorear consultas, usuarios, bloqueos y buffer pools.

**Tipos:**

* Reactivo: posterior a fallos.
* Proactivo: prevención de problemas.

**Herramientas:**

* pgAdmin Dashboard.
* pganalyze y otros.

## Optimización de Bases de Datos

### Comandos Clave en PostgreSQL

* **VACUUM:** limpia dead tuples.
* **VACUUM FULL:** recupera más espacio pero bloquea tablas.
* **REINDEX:** reconstruye índices.

### Optimización de Consultas

* Usar `EXPLAIN` / `EXPLAIN ANALYZE` para ver plan de ejecución.
* Revisar índices, rediseñar consultas, dividir en partes.

### Buenas Prácticas

* Usar `SELECT columna` en lugar de `SELECT *`.
* Revisar índices periódicamente.
* Ajustar buffer pools.
* Evitar consultas complejas innecesarias.

## Tipos de Índices

* **Primary Index:** sobre clave primaria.
* **Secondary Index:** sobre otras columnas.
* **Dense Index:** entrada por cada registro.
* **Multilevel Index:** jerárquico (como un árbol).

## Troubleshooting en Bases de Datos

### Preguntas Clave

* ¿Cuáles son los síntomas?
* ¿Dónde/cuándo ocurre?
* ¿Bajo qué condiciones?
* ¿Es reproducible?

### Pasos

1. Verificar servidor.
2. Verificar instancia.
3. Verificar conexión.
4. Verificar cliente.

### Herramientas

* Logs del sistema y del servidor.
* Comando `pg_isready`.

### Logs

* **Error Log, Event Log, Trace Log.**
* Contienen: timestamp, usuario, error, detalles.

## Automatización de Tareas

* **Tareas:** backups, limpieza, carga de datos, monitoreo.
* **Herramientas:** cron jobs, shell scripts.
* **Beneficios:** ahorro de tiempo, menos errores.

## Reportes, Notificaciones y Alertas

* **Reportes:** estado de salud, actividad.
* **Notificaciones:** intentos de login sospechosos.
* **Alertas:** problemas críticos (espacio, errores graves).
