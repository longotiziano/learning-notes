# JavaScript

## Variables
```javascript
let nombre = "tiziano" // variable normal
const edad = 25 // constante imposible de cambiar
var apellido = "longo" // forma vieja que hay que evitar\
```

## Data Types
```javascript
let nada = null // equivalente a None
let indenifido = undefined // existe en JS
let numero = 3.14 // no existen ints o floats, todo es un number (lo mismo aplica para los enteros)
```

## Funciones
```javascript
// 1. Función declarada
function saludar(nombre) {
    return `Hola ${nombre}`   // template literal con backticks
}

// 2. Función expresión
const saludar = function(nombre) {
    return `Hola ${nombre}`
}

// 3. Arrow function (la más usada hoy en día)
const saludar = (nombre) => `Hola ${nombre}`
```
Las arros functions 

## Condicionales
```javascript
// if/else
if (edad >= 18) {
    console.log("mayor")
} else {
    console.log("menor")
}

// for clásico
for (let i = 0; i < 5; i++) {
    console.log(i)
}

// equivalente al for..in de Python para arrays
const frutas = ["manzana", "pera", "uva"]
for (const fruta of frutas) {
    console.log(fruta)
}
```

## Operadores

### Aritméticos
Los mismos que los de Python

### Comparativos

- `&&`: and
- `||`: or
- `!`: not

#### `===` vs `==` 
```javascript
5 == "5"   // true  ⚠️ compara solo valor, convierte tipos
5 === "5"  // false ✅ compara valor Y tipo
```
Siempre hay que usar `===`

### Asignativos
```javascript
let x = 10
x += 5   // x ahora es 15
x -= 3   // x ahora es 12
x *= 2   // x ahora es 24
x++      // x ahora es 25 (incrementa 1, no existe en Python)
x--      // x ahora es 24 (decrementa 1, no existe en Python)
```

## Estructuras de datos

### Arrays
Equivalente a las listas en Python. Son mutables aunque estén definidas en una variable `const`. Un ejemplo de ellos puede ser un `NodeList` (que es una lista de elementos HTML).

#### Métodos
- `iterable.forEach(sec, <indice-opcional> => function)`: Recorre iterables, permitiendo operar sobre elementos y su respectivo índice.

### Objetos
Como los diccionarios en Python.

#### Métodos
- `Object.values(dict_)`: Es como si hicieras un `[valores for _, valores in dict_.items()]` 

## Asincronismo
En JS, cuando un proceso demora una cierta cantidad de tiempo el programa no se queda congelado en esa línea, si no que deja pendiente esa respuesta para cuando esté disponible. Esto permite que no se congele la página web en cada tarea que se realice.

### Callbacks
La forma más vieja.

### Promesas
Cuando tenemos un proceso que se espera que finalice en algún momento, se le denomina **promesa**. Estos son objetos `Promise(function())` que permiten definir los distintos casos de éxito de los procesos. Para esperar los procesos, se coloca `.then`s

### Async/Await
La forma más moderna. Es azúcar sintáctico para el método de `.then`.
- `async`: Le indica a la función que estás creando que dentro de ella va a haber `await`s adentro.
- `await`: Es como el `.then` en el método anterior.
```javascript
const obtenerDatos = async () => {
    console.log("1 - arranco")
    const respuesta = await fetch("https://jsonplaceholder.typicode.com/users/1")  // espera acá
    console.log("2 - llegó la respuesta")
    const datos = await respuesta.json()  // espera acá también
    console.log("3 - convertí a JSON")
    console.log(datos.name)
}

// Imprime en orden:
// 1 - arranco
// 2 - llegó la respuesta
// 3 - convertí a JSON
// Leanne Graham
```

## DOM
Lo que nos interesa en este momento. 

Es la raíz de cualquier documento HTML. Dentro de él, podemos encontrar los llamados **nodos**: estos mismos representan cualquier etiqueta del cuerpo, desde comentarios hasta el body incluso. Los nodos pueden también tener "características" que serían los atributos de las etiquetas.

## Eventos
- `click`: click del mouse. Recibe automáticamente un objeto `event` que permite ver si fue tecla (y que tecla), mouse, y más información. Esto también pasa para todos los eventos.
- `input`: cuando cambia un input.
- `submit`: se entiende.
- `keydown`: se entiende.
- `mouseover`: se entiende.
- `scroll`: se entiende.
- `DOMContentLoaded`: Cuando carga el HTML.

## Métodos

### Clases
```js
elemento.classList.add('active');      // agrega la clase
elemento.classList.remove('active');   // saca la clase
elemento.classList.toggle('active');   // si la tiene la saca, si no la tiene la agrega
elemento.classList.contains('active'); // devuelve true o false
```