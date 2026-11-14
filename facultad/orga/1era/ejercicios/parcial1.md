# 1. Sistemas de numeración y Operaciones

## 1a
1011 0110 -> Cómo cada 4 dígitos de un número hexadecimal representa un dígito en binario (por la relación 2^4 = 16), hacemos conversión directa: 
- 1011 = B
- 0110 = 6
Entonces [1011 0110]2 = [B6]16

Y para pasar 1011 0110 a decimal:
```
1*2^7 + 0 + 1*2^5 + 1*2^4 + 0 + 1*2^2 + 1*2 + 0 =
128 + 32 + 16 + 4 + 2 = 182
```
Resultado: [1011 0110]2 = [182]10

Comprobación: pasamos 172 a hexadecimal con divisiones sucesivas.
```
182 / 16 = 11 resto 6 -> 11 / 16 = 0 resto 11
```
Por lo tanto, nos queda 11 = B -> B6

## 2a
[3A7]16 a binario y decimal. Mismo procedimiento, pero al revés.
- 3 = 0011
- A = 1010
- 7 = 0111
[3A7]16 = [0011 1010 0111]2

Paso el hexadecimal a decimal:
```
3*16^2 + 10*16^1 + 7*16^0 = 3 * 256 + 160 + 7 = 768 + 167 = 935
```
[3A7]16 = [935]10

## 3a
Usando divisiones sucesivas:
```
157 / 2
1    78 / 2
    0   39 / 2
        1   19 / 2
            1   9 / 2
                1   4 / 2
                    0   2 / 2
                        0   1 / 2
                            1 --
```
[157]10 = [1001 1101]2

## b
Siendo carry el indicador de acarreo que sale del bit más significativo durante la suma, y overflow cuando un resultado de suma se sale del rango de representación (por lo tanto, depende también si es una operación con o sin signo).

### 1
```
  11111111
+ 00000001
1 00000000
```
Hay carry, y como la representación es unsigned, también hay overflow (el resultado no cabe en el espacio disponible).

### 2
```
  01111111
+ 00000001
  10000000
```
No hay carry, pero dado a que la representación es con signo, la suma de 2 valores positivos no puede dar un negativo. Esto indica que se fue de rango (overflow).

### 3
```
  10000000
+ 10000000
1 00000000
```
En este caso, hay carry (hay acarreo que sale del bit más significativo) y hay overflow, pero porque la suma de 2 números negativos está dando un positivo, y en operaciones con signo esto indica que se fue de rango.

# 2. IEE 754
Separo parte entera de la decimal:
- Entera: 13 = [1101]2
- Decimal: 0.25 x 2 = 0.5 x 2 = 1.0 -> 0.25 = [0.01]2
- Número en binario: 1101.01 -> 1.10101x2^3
- Exponente: 3 + 127 = [130]10 = 1000 0010
- Mantisa: 10101 00000 0000000000 000
- Signo: 1
Número:
```
1 10000010 10101000000000000000000
```

# 3. BCD empaquetado
Para representar un número en BCD, necesitamos colocar en cada dígito un número decimal, exceptuando por el último el cual indicará el signo: CAFE para positivo y BD para negativo.

### a.
-5837 = 5837B empaquetado en BCD. Dado que cada dígito es un nibble (4 bits), representamos el número en hexadecimal:
- 5: 0101
- 8: 1000
- 3: 0011
- 7: 0111
- B: 1011
- Configuración final: 01011000001101111011

### b.
Para obtener su tamaño, multiplicamos por cuatro la cantidad de dígitos. En este caso, 5x4 = 20.

### c.
- Binario convencional: representa el valor completo mediante una secuencia binaria.
- BCD: representa cada dígito decimal por separado.

# 4. Circuitos combinacionales

## a.
```
A B S
0 0 0
1 0 1
0 1 1
1 1 0
```

## b.
Representa la operación lógica XOR

## c.
S = A._B + _A.B

## d.
No puedo dibujar en markdown pero se hacerlo.

# 5. Teoría

## a.
- Formato: es el espacio físico y el significado que le da la computadora a cada uno de los bits.
- Configuración: es la aplicación de un formato a un caso particular. Por ejemplo, el paso de un número decimal como -13,25 a IEEE 754.

## b.
- PC (Program Counter): registro interno que guarda la dirección de memoria de la próxima instrucción.
- RI: es el registro que contiene a la instrucción del procesador: 1, 2, ..., F.
- RD: guarda la dirección del registro temporal.
- RDM: es la dirección de memoria a la que queremos acceder.

## c.
- Buscar: el procesador busca en memoria la instrucción.
- Decodificar: la unidad de control interpreta la instrucción, y la comunica con la unidad aritmético lógica.
- La UAL ejecuta la instrucción, enviando además los flags.

## d.
- Carry: representa el acarreo a partir del bit más significativo en una suma.
- Overflow: indica si el resultado de una suma está fuera del rango de representación.

## e.
Significa que, dentro de la instrucción, se especifica de manera directa la dirección de memoria objetivo de la operación.

En el caso del direccionamiento indirecto, en las instrucciones se coloca la dirección de memoria, aún sin especificar la operación.

## f. 
Ya expliqué en el punto e. No está disponible en la máquina de Brookshear debido a que es un concepto didáctico minimalista para poder entender el funcionamiento de las computadoras a menor escala.

# 6. Máquina de Brookshear
Para hayar el complemento de un número, podemos utilizar la operación lógica XOR (nro 9).

## a.
Utilizaría la máscara 00111100 (3B).

## b.
Utilizaría la operación lógica XOR, ya que cuando compara cualquier número con 0, no altera el resultado. En cambio, cuando se compara con 1, siempre alterna su dígito. 

Por lo tanto, con la máscara 00111100 estaría cambiando los 4 bits intermedios, dejando intactos los otros 4.

## c.
```
00 11AA # Leo el dato en AA
02 203B # Cargo la máscara
04 9310 # Realizo la operación
06 33BB # Guardo en BB
08 C000 # Finalizo la ejecución
```