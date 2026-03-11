# HTML and CSS in depth
Primero se toca el HTML.

# Hyper text markup lenguage (HTML)

## Etiquetas semánticas
Escenciales en desarrollo web para ayudar a los motores de búsqueda a entender qué es lo importante y qué tiene cada parte de la página.

Ejemplos:
- `<header>`: Título principal
- `<nav>`: Navegador en los que suelen estar listas desordenadas. 
- `<main>`: Contenido principal de la página
- `<article>`: Tienden a estar dentro del `<main>`. Es contenido independiente. En mi portfolio se podría aplicar a cada uno de los proyectos.
- `<section>`: Utilizado para dividir secciones den la página.
- `<footer>`: Se suelen ver contactos, derechos de autor, entre otros.
- `<aside>`: Barras laterales

Y otros datos curiosos como por ejemplo `<strong>` es mejor que `<b>` a la hora de resaltar palabras.

---

## Open Graph Protocol (OCP)
Las etiquetas **Open Graph** le permiten al navegador saber de que está tratando la página la cual estás por visitar, por ejemplo: cuando te envían un link por Whatsapp y aparece una preview con una imagen, descripción y tipo de dato lo que está ocurriendo es que el navegador está consultando a esas tags internamente.

Ejemplos:
```html
<meta property="og:title" content="Home - Tiziano Longo">
<meta property="og:description" content="Computer engineering student focused on IT Support.">
<meta property="og:image" content="https://tusitio.com/preview.png">
<meta property="og:url" content="https://tusitio.com">
<meta property="og:type" content="website">
<meta property="og:site_name" content="Tiziano Longo Portfolio">
<meta property="og:locale" content="en_US">
```

Las etiquetas Open Graph **NO** afectan al SEO

--- 

## Search Engine Optimization (SEO)
Es el análisis que realiza un sitio web para rankear tu página en la búsqueda. Se dedica a analizar el HTML completo para luego darte una puntuación específica.

---

## Metatags 
La tag `<meta>`. Esta misma puede ser utilizada para describir el tipo y el contenido de la página, donde la más tradicional podría ser `"viewport"`, que le dice cómo debe escalar y dimensionar la página en pantallas pequeñas.

Otros tipos podrían ser:
- `"description"`: Permite describir la página. Mejora CTR
- `"robots"`: Cómo analizar la página.
- `"author"`: Quién hizo la página. Sin impacto.

Por supuesto que se pueden utilizar muchas metatags al mismo tiempo, es lo común.

---

## Otras tags de HTML que pueden ser útiles
1. **Inputs**: email, tel, url, date, time, number, range, color, text, password, file, reset...
```html
<input type="input_type" id="para_reconocer_entre_js_css" name="identificar_datos_al_servidor">
```
También podemos encontrar otros como `<submit>` (con button), `<button>`, `<checkbox>` o `<radio>` que tienen parámetros adicionales.

2. Siguiendo con los inputs, tenemos **label**: permite dejar un nombre encima del input.
```html
<label for="id_del_campo"></label> 
```

---

# Cascade Style Sheets (CSS)

### Responsive Design
Se le denomina **responsive design** al diseño de páginas web que conserva la accesibilidad y la comodidad para múltiples dispositivos de múltiples resoluciones.

CSS lo permite, gracias a que dentro de él se ejecutan las **Media Queries**, que permiten aplican ciertas reglas cuando unas condiciones se cumplen, como la amplitud de la pantalla. 

Otro tema que colabora mucho con este diseño es el `display: flex/grid` de CSS

# Definiciones
- Responsive: Que se adapta al tamaño del dispositivo