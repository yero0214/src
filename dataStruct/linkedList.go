package dataStruct

import "fmt"

type Node struct {
	Next *Node
	prev *Node
	Val  int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(Val int) {
	if l.Root == nil {
		l.Root = &Node{Val: Val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: Val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.prev = prev
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *LinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

func (l *LinkedList) Empty() bool {
	return l.Root == nil
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.prev = nil
		}
		node.Next = nil
		return
	}

	prev := node.prev

	if node == l.Tail {
		prev.Next = nil
		l.Tail.prev = nil
		l.Tail = prev
	} else {
		node.prev = nil
		prev.Next = prev.Next.Next
		prev.Next.prev = prev
	}
	node.Next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail
	for node.prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.Val)
}
