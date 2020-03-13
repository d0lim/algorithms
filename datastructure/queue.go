package datastructure

// Queue is an interface which implements data structure queue.
type Queue interface {
	Enqueue(obj interface{})
	Dequeue() (interface{}, error)
	IsEmpty() bool
	Size() int
	Iterate() <-chan interface{}
}