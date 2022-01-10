package cachelookup

import (
	"testing"
)

func BenchmarkRowwiseLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rowwiseLookup()
	}
}

func BenchmarkColumnwiseLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		columnwiseLookup()
	}
}
