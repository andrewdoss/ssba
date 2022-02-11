section .text
global sum_to_n
sum_to_n:
; Looping answer below
; 			mov		rax, 0			; initialize return register to 0
; 			mov		rcx, 0			; set loop incrementer to 0
; 							; n is stored in rdi by the caller
; start_loop: 
; 			cmp		rcx, rdi		; compare incrementer to n
; 			jnle   	end_loop		; break from loop if incrementer not <= n
; 			add		rax, rcx		; add incrementer to accumulator
; 			inc		rcx
; 			jmp		start_loop
; end_loop:
; 	ret
; Closed form answer below
sum_to_n:
			xor		rdx, rdx
			mov		rax, 1
			mov		rsi, 2
			add		rax, rdi
			imul	rax, rdi
			idiv 	rsi
	ret
	
