# Registros internos de la CPU

| Registro          | Contiene                            | Pregunta que responde                |
| ----------------- | ----------------------------------- | ------------------------------------ |
| **PC**            | Dirección de la próxima instrucción | ¿Dónde está la próxima instrucción?  |
| **RI / IR**       | Instrucción actual                  | ¿Qué tengo que hacer?                |
| **RD**            | Dirección de memoria                | ¿Dónde tengo que acceder?            |
| **RDM**           | Dato transferido                    | ¿Qué dato estoy leyendo/escribiendo? |
| **R0, R1, R2...** | Datos/operandos                     | ¿Con qué datos trabajo?              |

## Adicionales
Operaciones como la suma, resta, AND, OR o XOR no tienen RD ni RDM, ya que estas no salen del procesador, y los registros RD y
RDM justamente hablan de memoria.

En caso de hacer un salto con el código de instrucción B, el PC se verá modificado.

## Diferencias entre tipos de registros
- Dedicados/Internos: los que estamos viendo
- Propósito general: los 16 registros, desde R0 hasta RF