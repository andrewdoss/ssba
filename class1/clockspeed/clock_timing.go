package clockspeed

import (
	"fmt"
	"time"
)

const iter = 1000000000
const trials = 10
const factor = 1.001

func ClockSpeed() {
	total := 0.0
	for i := 0; i < trials; i++ {
		start := time.Now()
		testLoop()
		elapsed := time.Since(start)         // ns
		frequency := iter / float64(elapsed) // GHz
		total += frequency
		fmt.Printf("Estimated clock-speed is %4.2f GHz.\n", frequency)
	}
	fmt.Printf("Average estimate over %d trials is %4.2f GHz.", trials, total/trials)
}

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

func testLoopMult() float64 {
	x := 0.0
	for i := 0; i < iter; i++ {
		x *= factor
	}
	return x
}

func testLoopDivide() float64 {
	x := 0.0
	for i := 0; i < iter; i++ {
		x /= factor
	}
	return x
}
