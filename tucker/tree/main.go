package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1()
	fmt.Println()
	tree.DFS2()
	fmt.Println()
	tree.BFS()
	fmt.Println()

	binaryTree := dataStruct.NewBinaryTree(5)
	binaryTree.Root.AddNode(3)
	binaryTree.Root.AddNode(2)
	binaryTree.Root.AddNode(4)
	binaryTree.Root.AddNode(8)
	binaryTree.Root.AddNode(7)
	binaryTree.Root.AddNode(6)
	binaryTree.Root.AddNode(10)
	binaryTree.Root.AddNode(9)

	binaryTree.Print()

	fmt.Println()

	if found, cnt := binaryTree.Search(3); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("not found 6 cnt:", cnt)
	}

	h := &dataStruct.Heap{}

	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 4 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())

	h.Push(-1)
	h.Push(3)
	h.Push(-1)
	h.Push(5)
	h.Push(4)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
