# Cheat sheet de Lua y Roblox development

## Base de datos
Cuando accedemos a una tabla en la base de datos de Roblox y esta NO existe, la crea:
```lua
local store = DataStoreService:GetDataStore("PrimeraVez")
```
En este ejemplo se crea la tabla "PrimeraVez"

## Operadores
- Los dos puntos `..` en Lua son el operador de concatenación, que une dos strings. Equivalente al `+` de Python.

## Buenas prácticas
Utilizar la función `WaitForChild` para evitar errores. Esta misma pausa el código hasta que un valor especificado exista.
```lua
local evento = game.ReplicatedStorage:WaitForChild("MostrarInterfaz")
```

## Remote events/functions
Son los eventos/funciones en los que participa la interacción del cliente con el servidor. Evento seria si NO necesitas una respuesta, y caso contrario en la función.

## Definición de funciones
En Lua, hay 2 maneras de definir una función:   
- Con nombre: que simplemente la llamás por el nombre
```lua
local function saludar()
    print("damn")
end
```
- O anónimas: que al no tener nombre, sirven para pasarlas como argumentos a otras funciones
```lua
-- La pasas directamente como argumento
pcall(function()
    return store:GetAsync(key)
end)

-- Es lo mismo que hacer esto:
local function obtenerDato()
    return store:GetAsync(key)
end
pcall(obtenerDato)
```

## Ejemplos de funciones y métodos

### ipairs(tabla)
Es una función de Lua para iterar sobre una tabla en orden, es básicamente un for each.

### pcall(function)
Significa **protected call**, y es lo mismo que el `try; except` de Python.

### GetAsync(tabla, clave)
Esta función devuelve un valor en base a una clave en una "tabla". Importantes las comillas, ya que las bases de datos en Roblox son ni más ni menos que JSONs.

### FireClient(jugador)
Método que usa el servidor para enviarle una señal a un cliente específico.

### GetChildren(objeto)
Devuelve todos los hijos de un objeto en forma de array numérico, tiende a combinarse con `ipairs` para ordenarlo numéricamente.

### IsA(variable, "tipo_de_dato")
Devuelve un valor booleano en caso de que el tipo de dato de la variable coincida con el asignado en el segundo parámetro.

### Connect(evento, funcion) - OnClientEvent:Connect(evento, funcion)
- Connect: Conecta un evento a una función, a la cual la consecuencia de ese evento caerá como parámetro en la función
- OnClientEvento:Connect: Lo mismo que Connect(), pero para LocalScripts

---

## Tipos de eventos

### PlayerAdded - CharacterAdded
- PlayerAdded: Se activa cuando un jugador entra al juego.
- CharacterAdded: Se activa cuando el jugador ya tiene su personaje en el juego o este revive.

### PlayerRemoving - CharacterRemoving
- PlayerRemoving: Se activa cuando el jugador sale del juego.
- CharacterRemoving: Se activa cuando el personaje muere o es removido

### ProximityPrompt
Evento que se ejecuta y permite interactuar con partes dentro del juego.

### MouseButtons
Sintáxis:
```lua
MouseButton[Identificador][Accion]
```
- **Identificadores**: 1 significa clic izquierdo, 2 significa derecho
- **Acciones**:
    1. Click: Presionado y soltado de botón completo
    2. Down: Presionado
    3. Up: Soltado

### Activated
Este evento es lo mismo que un MouseButton1Click, pero es compatible con gente de otras plataformas, como consolas o celulares.

---

## Paradigma programático de Lua
El paradigma de la programación de Lua se basa en la programación orientada a eventos.

Algo muy usado en estos paradigmas y en Lua es la función `Connect()`, la cual permite enviar funciones a los diversos eventos que ocurran en el juego.
```lua
Player.PlayerAdded:Connect(function(jugador) end)
```
En este caso, estaríamos enviando una función al evento `PlayerAdded`. Luego de esto la función recibe el parámetro.

### Proceso de la programación orientada a eventos
Es como si estuvieran ocurriendo eventos constantemente, los cuales tienen consecuencias que son tratadas con las diversas funciones y código que el desarrollador esté implementando.

Los eventos NO devuelven objetos, si no generan consecuencias.
```txt
Programación normal      →  vos controlás el flujo
                             A llama a B llama a C

Programación de eventos  →  el mundo exterior dispara cosas
                             y vos reaccionás a ellas
```

### Ejemplo
```txt
Explorer
├── ReplicatedStorage
│   └── MostrarInterfaz  ← RemoteEvent (el cable)
│
├── ServerScriptService
│   └── Script           ← dispara el evento
│
└── StarterGui
    └── ScreenGui
        └── LocalScript  ← escucha el evento
```

---

## Diseño responsive de UIs

### Argumentos de los más comunes parámetros
La estructura de la mayoría de parámetros es:
```txt
{escala, offset}
  ↑        ↑
porcentaje  píxeles fijos
del padre   extras
```
Por ejemplo, Size = x: {escala, offset}, y: {escala, offset} (define tamaño)

Esta estructura aplica también a AnchorPoint, que sería el punto de referencia; Position (posición)

### Evitar el uso de pixeles al diseñar interfaces 
Por ejemplo, al querer centrar una interfaz:
```txt
AnchorPoint 0.5, 0.5  →  el punto de referencia del Frame
                          es su centro, no su esquina

Position 0.5, 0.5     →  lo ubica en el 50% de la pantalla
                          tanto horizontal como vertical
```

## TweenService
Es el motor con el que Roblox desarrolla las animaciones de UI

### tweenInfo
Es la información de la animación, por ejemplo:
```lua
local tweenInfo = TweenInfo.new(0.3, Enum.EasingStyle.Quad, Enum.EasingDirection.Out)
```
- 0.3 → dura 0.3 segundos
- Quad → la curva de movimiento (empieza rápido, termina suave)
- Out → la desaceleración es al final

### TweenService:Create(objeto, tweenInfo, propiedad_destino)
Crea las animaciones. Para ejecutarlas simplemente ejecutamos `tween:Play()`.