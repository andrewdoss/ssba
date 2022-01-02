package main

import "fmt"

func testLoop() {
	for i := 0; i < 3; i++ {
		fmt.Println((i * 3) + 1)
		fmt.Println((i * 3) + 2)
		fmt.Println((i * 3) + 3)
	}
}

func main() {
	testLoop()
}
