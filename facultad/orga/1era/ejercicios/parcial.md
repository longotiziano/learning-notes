# Práctica de parcial

## Ejercicio 1.
- A = A3
- B = B7
Hacer A-B y A+B tomando los números en representación de enteros con signo y sin signo en 8 bits. Indicar todos los flags.

### Solución
A = 1010 0011, B = 1011 0111. Dado que B es negativo en el primer caso, saco el complemento a la base + 1, que nos queda 
entonces B = 0100 1001.
```
  1010 0011
+ 0100 1001
  1110 1100
```

Luego, la suma:
```
  1010 0011
+ 1011 0111
1 0101 1110 
```

### Flags
Con signo:
- C = 0 -> no hubo dígitos fuera de los 8 bits
- V = 0 -> Es quien me define si hubo un resultado válido, en el caso del signo. Al no haber overflow, es válido el resultado
- N = 1 -> El dígito más significativo es un 1
- Z = 0 -> El número es distinto del 0
- P = 1 -> Cantidad impar de unos

Y en la suma:
- C = 1
- V = 1
- N = 0
- Z = 0
- P = 1

> Un resultado es válido en operaciones sin signo cuando C = 0. Si V = 1, entonces es inválido en operaciones con signo.

## Ejercicio 2.

### Inciso a.
Tenemos el número -15,25 que está en decimal para pasar a punto flotante simple precisión.

```
-15,25 1111,01x10^0         0,25x2=0,5x2=1

Normalización de número:
1,11101x10^3 tengo que sumarle el exceso 127 -> 127+3=130
```
Por lo tanto, siguiento la estructura de signo-exponente-mantisa:
- S = 1
- EXP = 1000 0010
- MANTISA = 111010000...0000

### Inciso b.
C824F410 que está en PFSP. Pasarlo a decimal:
```
1 10010000 01001001111010000010000
```
Si separamos el número completo en los bits correspondentes (1, 8 y 23) podemos identificar cada componente. Usando esto,
nos queda que el número es negativo, el exponente es 127 + E = 144 -> E = 17, entonces:
```
1,01001001111010000010000x10^17
```
Corremos la coma 17 números y queda un número bastante grandecito.

## Ejercicio 3.
Tenemos 12745E como número empaquetado de 3 bytes.

### Inciso a.
¿Cuál es el valor en base 10?

Dado que finaliza en la E (y no en B o D), el número toma valor positivo, y cada dígito representa un número decimal, 
por lo tanto el valor en base 10 es +12745.

### Inciso b.
¿Y en base 7?

Sería positivo por los mismos motivos de antes, y el valor lo obtendrías luego de armar el número con los restos de cada
división.

## Ejercicio 4.
Diseñar un circuito con 4 entradas (A, B, C, D). La salida es 1 si el número es divisible por 3 o por 5.

### Inciso a.
Tabla de verdad.

### Inciso b.
Expresión lógica de verdad.
```
F(A, B, C, D) = \A\BCD + \AB\CD + \ABC\D + A\B\CD + A\BC\D + AB\C\D + ABCD
F(A, B, C, D) = \A\BCD + \AB\CD + \ABC\D + A\B\CD + A\BC\D + AB(\C\D + CD)
```

### Inciso c.
Dibujo del circuito.

## Ejercicio 5.
Hay un número almacenado en la dirección E0. Detectar si el bit 0 es 1 y el bit 7 es 0. En caso
afirmativo, poner 1 en el bit 4 y en 0 el bit 5. En caso contrario, hallar el complemento a la base
del número. 

En los 2 casos guardar el resultado en la dirección E1.

```
20 11E0 
22 2081
24 8201
26 2001  CASO FAVORABLE EN R0
28 B234  SI ES UN CASO FAVORABLE SALTA A XX

2A 23FF  MASCARA PARA XOR
2C 9413  R1 XOR R3 (complemento a la base -1)
2E 2501  R5 = 1
30 5645  R6 = R5 + 1 (complemento a la base)
32 B03C  SALTA INCONDICIONALMENTE (compara R0 con R0)

34 2710  R7 = 0001 0000
36 7817  R8 = R1 OR R7
38 29DF  R9 - 1191 1111
3A 8689  R6 - R8 AND R9
 
3C 36E1  ESCRIBIR
3E C000
```

## Ejercicio 6. Teórico
- ¿Cantidad máxima de instrucciones de la máquina de Brookshear? -> 16, ya que tiene 4 bits
- Diferencias entre modo de direccionamiento directo e indirecto -> En el indirecto te dan el lugar de memoria donde está el dato
- Que hay en la dirección de memoria del direccionamiento directo -> El dato, en el indirecto la dirección
- La máquina de Lucas tiene 3 bits, qué operaciones puede hacer la máquina de Lucas? -> 8
- Cuando la computadora se enciende, pone toda la memoria en 0? Lo vemos más adelante, pero es aleatorio.
- Qué entradas tiene la unidad aritmético lógica combinacional? La instrucción y el registro 