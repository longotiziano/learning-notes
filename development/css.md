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
- `flex-wrap`: Sirve para que los elementos que no llegan en la fila, pasen a la otra. Es como la función de Excel.
- `flex-direction`: Si es columna o fila.
- `align-self`: Sirve para alinear elementos individuales del flexbox.
- `gap`: Espacio entre los elementos.
- `flex`: <grow> <shrink> <basis>. Se aplica a las cajas que están dentro del contenedor.
    1. `flex-grow`: Toma el espacio sobrante del contenedor y lo distribuye entre el resto de cajas según la proporción de los valores. Añade píxeles por encima pero no a la proporción base, por lo que no afectan a los hijos de la caja.
    2. `flex-shrink`: Define como se achican las cajas en caso de que el contenedor sea más chico de lo requerido.
    3. `flex-basis`: Casi lo mismo que `width`. Revisa el tamaño del objeto para luego aplicar las otras medidas `flex-grow` y `flex-shrink`. Para utilizar la misma medida predeterminada de la caja se utiliza `flex-basis: auto`.

### Grid
Parecido a `flex`, este funciona más en dos dimensiones, como si de filas y columnas se tratase.

### Blox

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

---

## Responsive Design
Se le denomina **responsive design** al diseño de páginas web que conserva la accesibilidad y la comodidad para múltiples dispositivos de múltiples resoluciones.

CSS lo permite, gracias a que dentro de él se ejecutan las **Media Queries**, que permiten aplican ciertas reglas cuando unas condiciones se cumplen, como la amplitud de la pantalla. 

Otro tema que colabora mucho con este diseño es el `display: flex/grid` de CSS.
