package datastructure

// LinkedListNode is  astruct for common data structure. It can be used in stack, queue and whatever.
type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

// TreeNode is a struct for tree data structure.
type TreeNode struct {
	Item     interface{}
	Children []TreeNode
}
