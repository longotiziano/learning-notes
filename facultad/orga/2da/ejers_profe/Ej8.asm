
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt

org 100h

JMP inicio
nums DB 09h,01h,04h,40h,02h,00h,04h,00h,00h,01h
resultado DB 0h
;Mostrar en la pantalla la suma de 10 datos.
inicio:
    LEA SI, nums    ;puntero al primer dato
    MOV CX, 10      ;10 vueltas
    

repetir:  
     MOV AL, [SI]    ;traigo dato
     ADD resultado, AL      ;lo sumo al resultado
     INC SI          ;sig dato
     LOOP repetir
     
MOV AL, resultado    ;AL= resultado
AAM                  ;AH=decenas,AL=unidades
ADD AH, 30h          ;lo ubico a la region de nros de ascii
ADD AL, 30h          ;idem arriba
MOV BX, AX           ;BH=decenasy BL=unidades

MOV DL, BH           ;Imprimo la decena
MOV AH, 02h
INT 21h
   
MOV DL, BL
MOV AH, 02h          ;imprimo la unidad
INT 21h
 

ret




