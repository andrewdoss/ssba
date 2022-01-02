package clockspeed

const iter = 1000000000

func testLoop() {
	for i := 0; i < iter; i++ {
	}
}

func testLoopAdd() int {
	x := 0
	for i := 0; i < iter; i++ {
		x += 1
	}
	return x
}

func testLoopMult() int {
	x := 0
	for i := 0; i < iter; i++ {
		x = i * 146
	}
	return x
}

func testLoopDivide() int {
	x := 0
	for i := 0; i < iter; i++ {
		x = i / 146
	}
	return x
}
