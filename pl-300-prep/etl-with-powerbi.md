# Extract, Transform & Load Data in Power BI

---

## Data sources contra datasets

Mientras que el dataset es el objeto que se obtiene, el data source el de dónde provienen los datos.

---

**Limitaciones de Power BI**

Los archivos de Power BI para subir están capeados a 1GB, 16000 columnas y 1000 fuentes de datos.

> En caso de que, en Power BI, se reemplace la ubicación de un archivo, debe también ser cambiada en `File -> Options and Settings -> Data Source Settings` y allí tendrás todo lo que necesites acerca de fuentes de datos.

---

**Build Permissions en Power BI**

Power BI te permite setear permisos a los diversos usuarios que acceden a los gráficos, desde permiso para acceder y exportar los datos de tu reporte, hasta modificar los atributos del mismo.\
Cuando compartís un gráfico, podés agregar features tanto **promotions** (da un acceso más público) y **certifications** (más reservado, requiere permisos especiales para acceder a él).

---

### Local datasets contra shared datasets

Ventajas y desventajas:
- **Local datasets**: Accesibilidad rápida para quien lo porta pero débil para quienes no, control total para el individuo pero podría llevar a problemas de colaboración.
- **Shared datasets**: Acceso para todos simultáneamente, control sobre acceso de usuarios, la **promotion feature**, gobernanza, escalabilidad. La única desventaja es que requiere una configuración adicional, ya que de lo contrario empeorará el acceso de usuarios.

### Formas de almacenamiento en Power BI
Hay varias formas de almacenamiento, entre las que se destacan:
- **Import mode**: Se utiliza para cargar pequeños archivos de diversas fuentes, y se almacena en memoria, permitiendo así un rápido acceso. Las actualizaciones de datos se realizan de manera manual.
- **Direct query**: Es mejor para datasets más grandes, y cada vez que se acceda al reporte no se cargan los datos en memoria, si no que se consultan directamente a la fuente. Puede detener desventajas si la fuente de datos es lenta, o las consultas son complejas.
- **Dual mode**: Es como un híbrido entre ambos modos. Power BI Service se encarga de revisar cual es la query óptima para el trabajo. Es especialmente útil cuando se combinan fuentes tanto en Import Mode como en Direct Query.\
Una situación óptima para emplearlo sería por ejemplo tener las tablas fácticas en Direct Query y las dimensionales en Import Mode.

Es importante tener en cuenta que no se puede cambiar de Import Mode a Direct Query en medio de un reporte, ya que el Import Mode es irremplazable.

**¿Cómo almacenar los diversos tipos de datos?**

Simple, como ya sabemos, hay 3 tipos de datos, y el almacenamiento óptimo para los semi-estructurados y no estructurados son, por ejemplo, los **blobs** (binary large object) de cualquier cloud.\
Los blobs son ideales para archivos crudos de gran tamaño, fuera de toda idea de estructura.

---

**Conectores en Power BI**

Los conectores son los que permiten la conexión entre distintas fuentes de datos con la aplicación de Power BI.

> Es interesante agregar que se puede cronogramar la tasa de refresco de los datos, yendo a la página `app.powerbi.com` -> luego al reporte que interesa -> hacer clic en `Schedule Refresh`, con un máximo de 8 refreshs por día.

