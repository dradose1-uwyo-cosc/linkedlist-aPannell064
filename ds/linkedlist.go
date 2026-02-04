package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// insert at the tail
func (l *LinkedList) Insert(value string) {
	// Create a new node
	newNode := &Node{data: value}

	// If the list is empty, set Head and Tail to the new node.
	// Otherwise, just add it to the end and update Tail.
	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

// inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) InsertAt(position int, value string) error {
	// Handle invalid position and insertion at the ends
	if position > l.Size || position < 0 {
		return errors.New("invalid position")
	} else if position == l.Size {
		l.Insert(value)
		return nil
	} else if position == 0 {
		newNode := &Node{data: value}
		newNode.Next = l.Head
		l.Head = newNode
		l.Size++
		return nil
	}

	// Iterate to the position before the desired position
	current := l.Head.Next
	for i := 1; i < position-1; i++ {
		current = current.Next
	}

	// Insert the new node at the desired position
	tmp := current.Next
	current.Next = &Node{data: value}
	current.Next.Next = tmp
	l.Size++
	return nil
}

// remove first occurrence of the value
func (l *LinkedList) Remove(value string) error {
	current := l.Head
	previous := &Node{}

	// Traverse the list to find the node to remove
	for current != nil {
		if current.data == value {
			if current == l.Head {
				// Change head if needed
				l.Head = current.Next
				if current == l.Tail {
					l.Tail = nil
					l.Size--
					return nil
				}
			} else {
				// Skip the current node to remove it
				previous.Next = current.Next
			}
			if current == l.Tail {
				// Update tail if needed
				l.Tail = previous
				l.Tail.Next = nil
			}
			l.Size--
			return nil
		}
		previous = current
		current = current.Next
	}
	return errors.New("value not found")
}

// remove all occurrences of a value.
// This works about the same as Remove but keeps going after a removal
func (l *LinkedList) RemoveAll(value string) error {
	err := errors.New("value not found")
	current := l.Head
	previous := &Node{}

	// Traverse the list to find nodes to remove
	for current != nil {
		// If the current node matches the value, remove it
		if current.data == value {
			// Update head if needed
			if current == l.Head {
				l.Head = current.Next
				// If it was also the tail, update tail
				if current == l.Tail {
					l.Tail = nil
					l.Size--
					return nil
				}
			} else {
				// Skip the current node to remove it
				previous.Next = current.Next
			}
			if current == l.Tail {
				// Update tail if needed
				l.Tail = previous
				l.Tail.Next = nil
			}
			l.Size--
			err = nil
		} else {
			previous = current
		}
		current = current.Next
	}
	return err
}

// remove at a position, if index exists
func (l *LinkedList) RemoveAt(pos int) error {
	// Handle invalid position
	if pos >= l.Size || pos < 0 {
		return errors.New("invalid position")
	} else if pos == 0 {
		// Remove head and update tail if they are the same
		l.Head = l.Head.Next
		if l.Head == l.Tail {
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	current := l.Head.Next
	previous := l.Head
	// Traverse to the position
	for i := 1; i < pos; i++ {
		previous = current
		current = current.Next
	}

	// Remove the node at the position
	if current == l.Tail {
		l.Tail = previous
		l.Tail.Next = nil
	} else {
		previous.Next = current.Next
	}
	l.Size--
	return nil
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool { return l.Size == 0 }

// get size of ll
func (l *LinkedList) GetSize() int { return l.Size }

// reverse the list
func (l *LinkedList) Reverse() {
	var previous *Node = nil
	current := l.Head
	var next *Node

	l.Tail = current
	// Traverse the list and reverse the links
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	l.Head = previous

}

func (l *LinkedList) PrintList() {
	fmt.Print("[")
	current := l.Head
	// Traverse the list and print each node's data
	for current != nil {
		fmt.Print(current.data)
		if current.Next != nil {
			fmt.Print(" -> ")
		} else {

		}
		current = current.Next
	}
	fmt.Println("]")
}
