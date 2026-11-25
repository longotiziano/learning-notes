@5)Este programa halla la suma de dos números mediante una función que los recibe por la pila
@ https://cpulator.01xz.net/?sys=arm-de1soc

			.data				
        @ direcciones de memoria para los datos
			numero1:	.word 2	
			numero2:	.word 5
			resultado:	.word 0	

		   .text	
			.global main
	
main: 		
			@ como es procesador RISC, necesitamos un intermediario para encontrar el dato q esta en numero1
			ldr r0, =numero1 @ = -> "mete esta secuencia de 32 bits"
			@ el [] es fundamental pq sino copiamos la direc de memo. [] = *
			ldr r1, [r0]
			
			ldr r0, =numero2
			ldr r2, [r0]
			@ las llaves son conjuntos de registros, y son lo que utilizan los pops y pushs
			push {r1, r2}   @push {r2}
			@ bl = branch with link. Llama funciones 
			bl suma         @PC = dirección de suma LR = dirección de pop {r6}. LINK REGISTER (LR) = R14
			
			pop {r6}
			ldr r0, =resultado
			str r6,[r0]
			b fin
			
suma:		pop {r4, r5}     @pop {r5}
			add r3, r4, r5
			push {r3}
			bx lr            @ PC = LR, Branch and Exchange to Link Register
			
fin:		.end
