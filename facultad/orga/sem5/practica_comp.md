# Operaciones lógicas

El primer dígito es el número de operación OR AND XOR (7 8 9 respectivamente), el segundo es donde se guarda el resultado, y por ultimo los 2 números de registro.

Ejemplo:
```
    10100100 = R1
OR  10001110 = R2
   ____________
    10101110 = R3 = R1 OR R2 -> 7312
```

### Ejercicio 1
Se lee la dirección AA que tiene un número binario con signo. 
Determinar si es un número positivo o negativo. 
Si +, poner 0 en la dirección BB. Si -, poner 1.

10100110 <-- número
10000000 <-- sólo importa el signo

¿Es 10000000? si: es negativo, no: es positivo.
```
1C 2401 Número 1 en R4
1E 2300 Número 0 en R3
20 11AA Leer y guardar en R1 un número que está en AA 10100110 <-- número
22 2080 Guardamos 10000000 en R0
24 8210 R2 = R1 AND R0
26 B22C Si R2 es igual a R0 salta a la dirección 2C 
28 33BB Guardar 0 en BB
2A B02E Saltar a FIN
2C 34BB Guardar 1 en BB
2E C000
```

## Ejercicio 2
Leer el dato que está en la dirección AA. Forzarle a ese dato el bit 3 en 1 y el bit 5 en 0 y guardarlo en BB.

### Razonamiento
Leer un dato R1 <- M[AA]
```
x x x x x x x x
7 6 5 4 3 2 1 0
x x x x Y x x x <- Bit 3
x x Y x x x x x <- Bit 5
```
Bit 3 en 1:
a. Máscara:
```
x x x x Y x x x
0 0 0 0 1 0 0 0 <-- R2
08 en R2
```
b. Operación: OR -> 
R3 = R2 OR R1   Bit 3 en 1

Para poner un 1 en un bit la máscara tiene un sólo 1 en ese bit y la operación es OR.

Bit 5 en 0:
a. Máscara
```
x x x x Y x x x <-- Bit 5
1 1 0 1 1 1 1 1 (DF) <-- R4 Máscara
```
b. Operación: AND -> 
R5 = R4 AND R1   Bit 5 en 0

Para poner un 00 en un bit la máscara tiene un sólo 1 en ese bit la operación es AND.

### Ejecución
```
20 11AA  R1 <-- M[AA]
24 2208  R2 = 00001000 máscara para forzar 1 en bit 3 
26 24DF  R$2$ = 11911111 máscara para forzar 0 en bit 5 
28 7312  R3 = R1 OR R2 Bit 3 forzado a 1
2A 8534  R5 = R2 AND R4Bit 5 forzado a 0
2C 35BB  M[BB] <-- R5
2E C000  
```
### Orden de resolución
Primero identifico las tareas, y después escribo los códigos de resolución.