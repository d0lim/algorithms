package datastructure

import "errors"

type queueLinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	Length int
}

// NewQueueLinkedList returns queue datastructure based on linked list.
func NewQueueLinkedList() Queue {
	return &queueLinkedList{}
}

func (q *queueLinkedList) Enqueue(value interface{}) {
	oldTail := q.Tail
	q.Tail = &LinkedListNode{}
	q.Tail.Item = value

	if q.IsEmpty() {
		q.Head = q.Tail
	} else {
		oldTail.Next = q.Tail
	}
	q.Length++
}

func (q *queueLinkedList) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue linked list pop: queue is empty")
	}
	item := q.Head.Item
	q.Head = q.Head.Next
	q.Length--
	if q.Length == 0 {
		q.Tail = q.Head
	}
	return item, nil
}

func (q *queueLinkedList) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queueLinkedList) Size() int {
	return q.Length
}

func (q *queueLinkedList) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			if q.IsEmpty() {
				break
			}
			d, err := q.Dequeue()
			if err != nil {
				panic(err)
			}
			ch <- d
		}
	}()
	return ch
}
