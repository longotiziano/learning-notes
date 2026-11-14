name "orga"

; Este ejemplo imprime "Organizacion del Computador" 
; escribiendo directamente en la memoria de video.
; En la memoria VGA VRAM: el primer byte es un caracter ASCII,
; el siguiente byte es un atributo de caracter. 
; Si cambiamos el segundo byte, puede cambiar el color del caracter
; incluso despues de imprimirlo. El atributo de caracter es un valor de 8 bits.
; Los 4 bits mas significativos (MSB) establecen el color de fondo y 
; los 4 bits menos significativos (LSB) establecen el color de primer plano.


; hex    bin        color
; 
; 0      0000      negro
; 1      0001      azul
; 2      0010      verde
; 3      0011      cyan
; 4      0100      rojo
; 5      0101      magenta
; 6      0110      marron
; 7      0111      gris claro
; 8      1000      gris oscuro
; 9      1001      azul claro
; a      1010      verde claro
; b      1011      cyan claro
; c      1100      rojo claro
; d      1101      magenta claro
; e      1110      amarillo
; f      1111      blanco
 


org 100h

; Modo de video 3 (texto a color, 80 columnas por 25 filas de manera predeterminada)    
mov ax, 3     ; mode texto 80x25, 16 colores, 8 paginas (ax = 3 significa ah=0, al=3)
int 10h       ; ejecuta!

; Cancela el destello y establece los colores:
mov ax, 1003h ;al = 03 permite que el MSB controle la intensidad del color del caracter 
mov bx, 0 ;Cancela el destello
int 10h


; Define el registro de segmento en la VRAM para texto a color (B8000):
mov     ax, 0b800h
mov     ds, ax

; imprime Organizacion del Computador"
; El promer byte es un codigo ADCII ([02h] es O), el segundo el codigo de color([03h]).

mov [02h], 'O'

mov [04h], 'r'

mov [06h], 'g'

mov [08h], 'a'

mov [0ah], 'n'

mov [0ch], 'i'

mov [0eh], 'z'
 
mov [10h], 'a'

mov [12h], 'c'

mov [14h], 'i'

mov [16h], 'o'

mov [18h], 'n'

mov [1ah], ' '

mov [1ch], 'd'

mov [1eh], 'e'

mov [20h], 'l'

mov [22h], ' '

mov [24h], 'C'

mov [26h], 'o'
 
mov [28h], 'm'

mov [2ah], 'p'

mov [2ch], 'u'

mov [2eh], 't'

mov [30h], 'a'
   
mov [32h], 'd'

mov [34h], 'o'

mov [36h], 'r'


; colorea todos los cracteres:
mov cx, 27  ; numero de caracteres.
mov di, 03h ; Comienza desde el byte siguiente a 'O'

c:  mov [di], 11101100b   ; rojo claro(1100) sobre amarillo(1110)
    add di, 2 ; busca el siguiente codigo ascii en la memoria vga.
    loop c

; espera que se aprete alguna tecla para devolver el control al sistema operativo:
mov ah, 0
int 16h

ret

