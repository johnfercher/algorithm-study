package domain

import "fmt"

type Object struct {
	ID   int
	Name string
}

func (s *Object) Print() {
	fmt.Printf("ID: %d, Name: %s\n", s.ID, s.Name)
}

type Objects []*Object

func (s Objects) Print() {
	for _, value := range s {
		value.Print()
	}
}
