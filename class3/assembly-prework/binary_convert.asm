section .text
global binary_convert
binary_convert:
	xor 	rax, rax		; clear accumulator
start_loop:
	movzx	rdx, byte [rdi]	; move next byte from string
	cmp		rdx, 0
	je		end_loop
	shl		rax, 1			; multiply by 2
	inc		rdi				; increment string index
	cmp		rdx, 0x30			; check if value is '0'
	je		start_loop
	inc		rax
	jmp		start_loop
end_loop:
	ret
