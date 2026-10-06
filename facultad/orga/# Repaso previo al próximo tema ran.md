\# Repaso previo al próximo tema: rangos



Como definir el rango de un número:

\-> Teniendo un número con signo y n bits, rango de representación es...

\[-2^(n-1), +2^(n-1)-1]



Por ejemplo:

&#x20;00100101 = 1 + 4 + 32 = 37 en base 10

&#x20;01001011 = 1 + 2 + 8 + 64 = 75 en base 10

con n = 8, entonces rango es \[-128, +127], por lo tanto



37 - 75 = -38, por lo tanto está dentro del rango



\-> Si operamos sin signo, entonces...

\[0, 2^n -1]



Por ejemplo:

&#x20; 11001 A  25 en base 10

\- 10011 B  19 en base 10

&#x20;

&#x20; 01101 CB(B)

\+ 11001 A

1 00110  <- 6 en base 10, por lo tanto no se fue de rango, con C = 1



**IMPORTANTE**: cabe recalcar que en las operaciones **SIN SIGNO** la aparición

de carry representa que **NO está en rango.** 



Overflow:

\- Con signo: Si suma de números positivos da negativo

\- Sin signo: Si hay carry



Flag N:

\- Puede ocurrir incluso cuando realizamos operaciones sin signo, ya que

lo que revisa es que el primer dígito sea un 1.



\# Empaquetado BCB (Decimal Codificado en Binario)



Un byte está formado por 2 nibbles (nibble es una unidad de 4 bits)



Para almacenar un número, se debe pasar el número a base 10, luego 

colocar cada dígito en los nibbles dejando el último libre que 

representaría el signo, utilizando:

\-> C, A, F, E para positivos

\-> B, D para negativos

Ej: +123\[10] -> 00123A\[16]



\## Formato Zoneado

Representa positivos y negativos en hexadecimal



Se llama de esta forma porque se puede dividir el número en zonas, rellenando con F0 hasta alcanzar la cantidad de bytes usados.



\## Ejercicios

302111222 base 4

1. Pasarlo a base 16
2. Ver que tenga dígitos a la izquierda y signo en el nibble menos significativo
3. Escribir el número decimal



11 0010 0101 0110 1010

3   2     5    6    A  <- Empaquetado



\-> **Observación**: puedo agarrar cualquier dígito en base 4, siempre y 

cuando al momento se hacer el pase a hexadecimal, el resultante

tenga un dígito del 0 al 9 a la izquierda de la letra.



\### 20

a. 110236461011004652523442117]8

b. 001010110011010100111000001110010011000000100100]2



a. 001 001 000 010 011 110 100 110 001 000 001 001 000 000 100 110 101

010 101 010 011 100 100 010 001 001 111



Paso a hexadecimal:



0000 0100 1000 0100 1111 0100 1100 0100 0001 0010 0000 0100 1101 0101 

0101 0100 1110 0100 0100 0100 1111



0484F4C41204D554E444F



\# Representación de decimales y puntos flotantes (IEEE 754)

1,01110 x 2^(5)



Este 1 de la parte entera se llama bit implícito, y es la parte 

representativa del número. En caso de que sea un 1, entonces es número

negativo, si es 0 entonces positivo (esta parte ocupa 1 bit).



Después tenemos la mantisa que es la parte decimal del número (que 

ocupa 23 bits). Y finalmente tenemos el exponente (que ocupa 8 bits, 

sumando entre todos 32 bits).



\## Doble precisión (64-bits)

El signo ocupa 1 bit, el exponente 11 y la mantisa 52.



\## Ejercicios

12a. 0.100111\[2] positivo, mandamos coma hacia la derecha para llevar

al número a forma 1.(mantisa) y nos queda 1.0111x2^-1. 



Después tendríamos que completar los 23 bits de la mantisa con ceros.









