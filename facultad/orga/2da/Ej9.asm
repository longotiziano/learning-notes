; Un byte se encuentra en la dirección 0100. Ubicar en la dirección 0101 su
; nible más significativo. 
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt

org 100h

; Un byte se encuentra en la direccion 0100. Ubicar en la direccion 0101 su
;nible mas significativo.
MOV BYTE PTR [0100h], 0C6h    ;[0100h] = C6
                              
MOV AL, [0100h]               ;Me ubico en 0100
MOV CL, 4                     ;Me desplazo 4 a la derecha
SHR AL, CL                    ;En AL guardo el nibble alto, shr desplaza 4 bits a la derecha 
MOV [0101h], AL               ;En 0101 guardo lo que hay en AL quedaria 0000 1100, en teoria
    

ret




