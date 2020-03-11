package datastructure

import "fmt"

func Example_stackArray() {
	fmt.Println("Hello")
	// Output:
	// Hello
}

func Example_stackLinkedList() {
	s := NewStackLinkedList()
	fmt.Println(s.IsEmpty())
	s.Push(5)
	s.Push(3)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	// Output:
	// true
	// false
	// 2
	// 3 <nil>
	// 1
	// 5 <nil>
}