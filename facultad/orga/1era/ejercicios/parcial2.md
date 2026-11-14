# 1. Conversión y aritmética binaria

## a. 

### 1. 1 1010 1101 

#### Hexadecimal
Dado que tenemos la relación de 2^4 = 16, podemos hacer la conversión directa, donde 4 dígitos del número binario representan un número hexadecimal.

Como tenemos 9 dígitos, tenemos que rellenar el número de 0s para completar los 12 dígitos (1 1010 1101 -> 0001 1010 1101)
- 0001: 1
- 1010: A
- 1101: C

Número: 1AC

#### Decimal
```
1x2^8 + 1x2^7 + 0 + 1x2^5 + 0 + 1x2^3 + 1x2^2 + 0 + 1 =
256 + 128 + 32 + 8 + 4 + 1 = 529
```

## b.

### 1.
```
  01101101
+ 00110111
  10100100
```
No hay ni carry ni overflow, ya que no hay acarreo del bit más significativo, y el resultado entra en el rango de representación sin signo.

### 2.
```
  11110110
+ 00101101
1 00100011
```
Interpretado sin signo, hubo un acarreo en el bit más significativo, por lo tanto hay carry y, además, el resultado excede el rango de representación de 8 bits, por lo tanto hay overflow.

### 3.
```
  10000001
+ 11111111
1 10000000
```
Interpretado con signo, hubo un acarreo (mismas razones habladas anteriormente), sin embargo no hubo overflow, ya que la suma de dos negativos dio un negativo, lo que significa que está dentro del rango de representación con signo.

# 2. -45,625 -> IEEE 754
```
45 / 2
1   22 / 2
    0   11 / 2
        1   5 / 2
            1   2 / 2
                0   1 / 2
                    1
```
1. Parte entera: 101101
2. Parte decimal: 0.625x2 = 1,250x2 = 0,5x2 = 1,0 -> 0101
3. Exponente: 101101,0101 = 1,011010101x2^5
4. Bit de signo: 1
5. Exponente con sesgo: 5 + 127 = [132]10 = [1000 0101]
6. Configuración final del número (1 bit signo, 8 bits exponente, 23 bits de mantisa):
```
1 10000101 01101010100000000000000
```

# 3. -6942 a BCD empaquetado
Ya que en BCD cada dígito es un nibble (4 bits) representa un número decimal, y el último nibble el signo del número (CAFE para positivos y BD para negativos).

## 1. Hexadecimal
- 6 = 0110
- 9 = 1001
- 4 = 0100
- 2 = 0010
- Signo = B = 1011
Número: 0110 1001 0100 0010 1011

## 2. Representación binaria
```
6942 / 2
0    3471 / 2
    1   1735 / 2
        1   867 / 2
            1   433 / 2
                1   216 / 2
                    0   108 / 2
                        0   54 / 2
                            0   27 / 2
                                1   13 / 2
                                    1   6 / 2
                                        0   3 / 2
                                            1   1 / 2
                                                1
```
El número positivo es: 1101100011110, y para encontrar el número negativo hayamos el complemento a la base + 1.

```
1101100011110 -> 0010011100001 + 1 = 0010011100010
```
Resultado final: 0010011100010

## 3. Cantidad total de bits
Ya que cada dígito es un nibble (4 bits), podemos multiplicar la cantidad de nibbles por 4 para obtener los bits. En este caso, 5x4 = 20 bits.

## 4. Por qué es distinto
Es distinto ya que no únicamente que ocupa menos tamaño, si no que en el caso de la conversión directa sería únicamente una secuencia binaria, mientras que el otro tiene un formato particular.

# 4. Circuito combinacional

## a. Tabla de verdad
```
A B C S
0 0 0 0
1 0 0 0
1 1 0 1
1 1 1 1
0 1 1 1
0 0 1 0
1 0 1 1
0 1 0 0
```

## b. S(A, B, C)
S(A, B, C) = AB + AC + BC

## c. 
No le veo mucho sentido a simplificar acá

## d. 
No puedo dibujar en markdown pero sé hacerlo

## e. 
Porque rompe con las condiciones de la función. XOR para devolver 1, debe recibir un único 1, de lo contrario será 0. Por lo tanto fallaría en todos los casos.

# 5. Punto teórico

## a. Registros internos
Función de los registros internos y comportamiento durante la lectura de memoria.
- PC: Program Counter. Indica la dirección de memoria de la próxima instrucción.
- RI: Dirección de la instrucción actual. Es quien dice qué hacer.
- RD: El dato del registro obtenido en la lectura.
- RDM: La dirección de memoria asociada a la instrucción. La ubicación que vamos a buscar.

## b. Direccionamiento
