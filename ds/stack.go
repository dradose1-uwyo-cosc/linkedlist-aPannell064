package ds

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.InsertAt(0, value)
}

// remove the head
/*
	Returns a string and a bool value because that
	is how the assignment specified it should be done
	even through it's inconsistent with other methods in
	the project.
*/
func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty() {
		return "", false
	}

	head := s.list.Head.data
	err := s.list.RemoveAt(0)
	return head, err == nil
}
