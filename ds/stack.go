package ds

import (
	"errors"
)

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.InsertAt(0, value)
}

// remove the head
func (s *Stack) Pop() (string, error) {
	if s.list.IsEmpty() {
		return "", errors.New("stack is empty")
	}

	head := s.list.Head.data
	err := s.list.RemoveAt(0)
	return head, err
}
