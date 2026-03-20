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

## 2da clase
Hay muchas maneras de iterar en bash, por ejemplo:
- `for`: sabés la cantidad de iteraciones.
- `while`: no sabés la cantidad de iteraciones.
- `until`: lo mismo que while, pero evalúa la condición luego de ejecutar la primera vez.

### Pipelines
Nos permiten utilizar la salida de un comando como entrada de otro con el símbolo `|`.

```bash
cat ejemplo.txt | wc -l
```

### grep y sed
- `grep` es de los mejores operadores. Permite buscar desde una palabra en un único archivo hasta en un conjunto de archivos y devuelve la línea donde lo encontró.\
Además, `grep` acepta **expresiones regulares** con la flag `-E`.

- `sed` permite realizar cambios en los archivos.
```bash
sed -i 's/old-text/new-text/g' input.txt # Todo esto lo expliqué unas líneas arriba
```

### Expresiones regulares
Cadenas que agrupadas sirven para realizar validaciones de texto. La idea de la clase no es que las escribas, si no que puedas leerlas.

Por ejemplo:
```bash
a? # Una vez o ninguna
a* # Las veces que sea (incluso ninguna)
a+ # Al menos una vez
a{5,9} # Entre n y m veces. Puede omitirse uno de los límites.
```

### Operadores de redirección
Lo que hace es redirigir la salida de un script, que sería el **stdout**. En caso de que no exista el archivo, lo crea.
- `>`: Reescribe por completo el archivo.
- `>>`: Agrega el contenido a la última línea.

También existen los operadores de entrada, que se denominan **stdin**.
- `<`

Y por último, los **stderr**, que encapsulan a los errores y excepciones. Es un logger.
- `2>`
- `2>>`