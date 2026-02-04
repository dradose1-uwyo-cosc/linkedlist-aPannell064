//[Your Name Here]
//COSC 3750
//[Date Here]
//
/*
	Don't forget to run your go mod init command in your terminal
	Review the assignment instructions for running your code
	All the code you need to write should be put in the /ds/ package files
	Uncomment the import statement for your module name
	you can uncomment the tests in main as you go to test
	The code in main is not an extensive test, you should add more and test your code as needed
*/
package main

import (
	"fmt"
	"hw-linked/ds"
)

func main() {
	//fmt.Println("Only here so the import doesn't leave an error")

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
