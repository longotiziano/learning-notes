# HTML and CSS in depth
Primero se toca el HTML.

## GET vs POST

### GET request

Cuando haces una GET request, los datos se envían al servidor a través de la URL (query parameters).

Ejemplo:
`site.com/login?user=admin&pass=123`

Ventajas:
- Puede ser cacheado
- Permite guardar la URL en bookmarks
- Permite compartir enlaces
- Es el método estándar para obtener datos

Desventajas:
- Los datos quedan visibles en la URL
- Límite de tamaño en la URL
- No se recomienda para datos sensibles

### POST request

Cuando haces una POST request, los datos se envían al servidor dentro del body de la petición HTTP.

Ventajas:
- Permite enviar grandes cantidades de datos
- Los datos no aparecen en la URL
- Se usa para crear o enviar información al servidor

Desventajas:
- Normalmente no es cacheado
- No se puede guardar o compartir fácilmente

# Hyper text markup lenguage (HTML)

## Etiquetas semánticas
Escenciales en desarrollo web para ayudar a los motores de búsqueda a entender qué es lo importante y qué tiene cada parte de la página.

Ejemplos:
- `<header>`: Título principal.
- `<nav>`: Navegador en los que suelen estar listas desordenadas. 
- `<main>`: Contenido principal de la página.
- `<article>`: Tienden a estar dentro del `<main>`. Es contenido independiente. En mi portfolio se podría aplicar a cada uno de los proyectos.
- `<section>`: Utilizado para dividir secciones den la página.
- `<footer>`: Se suelen ver contactos, derechos de autor, entre otros.
- `<aside>`: Barras laterales.
- `<figure>`: Imágenes.

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

## Formularios
Los formularios son bloques que contienen todo lo relacionado al ingreso de datos de parte del usuario.
```html
<form action="direccion_de_datos" method="tipo_de_request">
  <input type="text">DAMN
</form>
```
- **Direccion de datos**: A dónde enviar los datos. Puede ser tanto una URL, una ruta absoluta o incluso una relativa. 
- **Tipo de request**: GET o POST

**¿Y para que lo de la dirección de datos?**

Sirve en caso de que quieras enviar la información a tu servidor, o a otra URL.

---

## Inputs
Permiten dejarle al usuario la introducción de datos y valores. 

### Sintaxis y tipos

Podemos encontrar una gran variedad de tipos de inputs, como email, tel, url, date, time, number, range, color, text, password, file, reset...

Su mayoría tiene esta sintaxis:
```html
<input type="input_type" id="para_reconocer_entre_js_css" name="identificar_datos_al_servidor">
```

### Ejemplos de inputs importantes

#### Radio
Los input `radio` permiten las opciones múltiples.
```html
<form method="POST">
    <fieldset id="size"> <!-- Es importante que el fieldset y los input compartan ID -->
        <input type="radio" value="2" name="size"> 2-person table
        <input type="radio" value="4" name="size" checked> 4-person table <!-- checked deja marcado por defecto -->
    </fieldset>
</form>
```

#### Range
Los input `range` permiten elegir un valor en un rango.
```html
<input type="range" min="0" max="100" value="50"> <!-- "value" es el valor por defecto -->
```

#### Select
Únicamente podés elegir entre opciones existentes.
```html
<select>
  <option>Chrome</option>
  <option>Firefox</option>
</select>
```

### Validaciones client-side

En caso de que NO se haga (o se haga incorrectamente) la validación de input, podemos gastar recursos del server, ya que la HTTP request ya se hizo.

**Ejemplos de validaciones**:
```html
<input type="text" id="user" name="user" required> <!-- No puede estar vacío -->
<input type="text" id="user" name="user" required minlength="3" maxlength="3"> <!-- Longitud requerida -->
```

---

## Carga de media en HTML
Formatos comúnmente aceptados:
1. **Video**:
  - `.mp4`
  - `.webm`
  - `.ogg`
```html
<video> <!-- Acá es donde se agregan los parámetros, como por ejemplo loop, controls, muted... -->
  <source src="[ruta_archivo].[formato]" type="video/[formato]">
  <!-- <source src="[ruta_archivo].[formato2]" type="video/[formato2]"> -->
</video>

<!-- Se pueden agregar múltiples sources, para aquellos navegadores
que prefieran un formato antes que otro. -->
```

2. **Audio**:
  - `.mp3`
  - `.wav`
  - `.ogg`
```html
<audio>
  <source src="[ruta_archivo].[formato]" type="audio/[formato]">
  <!-- En caso de .mp3, seria type="audio/mpeg" -->
</audio>
```

3. **Imágenes**
- `.APNG` – Animated Portable Network Graphics 
- `.AVIF` – AV1 Image Format 
- `.GIF` – Graphics Interchange Format 
- `.JPEG` / .JPG – Joint Photographic Expert Group image format 
- `.PNG` – Portable Network Graphics 
- `.SVG` – Scalable Vector Graphics 
- `.WEBP` – Web Picture Format 
```html
<figure> <!-- Elemento semántico -->
   <img src="photo.png" width="320" alt="My Profile Photo"> <!-- "alt" es el atributo alternativo en caso de que la imagen no cargue -->
   <figcaption>A photo of myself on a beach in 2015</figcaption> 
</figure>
```

**PD**: Siempre es recomendado aligerar los archivos de media a través de distintas aplicaciones. Por ejemplo, en imágenes mantener el formato `webp`.

---

## Otros elementos de HTML que pueden ser útiles

### Label
Permite dejar un nombre encima del input.
```html
<label for="id_del_campo"></label> 
```

### Datalist
Da sugerencias o autocompletado a un `<input>`.
```html
<label for="browser">Choose a browser:</label>

<input list="browsers" id="browser" name="browser">

<datalist id="browsers">
  <option value="Chrome">
  <option value="Firefox">
  <option value="Safari">
  <option value="Edge">
</datalist>
```

### iFrames
Incrusta otra página web o contenido externo dentro de tu página. Casos de uso comunes: Google Maps, videos de YouTube y widgets de pago. No recomendado para contenido propio ni navegación interna.

Los iFrames afectan el SEO, la seguridad y el performance. Son herramientas específicas para casos específicos. Los atributos `sandbox` y `allow` son útiles para tener un poco más de resistencia.
```html
<iframe src="https://www.google.com/maps/embed?" width="400" height="300">
  Tu browser no soporta iframes.
</iframe>
```

### Canvas
Elemento utilizado para hacer animaciones en JS.

### Scalable Vector Graphics (SVG)
Son elementos desarrollados por cálculos matemáticos. Muy utilizado para la creación de logos y otras formas, ya que son responsive y no se pixelan ante nada.

---

# Definiciones
- Responsive: Que se adapta al tamaño del dispositivo.
- 2D Canvas: Es una API del browser. Se utiliza con JS.
- WebGL: Es una API del browser. Se utiliza con JS.
- Tabindex: Es un atributo que puede ser utilizado para definir el orden en los que la tecla TAB va a navegar.