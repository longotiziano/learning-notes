# issues_path.md
En este archivo almacenaré todos los diversos problemas a los que me he enfrentado durante la creación de proyectos
que abarcan horas y horas de debugging, con el trayecto del como he enfrentado el problema y cual habría sido la 
solución

## analisis-de-bolsa
* Los diversos problemas en los que me he enfrentado en este proyecto son:
    1. **Problema de logging**: Fue mi primera vez utilizando un sistema de logging serio en Python con la libreria
    nativa de `logging`, por lo que aprendí de la configuración, headers, entre otros. Resulta que no se estaba
    aplicando la configuración en todos los logs nuevos que creaba, ya que si bien estaban bien definidos los
    handlers y el `propagate = True`, no se estaba definiendo de nuevo el `propagate = True` al momento de
    la definición individual del logger y además no definí la raiz de los loggers, lo cual es fundamental si estoy
    usando lo de propagate. Me pude dar cuenta ya que, una vez que utilizaba un logger centralizado con el mismo 
    nombre, andaba todo perfecto. Otro dato es que al momento de loggear con los loggers hijos, no van a heredar
    los handlers, si no que van a enviar los logs al logger padre que si que los tiene y los procesa.
    2. **Obtención de la TNA de LECAPs**: Que manera de sufrir! Sin APIs como la gente y escasez de web scraping
    por culpa de páginas web dinámicas que requieren descargar el HTML desde un navegador externo... Por suerte
    utilicé la imagen oficial de Microsoft que viene con navegadores preinstalados, llegando así a la solución
    con un contenedor externo al flujo principal que corre en distinto cronograma (el scraping con navegadores
    es bastante pesado).
    3. **Directorios que no existen en Docker**: Cuando vos elegís volúmenes en un servicio de Docker lo que
    estás haciendo es mantener una conexión directa entre el host y el contenedor. Si el directorio elegido
    no existe siempre se creará, por lo que puede llegar a haber muchas confusiones si aparecen carpetas que
    no necesitás. Por lo tanto, en caso de que no haga falta tener conexión directa con algún directorio en
    especial es recomendado montarlo en la imagen del contenedor (Dockerfile), de manera tal que se cree
    la carpeta en el contenedor aislada