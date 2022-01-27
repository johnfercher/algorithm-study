package domain

import "fmt"

type Counter struct {
	qtd   int
	label string
}

func NewCounter(label string) *Counter {
	return &Counter{
		qtd:   0,
		label: label,
	}
}

func (s *Counter) Increment() {
	s.qtd++
}

func (s *Counter) Print() {
	fmt.Printf("%s counter: %d\n\n", s.label, s.qtd)
}
