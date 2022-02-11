package clockspeed

import (
	"testing"
)

func BenchmarkTestLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testLoop()
	}
}

func BenchmarkTestLoopAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testLoopAdd()
	}
}

func BenchmarkTestLoopMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testLoopMult()
	}
}

func BenchmarkTestLoopDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testLoopDivide()
	}
}
