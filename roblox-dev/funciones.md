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

### GetPartsInPart()