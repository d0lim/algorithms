package datastructure

import "fmt"

func Example_queueLinkedList() {
	q := NewQueueLinkedList()
	fmt.Println(q.Size())
	fmt.Println(q.IsEmpty())
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)
	fmt.Println(q.Size())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	for k := range q.Iterate() {
		fmt.Println(k)
	}
	// Output:
	// 0
	// true
	// 4
	// false
	// 3 <nil>
	// 5 <nil>
	// 7
	// 9
}