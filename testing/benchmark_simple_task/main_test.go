package main

import (
	"testing"
)

func BenchmarkAtomicCounter(b *testing.B) {
	input := 5000

	for b.Loop() {
		_ = GoroutineAtomicCounter(input)
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	input := 5000

	for b.Loop() {
		_ = GoroutineMutexCounter(input)
	}
}
