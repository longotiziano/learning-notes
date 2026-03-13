# Notas de Introducción al Desarrollo de Software

## 1era clase: Bash

### Comandos

Para encontrar información acerca de un comando en específico podemos usar:
```bash
# De ayuda
date --help # Te dice que opciones le podes meter
man date # Sería un gran manual del comando
info date # No se

# Interacción
echo -n "Ingrese una nota: " # La flag -n es para anular el salto de línea 
read nota # Permite al usuario realizar un input, y guardando la data en $nota

# Procesamiento de archivos
cat archivo.txt # Te permite leer que hay adentro
sed archivo.txt # Modificación de texto
wc archivo.txt # Cuenta la cantidad de palabras que hay

# Comparadores lógicos
-ge # greater or equal
-le # lower or equal
-gt # greater than
-lt # ...

if [ $nota -ge 4 ]; then # El flag -ge significa greater or equal
echo "Tu examen está aprobado"
else
echo "Tu examen no está aprobado"
fi 
```

---

### Tipos de variablers

Nos permiten almacenar información de algún tipo.

**Las variables pueden ser:**
- Variables de ambiente: Globales y presentes en todos los Shells (.env), y suelen estar en mayúsculas
```shell
env # Este comando despliega las variables de entorno
printenv EJEMPLO # Entre los ejemplos encontramos PATH, USER, USERNAME, SHELL... 
```
- Variables locales: Definidas por el usuario.
```shell
VARNAME="value" # Valida únicamente para la la shell actual
export VARNAME="value" # Para exportar a otras sub-shells 
```
>Las variables son key-sensitive.

---

### Definición scripts

Un script una secuencia de comandos y operaciones que el shell puede interpretar y ejecutar. Por eso se dice que no es un lenguaje de programación, si no un interprete.

Estos scripts se identifican con .sh y se ejecutan, en nuestra materia, con `bash script.sh`.

Y finalmente, deberían ejecutarse sin errores, cumpla su función, sea claro y reusable.

---

### Ejercicios
**Primer ejercicio**:

```bash
# Crear directorios y archivo
mkdir Intro
mkdir Intro/Ejercicio1
touch Intro/Ejercicio1/datos_personales.txt

# Listar el contenido y introducir la información
cd Intro/Ejercicio1
cat datos_personales.txt
echo "información" >> datos_personales.txt

# Visualizuar nuevamente la información
cat Intro/Ejercicio1/datos_personales.txt

# Realizar una copia del archivo
cp datos_personales.txt datos_personales_modif.txt

# Modificar el archivo copiado y reemplazar la palabra soltero por casado
sed -i "s/soltero/casado/g" datos_personales_modif.txt 

# El flag -i significa in-place, que modifica el archivo original
# La s significa substitute
# La g significa global (todas las coincidencias de la línea)

# Contar cantidad de letras
wc -m datos_personales.txt # El flag -m da la cantidad de caracteres
```

**Segundo ejercicio**:

```bash

```