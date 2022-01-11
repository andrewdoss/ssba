section .text
global pangram
pangram:
start_loop:
	movzx	rdx, byte [rdi]	; move next byte from string
	cmp		rdx, 0			; check for null terminator
	je		end_loop
	cmp		rdx, 97		; if < 'a': shift to lowercase
	jl		to_lower
end_to_lower:
	cmp		rdx, 97		; if still not within 'a'-'z', continue
	jl		continue
	cmp		rdx, 122
	jnle	continue
	; At this point, the byte is a lower case ASCII letter
	sub		rdx, 97			; normalize as offset from 'a'
	bts		rax, rdx		; set bit at offset position to '1'
	jmp     continue
end_loop:
	; Test whether all lower 26 bits are set
	add		rax, 1
	mov		rbx, 1
	shl		rbx, 26
	cmp		rax, rbx
	je		return_true		; jump to return 1 if equal
	; Set to 0 and return
	mov		rax, 0
	ret
return_true:
	mov		rax, 1
	ret
to_lower:
	add		rdx, 32
	jmp		end_to_lower
continue:
	inc		rdi
	jmp		start_loop	
