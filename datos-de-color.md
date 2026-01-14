# Datos de color

En este archivo almaceno conceptos y observaciones técnicas que me resultan interesantes y que considero valiosas para mi formación en ingeniería de datos.

## Desarrollo

* Al realizar **web scraping** en Python, suele recomendarse el uso de la biblioteca **lxml**, ya que está implementada principalmente en C, lo que la hace entre **5 y 10 veces más rápida** que `html.parser`. Además, es más tolerante a HTML malformado y ofrece mayor compatibilidad con herramientas como XPath y XSLT.
* Para la ingesta y procesamiento de **datos no estructurados** a nivel empresarial, se utilizan herramientas ETL/ELT como **Talend** o **Informatica**, especialmente cuando se requiere orquestación, control de calidad y escalabilidad.


## Virtualización de Windows
Unos pequeños datos acerca de mi trayecto en la creación de la VM en Windows para usar Power BI y Excel sin drama.

* **Paquetes instalados**
    - `qemu-kvm` → emulación + aceleración
    - `libvirt-daemon-system` → servicio central
    - `libvirt-clients` → CLI (virsh)
    - `virt-manager` → GUI

## Linux
Conocimientos interesantes acerca de la arquitectura de Linux.

### En Linux TODO es un archivo que tiene:
- Dueño (user)
- Un grupo: Mecanismo de control de acceso. Habilita capacidades concretas (docker, sudo...).
- Un conjunto de permisos

Y se separan de la siguiente manera:
```bash
[user] [group] [others]
```

---

**Supongamos**:
```bash
-rw-rw-r-- 1 longo-tiziano longo-tiziano 50510 Jan 10 16:40 job_seeking.ods
```
- El primer caracter es el tipo de archivo (`d` para directorios, `c` para dispositivos...)\
`-` = Archivo común
- Usuario con permisos de lectura y escritura
- Cualquier usuario dentro del grupo puede leer y escribir
- Y otros usuarios pueden únicamente leer

Luego, tenemos dueño y grupo del archivo, en ese orden. En este caso, ambos son `longo-tiziano`.\
`50510` es el peso en bytes del archivo, y por último tenemos la última modificación.

## Bash

* Al usar `grep -c`, la opción `-c` hace que el comando devuelva **la cantidad de coincidencias**, no las líneas coincidentes.
* Con `egrep` (equivalente a `grep -E`) se pueden utilizar **expresiones regulares extendidas (ERE)**, que habilitan operadores adicionales como:

  * `|` : OR lógico
  * `+` : una o más repeticiones
  * `?` : elemento opcional
  * `()` : agrupación de patrones

## Adicionales
- **I/O**: Input/Output

