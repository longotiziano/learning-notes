# Cheat sheet de Lua y Roblox development

---

## Base de datos
Cuando accedemos a una tabla en la base de datos de Roblox y esta NO existe, la crea:
```lua
local store = DataStoreService:GetDataStore("PrimeraVez")
```
En este ejemplo se crea la tabla "PrimeraVez"

## Operadores
- Los dos puntos `..` en Lua son el operador de concatenación, que une dos strings. Equivalente al `+` de Python.
- `~=` significa "distinto de".
---

## Buenas prácticas
Utilizar la función `WaitForChild` para evitar errores. Esta misma pausa el código hasta que un valor especificado exista.
```lua
local evento = game.ReplicatedStorage:WaitForChild("MostrarInterfaz")
```

---

## Remote events/functions
Son los eventos/funciones en los que participa la interacción del cliente con el servidor. Evento seria si NO necesitas una respuesta, y caso contrario en la función.

---

## Tipos de datos
- Enum: Opción de lista fija. Cosas que no están en el workspace. Ejemplo: `Enum.Material.Grass`
- Vector3: Posiciones. Esta en R3 y son literalmente vectores.
- Number, string, bool...

---

## Servicios útiles

### Players

### UserInputService

### RunService

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

## Tipos de dato en Lua

### UDim2
`UDim2` es el tipo de dato que usa Roblox para definir tamaños y posiciones de UI. Tiene 4 valores:
```lua
UDim2.new(ScaleX, OffsetX, ScaleY, OffsetY)
```