# Linux y Shell Scripting

Linux es un sistema operativo basado en Unix, especialmente en su modelo de scripting y estructura interna. Existen diversas distribuciones (distros), cada una orientada a distintos objetivos y perfiles de usuario.

## Distribuciones de Linux

* **Debian**: estable, soporte amplio, actualizaciones constantes.
* **Ubuntu**: basada en Debian, ideal para principiantes.
* **Arch Linux**: minimalista, altamente personalizable.
* **Fedora**: pensada para entornos empresariales, respaldada por Red Hat.

## Estructura de Linux

1. **Interfaz de usuario**: interacción visual (mouse, teclado).
2. **Aplicaciones**: shells, herramientas y programas.
3. **Sistema operativo**: administra archivos, tiempo, errores.
4. **Kernel**: gestiona memoria, procesos, seguridad, drivers.
5. **Hardware**: dispositivos físicos.

## Editores de texto en Linux

* **Gedit**: GUI.
* **Nano**: simple, en terminal.
* **Vim**: potente, curva de aprendizaje más alta.

## ¿Qué es Shell?

Interfaz que permite interactuar con el sistema a través de comandos. El más común es **Bash**.

## Usos comunes de Bash

* **Información del sistema**: `whoami`, `id`, `uname`, `ps`, `top`, `df`
* **Navegación**: `ls`, `pwd`, `cd`, `mkdir`, `cp`, `mv`, `rm`, `touch`, `chmod`, `find`
* **Texto**: `cat`, `more`, `head`, `tail`, `echo`
* **Compresión**:

```bash
tar -czf archivo.tar.gz carpeta/
zip -r archivo.zip carpeta/
```

* **Red**: `hostname`, `ping`, `ifconfig`, `curl`, `wget`
* **Fecha y hora**: `date`
* **Automatización**: `cron`, `at`

## Flags y Opciones

* `ls -l` → formato detallado
* `rm -r carpeta/` → elimina carpeta de forma recursiva
* `--help` → muestra ayuda del comando

## Visualización de archivos

```bash
cat archivo      # Muestra contenido completo
more archivo     # Pagina el contenido
head -n 10 arch  # Primeras 10 líneas
tail -n 10 arch  # Últimas 10 líneas
```

## Scripts en Shell

* Lista de comandos interpretados por bash u otros intérpretes.
* Más lentos que programas compilados, pero más fáciles de desarrollar.
* Usos:

  * ETLs
  * Backups
  * Administración de sistema

### Shebang

Primera línea de un script, indica qué intérprete usar:

```bash
#!/bin/bash
echo "Hola mundo"
```

Para hacerlo ejecutable:

```bash
chmod +x script.sh
```

## Pipes y Filters

* **Pipes (`|`)**: encadenan comandos.
* **Filters**: procesan datos de entrada, e.g. `cat`, `grep`, `sort`.

## Metacaracteres

* `#` → comentario
* `;` → separador de comandos
* `*` → comodín (varios caracteres)
* `?` → comodín (un carácter)
* `\` → interpreta literal el siguiente carácter
* `" "` → permite variables y expansión
* `' '` → literal, sin interpretación

## Redirección de Entrada/Salida

* `>` → sobrescribe archivo con salida estándar
* `>>` → agrega al final del archivo
* `2>` → redirige errores
* `<` → usa archivo como entrada

## Modos de ejecución

* **Batch Mode**: `com1; com2`
* **Concurrent Mode**: `com1 & com2`
