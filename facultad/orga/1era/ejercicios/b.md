El puerto de un dispositivo de E/S está en la dirección CA.
Si los bits 1 y 3 recibidos son "0", debe enviar un "1" por el bit 7.  
Si los bits 1 y 3 recibidos son "1", debe enviar un "0" por el bit 6.  

En los demás casos debe enviar un "0" por el bit 5 y un "1" por el bit 4.
En ningún caso debe modificar los bits restantes al escribir en el puerto.

```
00 210A -> 0000 1010
02 12CA -> OG
04 8012 -> R1 AND R2 en R0
06 2300 -> R3 = 0000 0000
08 B118 -> Si el 0000 1010 es igual al resultado del AND...
0A B322
-> Estoy en el 3er caso, por lo tanto: XX01 XXXX
-> Primero uso 0011 0000 con OR
-> Despues uso 0010 0000 con XOR para asegurarme que va a haber un 0
0C 2430 -> OR
0E 2520 -> XOR
10 7642 -> R4 o R2 (R2 es el número original)
12 9756 -> XOR con el resultado recién obtenido
14 37CA
16 C000
18 2450 -> Estamos en el 2do caso, y hago el mismo metodo de OR (con 0100 0000) -> XOR
1A 7542
1C 9645 -> Reutilizamos la máscara, pero con XOR del resultado del OR
1E 36CA
20 C000
22 2480 -> 1er caso, con un OR ya nos aseguramos el 1
24 7542
26 35CA
28 C000
```
0000 0010 -> XX01 XXXX
0000 1000 -> XX01 XXXX
0000 0000 -> 1XXX 0X0X
0000 1010 -> X0XX 1X1X

1111 0110 + 
0000 1010
0000 0000