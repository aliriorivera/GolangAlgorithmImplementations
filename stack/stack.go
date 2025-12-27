package stack

// El Stack es LIFO : Last In Frist Out.
// el ultimo que ingresa es el primero que sale, este es ideal para reversar un arreglo
// entra 1 2 3  sale 3 2 1

import (
	"errors"
)

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: []int{}}
}

func (s *Stack) Push(newValue int) {
	s.values = append(s.values, newValue)
}

func (s *Stack) Pop() (int, error) {
	if len(s.values) == 0 {
		return 0, errors.New("Stack is empty")
	}
	last := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return last, nil
}

func (s *Stack) IsStackEmpty() bool {
	return len(s.values) == 0
}
