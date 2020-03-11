package datastructure

import "errors"

type stackLinkedList struct {
	Current *node
	Length  int
}

type node struct {
	Item interface{}
	Next *node
}

// NewStackLinkedList returns stack made with linked list.
func NewStackLinkedList() Stack {
	return &stackLinkedList{}
}

func (s *stackLinkedList) Push(obj interface{}) {
	newNode := &node{}
	newNode.Item = obj
	newNode.Next = s.Current
	s.Current = newNode
	s.Length++
}

func (s *stackLinkedList) Pop() (interface{}, error) {
	if !s.IsEmpty() {
		item := s.Current.Item
		s.Current = s.Current.Next
		s.Length--
		return item, nil
	}
	return nil, errors.New("stack linkedlist pop: stack is empty")
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.Length == 0
}

func (s *stackLinkedList) Size() int {
	return s.Length
}

func (s *stackLinkedList) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			if s.IsEmpty() {
				break
			}
			p, err := s.Pop()
			if err != nil {
				panic(err)
			}
			ch <- p
		}
	}()
	return ch
}