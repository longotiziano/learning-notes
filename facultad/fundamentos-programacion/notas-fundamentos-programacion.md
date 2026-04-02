# 1era clase

## Introducción a C
 
### Origen y evolución
- Entre 1969  y 1973, Dennis Ritchie, crea el lenguaje en  los Laboratorios Bell.
- Modifica el Lenguaje B, creado por Ken Thompson en 1969.
- Reescribe el Sistema Operativo Unix.
- En 1978 se publica el libro "El lenguaje de programación C", por Ritchie y Kernighan.
- Amediados de los 80, se crea el Lenguaje C++, que incorpora orientación a objetos.

### Características
- Es compilado: Utiliza un compilador que convierte programas en lenguaje máquina.
- Es imperativo: Control total. Permite al programador especificar el orden y la cantidad de pasos que quieras para alcanzar un objetivo (Python, Java, C), a diferencia de los declarativos que son más parecidos al lenguaje humano, donde uno simplemente especifica lo que quiere (SQL o algún framework).
- De alto nivel: Cercano al lenguaje humano y lejano al lenguaje máquina, aunque se lo puede considerar de nivel intermedio porque se pueden manejar cosas como la memoria que en la mayoría de lenguajes no es posible.
- Fuertemente tipado (depende): Tiene sistema de tipos estático, pero su disciplina de tipos de débil. Te permite hacer operaciones entre diversos tipos que a lo mejor no debería.
- Estructurado: Permite programar en base a las reglas del paradigma de la programación estructurada.

### Curiosidades
- Linux (en su mayor parte), una parte de Windows, Python, PHP, y demás están escritos en C.

### Tipos de datos
- `unsigned long`: Cuando se pone el `unsigned` significa "sin signo". La computadora siempre usa bits para detectar si el número es positivo o negativo, pero cuando se utiliza esta letra siempre se guarda ese bit para poder almacenar números más grandes. La parte de `long` asegura tener al menos 32 bits, lo que asegura que podés multiplicar la máxima capacidad de almacenamiento.

### Operadores
Los operadores son los mismos que en JS. 

En caso de que se este haciendo una división entre 2 enteros, dará un resultado entero. Esto sigue una regla que es que el tipo de dato más complejo, es el más fuerte.

### Definición de constantes
Se definen al princpio del código de una manera similar a las bibliotecas.
```c
# define PI 3.1415
```

### Placeholders en C (especificadores de formato en C)
Los placeholders son iguales que en Python, con la diferencia de que cada placeholder tiene su formato específico.
- `%i` y `%d`: enteros.
- `%f`: floats.
- `%c`: char.
- `s`: string.
```c
int edad = 1
printf("Edad: %i", edad)
```
Estos especificadores de formato tambien permiten hacen combinaciones, como `%8.2f`sería "tamaño total de 8 dígitos, con dos decimales".

### Lectura de datos
Hay varias formas, pero la que vamos a estar usando hoy es `scanf`.
```c
scanf(tipo, &variable)
scanf("%i", &edad) // siguiendo el ejemplo de antes
```