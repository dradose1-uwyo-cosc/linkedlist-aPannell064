// Adam Pannell
// COSC 3750
// 02/05/2026

package main

import (
	"fmt"
	"hw-linked/ds"
)

func main() {
	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	linkedlist.Remove("first")
	linkedlist.PrintList()
	fmt.Println("List empty?", linkedlist.IsEmpty())
	fmt.Println("Try to remove from empty list: ", linkedlist.Remove("first"))
	fmt.Println("-------------")
	linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	fmt.Println("Remove beyond the scope: ", linkedlist.RemoveAt(10))
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	stack := &ds.Stack{}
	data, err := stack.Pop()
	fmt.Println("Pop from empty stack: ", err)
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, _ = stack.Pop()
	println("Popped from stack:", data)

	fmt.Println("-------------")

	queue := &ds.Queue{}
	data, err = queue.Pop()
	fmt.Println("Pop from empty queue: ", err)
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data, _ = queue.Pop()
	println("Popped from queue:", data)
}
