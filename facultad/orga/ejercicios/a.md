Se desea complementar los cuatro bits intermedios de un byte
dejando los otros cuatro bits intactos. ¿Qué máscara y qué 
operación debe usar? Resolver el programa. 

El dato está en la dirección AA.
Guardar el resultado en AB.

```
20 11AA R1 = M[AA]
22 223C R2 = 3C = 0011 1100
24 9312 R3 = R1 XOR R2
26 33AB M[AB] = R3
28 C000
```

En la maquina de brookshear cada instrucción tiene 4 dígitos hexadecimales (2 bytes):
```
[Dirección de Memo de la instrucción] [Código de operación] [Registro destino] [Dirección de memoria objetivo de la operación]
```

Por ejemplo, en nuestro ejemplo:
```
20 (1) Cargá, en el registro 1, el dato de la dirección de memoria AA (M[AA]).
22 (2) Cargá, en el registro 2, ESTE dato (no tiene nada que ver con la memoria)
```

> Complementar un bit significa invertirlo

> Máscara: número que elegís para comparar el valor original

La máscara 0011 1100 con la operación XOR resulta ideal, ya que, en la máscara, 0 deja el bit original
y 1 lo invierte.