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

func (s *Counter) Add() {
	s.qtd++
}

func (s *Counter) Print() {
	fmt.Printf("%s counter: %d\n", s.label, s.qtd)
}
