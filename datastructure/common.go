package datastructure

// Node is struct for common data structure. It can be used in stack, queue and whatever.
type Node struct {
	Item interface{}
	Next *Node
}