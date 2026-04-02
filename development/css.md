# Cascade Style Sheets (CSS)
El maldito CSS.

---

## Displays
Los displays permiten organizar la página web en secciones de una manera más eficiente y rápida. 

Cada uno de los displays se llaman con la propiedad `display`.

### Flexbox
Este tipo de display sirve cuando se trata de organizar elementos de una dimensión.

#### Propiedades
- `justify-content`: Controla la distribución y alineamiento de las cajas, como con `center`, `flex-start`... No sirve en caso de que el item tenga algún `flex-grow`.
    1. `space-evenly`: Misma distancia entre elementos y bordes
- `flex-wrap`: Sirve para que los elementos que no llegan en la fila, pasen a la otra. Es como la función de Excel.
- `flex-direction`: Si es columna o fila.
- `align-items`: Sirve para determinar en que parte del flexbox van a estar ubicados los items: principio, centro, final...
    1. `stretch`: Setea la misma altura con respecto al contenedor flex.
- `align-self`: Sirve para alinear elementos individuales del flexbox.
- `gap`: Espacio entre los elementos.
- `flex`: <grow> <shrink> <basis>. Se aplica a las cajas que están dentro del contenedor.
    1. `flex-grow`: Toma el espacio sobrante del contenedor y lo distribuye entre el resto de cajas según la proporción de los valores. Añade píxeles por encima pero no a la proporción base, por lo que no afectan a los hijos de la caja.
    2. `flex-shrink`: Define como se achican las cajas en caso de que el contenedor sea más chico de lo requerido.
    3. `flex-basis`: Casi lo mismo que `width`. Revisa el tamaño del objeto para luego aplicar las otras medidas `flex-grow` y `flex-shrink`. Para utilizar la misma medida predeterminada de la caja se utiliza `flex-basis: auto`.
- `flex-flow`: <flex-direction> <flex-wrap>

### Grid
Parecido a `flex`, este funciona más en dos dimensiones, como si de filas y columnas se tratase. Los espacios que estan entre cada celda se llaman **gutters** o **gaps**.

Se tiende a dividir en 12 columnas, ya que a partir de ahí podés especificar mejor las dimensiones de cada uno de los bloques dentro del grid. Ya que 12 es un número con varios divisores (1, 2, 4, 6 y 12). 
```css
.container {
  display: grid;
  grid-template-columns: repeat(12, 1fr);
}

.elemento {
  grid-column: span 6; /* ocupa la mitad */
}

.sidebar {
  grid-column: span 4; /* ocupa un tercio */
}
```

#### Propiedades
- `grid-template-<columns/rows>`: Define cuántas columnas tiene el grid y el tamaño de cada una de ellas. Las unidades de medida pueden ser cualquiera pero son muy utilizadas `fr` (espacio disponible) y `auto` (se acomoda el valor).
    1. `repeat(cant_cols, tamaño)`: Para que todas tengas la misma.
    2. `tamañox, tamañoy, tamañoz`: Para definr individualmente.
- `grid-<column/row>`: Define que tanto se va a expandir una columna.
    1. Ej: `grid-column: 1 / 3` -> "Desde la primera hasta la tercera".
    2. Ej: `grid-row: 1 / span 3` -> "Desde la primera y cubrí 3 columnas más".
    2. Ej: `grid-row: 1 / -1` -> "Desde la primera y a la última".
- `place-items`: Manera abreviada de utilizar `align-items` y `justify-content` (en ese orden ).
- `grid-template-areas`: Este de acá es bastante útil. Te permite definir las áreas de un grid con nombres. De manera tal de que, cuando estemos haciendo las clases de los elementos que estarán contenidos dentro del grid, podríamos hacer `grid-area: <nombre_del_area>`. Por ejemplo:
```css
.container {
  display: grid;
  grid-template-columns: 1fr 3fr 1fr;
  grid-template-rows: auto 1fr auto;
  grid-template-areas:
    "header  header  header" /* Se ponen 3 headers por que querríamos que ocupe 3 columnas */
    "sidebar content aside"
    "footer  footer  footer";
}
```

### Blox

---

## Posiciones
En los navegadores, cada elemento sabe dónde se posiciona y cuánto espacio ocupa. Cuando se agrega otro elemento adelante, este respeta lo dicho y se acomoda en función a las reglas y dimensiones disponibles y por ocupar. 

Estas cuestiones pueden ser modificadas con la ayuda de diversas propiedades.

