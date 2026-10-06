# Registros internos de la CPU

Los **registros internos** son pequeñas unidades de almacenamiento dentro de la CPU. Se utilizan para guardar temporalmente **instrucciones, direcciones y datos** mientras el procesador ejecuta un programa.

## 1. Contador de Programa (PC)

El **Contador de Programa (PC, Program Counter)** almacena la **dirección de memoria de la próxima instrucción que debe buscar la CPU**.

Por ejemplo, si tenemos:

```text
Dirección    Contenido
20           11AA
22           2080
24           8210
```

Si:

```text
PC = 20
```

la CPU sabe que debe buscar la instrucción ubicada en la dirección `20`:

```text
M[20] = 11AA
```

Una vez que la CPU obtiene esa instrucción, el PC normalmente se actualiza para apuntar a la siguiente instrucción:

```text
PC = 22
```

Entonces:

```text
PC = 20
 ↓
M[20] = 11AA
 ↓
RI = 11AA
 ↓
PC = 22
```

> **Importante:** el PC contiene una **dirección**, no la instrucción en sí.

**En resumen:**

> El PC indica **dónde está la próxima instrucción que debe ejecutarse**.

---

## 2. Registro de Instrucción (RI / IR)

El **Registro de Instrucción** almacena la **instrucción que la CPU está ejecutando actualmente**.

Por ejemplo, si la instrucción ubicada en memoria es:

```text
11AA
```

la CPU la carga en el Registro de Instrucción:

```text
RI = 11AA
```

A partir de ahí, la Unidad de Control interpreta la instrucción y determina qué operación debe realizar.

**En resumen:**

> El RI indica **qué tiene que hacer la CPU**.

---

## 3. Registro de Dirección de Memoria (MAR)

El **Registro de Dirección de Memoria** almacena la **dirección de memoria a la que la CPU quiere acceder**.

Por ejemplo, si queremos acceder a:

```text
M[AA]
```

se coloca la dirección `AA` en el MAR:

```text
MAR = AA
```

La memoria utiliza esa dirección para determinar qué posición debe leer o escribir.

**En resumen:**

> El MAR indica **dónde está el dato que queremos acceder**.

---

## 4. Registro de Datos de Memoria (MDR)

El **Registro de Datos de Memoria** almacena temporalmente el **dato que se está transfiriendo entre la memoria y la CPU**.

Por ejemplo, si:

```text
M[AA] = 10100110
```

al leer esa posición:

```text
MDR = 10100110
```

Luego, ese dato puede ser transferido a un registro general:

```text
R1 = 10100110
```

El MDR también se utiliza cuando la CPU quiere **escribir** un dato en memoria.

**En resumen:**

> El MDR contiene **el dato que entra o sale de la memoria**.

---

# 5. Registros generales

Los **registros generales** (`R0`, `R1`, `R2`, etc.) almacenan datos que la CPU necesita para realizar operaciones.

Por ejemplo:

```text
R1 = 10100110
R2 = 10000000
```

Podemos realizar:

```text
R3 = R1 AND R2
```

Los registros generales son especialmente importantes para las **operaciones aritméticas y lógicas**.

---

# Ejemplo completo

Supongamos que tenemos:

```text
Dirección    Contenido
20           11AA
22           2080
```

Y queremos ejecutar la instrucción ubicada en `20`.

### 1. El PC indica dónde está la próxima instrucción

```text
PC = 20
```

La CPU sabe que debe buscar la instrucción en `20`.

### 2. La dirección pasa al MAR

```text
MAR = PC
MAR = 20
```

### 3. Se lee la memoria

```text
M[20] = 11AA
```

La instrucción obtenida se carga en el RI:

```text
RI = 11AA
```

### 4. El PC avanza

La CPU actualiza el PC para apuntar a la siguiente instrucción:

```text
PC = 22
```

### 5. La Unidad de Control interpreta la instrucción

Ahora la CPU analiza:

```text
RI = 11AA
```

y determina qué operación debe realizar.

---

# Resumen

| Registro          | Contiene                            | Pregunta que responde                |
| ----------------- | ----------------------------------- | ------------------------------------ |
| **PC**            | Dirección de la próxima instrucción | ¿Dónde está la próxima instrucción?  |
| **RI / IR**       | Instrucción actual                  | ¿Qué tengo que hacer?                |
| **MAR**           | Dirección de memoria                | ¿Dónde tengo que acceder?            |
| **MDR**           | Dato transferido                    | ¿Qué dato estoy leyendo/escribiendo? |
| **R0, R1, R2...** | Datos/operandos                     | ¿Con qué datos trabajo?              |

Una forma sencilla de recordarlos:

```text
PC  → DÓNDE está la próxima instrucción
RI  → QUÉ hacer
MAR → DÓNDE acceder en memoria
MDR → QUÉ dato transferir
Rn  → CON QUÉ datos trabajar
```

> **Nota:** Los nombres pueden variar según la arquitectura o la bibliografía. Por ejemplo, el Registro de Dirección de Memoria puede aparecer como **MAR**, y el Registro de Datos de Memoria como **MDR**, **MBR** o incluso con otras denominaciones.
