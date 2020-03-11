package datastructure

// Stack includes general function of stack.
type Stack interface {
	Push(item interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
	Size() int
	Iterate() <-chan interface{}
}