### position
Sirve para determinar la posición de un elemento y en relación a qué. Te permiten modificar la cantidad de distancia a la que se mueven respecto a otros elementos.
- `static`: Es el predeterminado, no te permite modificar la distancia dicha.
- `relative`: Sigue el flujo antes dicho, pero te permite modificar la distancia.
- `absolute`: Este es interesante, ya que sale del flujo común y permite posicionarse en base al ancestro más cercano con `position: relative`. Si no hay ninguno, se posiciona relativo a la página entera.
- `fixed`: Sale del flujo y se posiciona relativo a la ventana del browser. No se mueve aunque hagas scroll.
- `sticky`: Parecido al anterior, solamente que podés medir la distancia del "sticky" en el elemento.

---

## Unidades de medida
Las unidades de medida que decidí incorporar en el cheat sheet son las relativas, que dependen del tamaño de pantalla del usuario.
- `em`: Font size of the parent where present.
- `ex`: x-co-ordinate or height of the font element.
- `ch`: Width of the font character.
- `rem`: Font size of the root element.
- `lh`: Value computed for line height of parent element. 
- `rlh`: Value computed for line height of root element which is <html>. 
- `vw`: 1% of the viewport width.
- `vh`: 1% of the viewport height.
- `vmin`: 1% of the smaller dimension of viewport.
- `vmax`: 1% of the larger dimension of viewport.
- `%`: Denotes a percentage value in relation to its parent element. 

**Las más usadas podrían ser:**
- `rem`: Para cosas relacionadas al texto, padding o margin, ya que escala con la tipografía del usuario.
- `%`: Cuando querés que el tamaño sea relativo al padre.
- `vh/vw`: Típico para cosas más grandes, aquellas que son la base de la página.
- `px`: Bordes, sombras, gaps o cosas que no querés que escalen, como el nav.
- `fr`: Unidad de medida propia de los grid. Significa "fracción de espacio disponible".

---

## Selectores
Son las maneras de referirse a los diferentes elementos creados en HTML desde CSS.

### de Etiquetas
Como por ejemplo, cuando utilizamos `nav {...}` lo que estaríamos haciendo es referirnos a todas las etiquetas `nav` a través de un selector de etiquetas.

### de Clases
En cambio, cuando hacemos `.[nombre-de-la-clase]` estaríamos usando el selector de clases.

### de Atributos
Algo curioso de los selectores es que podemos referirnos a clases y etiquetas que tengan atributos específicos, por ejemplo `a[href="lol.com"]`.  

### nth-of-type & nth-child
Permiten agarrar al enésimo elemento, como por ejemplo:
- `li:nth-of-type(2) {...}`: Afecta al segundo elemento de `<li>`
- `li:nth-child(2) {...}`: Afecta al segundo hijo de `<ul>`

### de Estrella (*)
Seleccionan TODO.

### Agrupados
Si pones comas podes ahorrarte escribir lo mismo muchas veces

### Orden de selectores
En HTML > IDs > Clases, atributos y pseudo clases > Elementos

---

## Responsive Design
Se le denomina **responsive design** al diseño de páginas web que conserva la accesibilidad y la comodidad para múltiples dispositivos de múltiples resoluciones.

CSS lo permite, gracias a que dentro de él se ejecutan las **Media Queries**, que permiten aplican ciertas reglas cuando unas condiciones se cumplen, como la amplitud de la pantalla. 

Otro tema que colabora mucho con este diseño es el `display: flex/grid` de CSS.

---

## Pseudoelementos y pseudoclases

### Pseudoelementos
Sirven para colocar decoraciones sin manchar el HTML. Se llaman pseudoelementos porque parecieran no existir, pero son elementos del HTML.
```css
.nav-button::after {
  content: ''; /* importante, para que pueda existir */
  position: absolute; /* para poder jugar con la posición del objeto */
  /* resto-de-propiedades-que-crean-el-objeto: ...; */
}
```
`after` coloca el elemento luego del tag, y `before` antes. Es indiferente usar esta particularidad en caso de que estemos usando el atributo `position`.

### Pseudoclases
Son detecciones de estado, no eventos.

#### Estados de interacción
```css
:hover   /* mouse encima */
:active  /* mientras clickeás */
:focus   /* elemento seleccionado (tab o click en inputs) */
:visited /* link ya visitado */
```

#### Estados de formularios
```css
:checked  /* checkbox o radio seleccionado */
:disabled /* elemento deshabilitado */
:enabled  /* elemento habilitado */
:required /* input con required */
:valid    /* input con valor válido */
:invalid  /* input con valor inválido */
```

#### Posición entre hermanos
```css
:first-child  /* primer hijo */
:last-child   /* último hijo */
:nth-child(2) /* segundo hijo */
:nth-child(odd)  /* hijos impares */
:nth-child(even) /* hijos pares */
```
#### Negación
```css
:not(.activo) /* todos los elementos que NO tengan la clase activo */
```