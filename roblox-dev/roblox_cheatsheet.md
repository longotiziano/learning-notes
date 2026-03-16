# Roblox Services - Cheat Sheet

## Servicios Comunes

### Players
```lua
local Players = game:GetService("Players")

Players.PlayerAdded:Connect(function(jugador) end)    -- jugador entra
Players.PlayerRemoving:Connect(function(jugador) end) -- jugador sale
Players.LocalPlayer                                    -- jugador actual (solo LocalScript)
Players:GetPlayers()                                   -- lista de jugadores
```

### ReplicatedStorage
Es un almacenamiento compartido entre el servidor y el cliente.
```lua
local ReplicatedStorage = game:GetService("ReplicatedStorage")

-- Guardar cosas accesibles por cliente y servidor
-- RemoteEvents, ModuleScripts, objetos compartidos
ReplicatedStorage:WaitForChild("NombreObjeto")
```
Se diferencia de otros almacenamientos como ServerStorage (cosas privadas del servidor) o ServerScriptService (scripts privados del servidor).

### DataStoreService
```lua
local DataStoreService = game:GetService("DataStoreService")

local store = DataStoreService:GetDataStore("NombreTabla")
store:GetAsync(key)                                    -- leer dato
store:SetAsync(key, valor)                             -- guardar dato
store:RemoveAsync(key)                                 -- eliminar dato
store:UpdateAsync(key, function(valorViejo) end)       -- actualizar dato
```

### ServerScriptService
```lua
-- No se accede por codigo, es un contenedor
-- Aca van todos los Scripts del servidor
-- Los clientes no pueden ver su contenido
```

### Workspace
```lua
local Workspace = game:GetService("Workspace")

Workspace:FindFirstChild("NombreObjeto")  -- buscar objeto (nil si no existe)
Workspace:WaitForChild("NombreObjeto")    -- esperar a que exista
Workspace.CurrentCamera                   -- camara actual
```

### RunService
```lua
local RunService = game:GetService("RunService")

RunService.Heartbeat:Connect(function(dt) end)   -- cada frame (servidor)
RunService.RenderStepped:Connect(function() end)  -- cada frame (cliente)
RunService:IsServer()                             -- retorna true si es servidor
RunService:IsClient()                             -- retorna true si es cliente
```

### TweenService
```lua
local TweenService = game:GetService("TweenService")

local info = TweenInfo.new(
    1,                          -- duracion en segundos
    Enum.EasingStyle.Quad,      -- estilo
    Enum.EasingDirection.Out,   -- direccion
    0,                          -- repeticiones
    false,                      -- reversa
    0                           -- delay
)

local tween = TweenService:Create(objeto, info, {propiedad = valor})
tween:Play()
tween:Cancel()
tween:Pause()
```

### HttpService
```lua
local HttpService = game:GetService("HttpService")

HttpService:JSONEncode(tabla)    -- tabla a JSON
HttpService:JSONDecode(json)     -- JSON a tabla
HttpService:GetAsync(url)        -- GET request
HttpService:PostAsync(url, data) -- POST request
```

### UserInputService
```lua
local UserInputService = game:GetService("UserInputService") -- solo LocalScript

UserInputService.InputBegan:Connect(function(input, gameProcessed)
    if input.KeyCode == Enum.KeyCode.E then
        -- jugador presiono E
    end
end)

UserInputService.InputEnded:Connect(function(input) end)  -- solto tecla
UserInputService:IsKeyDown(Enum.KeyCode.W)                -- tecla presionada?
```

### SoundService
```lua
local SoundService = game:GetService("SoundService")

SoundService:PlayLocalSound(sound) -- reproducir sonido local
```

---

## RemoteEvents (comunicacion cliente y servidor)
```lua
-- Servidor a Cliente
remoteEvent:FireClient(jugador)
remoteEvent:FireAllClients()

-- Cliente a Servidor
remoteEvent:FireServer()

-- Escuchar en servidor
remoteEvent.OnServerEvent:Connect(function(jugador) end)

-- Escuchar en cliente
remoteEvent.OnClientEvent:Connect(function() end)
```

## RemoteFunctions (cuando necesitas una respuesta)
```lua
-- Cliente pregunta, servidor responde
remoteFunction.OnServerInvoke = function(jugador)
    return "respuesta"
end

-- En el cliente
local respuesta = remoteFunction:InvokeServer()
```

---

## Patrones Comunes

### Acceder a un objeto de forma segura
```lua
local objeto = Workspace:WaitForChild("NombreObjeto") -- espera a que exista
local objeto = Workspace:FindFirstChild("NombreObjeto") -- retorna nil si no existe
```

### Proteger llamadas al DataStore con pcall
```lua
local ok, resultado = pcall(function()
    return store:GetAsync(key)
end)

if not ok then
    warn("Error:", resultado)
end
```

### Verificar si es servidor o cliente
```lua
local RunService = game:GetService("RunService")

if RunService:IsServer() then
    -- logica del servidor
end

if RunService:IsClient() then
    -- logica del cliente
end
```

### Obtener el personaje del jugador
```lua
local jugador = Players.LocalPlayer
local personaje = jugador.Character or jugador.CharacterAdded:Wait()
local humanoid = personaje:WaitForChild("Humanoid")
```

---

## Donde va cada script

| Que hace | Tipo de script | Donde va |
|---|---|---|
| Mostrar UI | LocalScript | StarterGui |
| Input del jugador | LocalScript | StarterPlayerScripts |
| Guardar datos | Script | ServerScriptService |
| Logica de combate | Script | ServerScriptService |
| Funciones compartidas | ModuleScript | ReplicatedStorage |
| Animaciones | LocalScript | StarterCharacterScripts |
