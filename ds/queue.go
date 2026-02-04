package ds

import (
	"errors"
)

type Queue struct {
	list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

// remove the head
func (q *Queue) Pop() (string, error) {
	if q.list.IsEmpty() {
		return "", errors.New("queue is empty")
	}

	head := q.list.Head.data
	err := q.list.RemoveAt(0)
	return head, err
}
