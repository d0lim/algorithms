package datastructure

import "fmt"

func Example_treeRemove() {
	tree := NewTree("root")
	tree.root.AddChild("First")
	tree.root.AddChild("Second")
	fmt.Println(tree.root.Children[0].Item)
	fmt.Println(tree.root.Children[1].Item)
	err := tree.root.RemoveChild("notExist")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// First
	// Second
	// Tree RemoveChild: item not found
}
