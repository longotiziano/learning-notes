### Instrucciones

| Opcode | Nombre | Operación |
| :---: | :--- | :--- |
| **`1`** | LOAD M | `R = Memoria[XY]` |
| **`2`** | LOAD C | `R = XY` |
| **`3`** | STORE | `Memoria[XY] = R` |
| **`4`** | MOVE | `S = R` |
| **`5`** | ADD | `R = S + T` |
| **`6`** | ADD Float | `R = S + T` (Flotante) |
| **`7` / `8` / `9`** | Lógica | `OR` / `AND` / `XOR` entre `S` y `T` |
| **`A`** | ROTATE | Rota bits de `R` a la derecha |
| **`B`** | JUMP | Si `R == R0` salta a `XY` |
| **`C`** | HALT | Detener máquina |

Curiosamente, cuando apretamos `CTRL + C` también corta los procesos de terminal!