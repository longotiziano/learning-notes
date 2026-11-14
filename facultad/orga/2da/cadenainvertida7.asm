;la pila
.model small

.data

unNombre db 'Microprocesador $'
tamNombre = ($ - unNombre) - 1

.code
main PROC
    
    mov ax,@data
    mov ds,ax
    
    ; Pone el nombre en la pila.
    mov cx,tamNombre
    mov si,0
    
L1: mov al,unNombre[si]; obtine el caracter
    push ax ; lo pone en la pila
    inc si
    loop L1
    
    ; Saca el nombre de la pila en orden inverso,
    ;y lo almacena en el arreglo unNombre.
    mov cx,tamNombre
    mov si,0
    
L2: pop ax ; obtiene el caracter
    mov unNombre[si],al;lo almacena en la cadena
    inc si
    loop L2
    
    mov ah,9
    lea dx, unNombre
    int 21h
    
    hlt
;Muestra el nombre.
   ; mov dx, OFFSET unNombre
   ; call Writestring
    ; call Crlf
main ENDP
END main