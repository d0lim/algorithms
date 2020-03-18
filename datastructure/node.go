package datastructure

// LinkedListNode is  astruct for common data structure. It can be used in stack, queue and whatever.
type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

type level int

// TreeNode is a struct for tree data structure.
type TreeNode struct {
	Item     interface{}
	Level    level
	Children []*TreeNode
}
