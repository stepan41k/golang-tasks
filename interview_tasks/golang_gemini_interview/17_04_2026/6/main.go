package main

import (
	"fmt"
	"testing"
)

type Performer interface {
	Do()
}

type FastPerformer struct {
	ID int
}

func (f FastPerformer) Do() {}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := FastPerformer{ID: i}
		execute(p)
	}
}

func execute(p Performer) {
	p.Do()
}

// Происходит ли здесь аллокация в куче (heap) при вызове execute(p)? Почему?
// Что изменится, если Do() будет принимать указатель: func (f *FastPerformer) Do()?
// Как escape analysis (анализ побега) влияет на интерфейсы?
