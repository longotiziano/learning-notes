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

### RenderStepped
Se activa cada vez que pasa un frame en el juego. Depende del jugador.