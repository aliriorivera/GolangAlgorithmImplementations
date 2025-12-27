package stack

import "errors"

type StackWithArray struct {
	index  int
	values []int
	max    int
}

func NewStackWithArray() *StackWithArray {
	return &StackWithArray{
		index:  -1,
		values: make([]int, 10),
		max:    10,
	}
}

func (s *StackWithArray) Pop() (int, error) {
	if s.index < 0 {
		return s.index, errors.New("Stack is empty")
	}
	currentValue := s.values[s.index]
	s.index--
	return currentValue, nil
}

func (s *StackWithArray) Push(newValue int) error {
	if s.index+1 >= s.max {
		return errors.New("no more capacity")
	}
	s.index++
	s.values[s.index] = newValue
	return nil
}

func (s *StackWithArray) IsEmpty() bool {
	return s.index < 0
}
