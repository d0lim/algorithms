package datastructure

import "errors"

// NewTree returns new tree
func NewTree(rootItem interface{}) Tree {
	newNode := TreeNode{
		Item:  rootItem,
		Level: 0,
	}
	return Tree{
		root: &newNode,
	}
}

// Tree is struct that implements datastructure tree
type Tree struct {
	root *TreeNode
}

// removeSliceIndex removes item at index in slice
func removeSliceIndex(s []*TreeNode, index int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// AddChild adds child to parent.
func (t *TreeNode) AddChild(item interface{}) {
	newChild := &TreeNode{}
	newChild.Item = item
	newChild.Level = t.Level + 1
	t.Children = append(t.Children, newChild)
}

// RemoveChild removes child from parent
func (t *TreeNode) RemoveChild(item interface{}) error {
	success := false
	for index, child := range t.Children {
		if item == child.Item {
			t.Children = removeSliceIndex(t.Children, index)
			success = true
		}
	}
	if !success {
		return errors.New("Tree RemoveChild: item not found")
	}
	return nil
}
