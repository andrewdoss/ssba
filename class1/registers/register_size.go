package main

import (
	"fmt"
	"unsafe"
)

// Short program to demonstrate register size
// Assumes non-sized "int" takes register value
func main() {
	last := 1
	for i := 0; i < 100; i++ {
		if (1 << i) < last {
			fmt.Println(i + 1) // Should print 64
		}
		last = 1 << i
	}
	// Can also print size of a pointer
	fmt.Println(8 * unsafe.Sizeof(&last))
}
