package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	queue := []int{}
	for i := 1; i < 6; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	queue2 := dataStruct.NewQueue()
	for i := 1; i < 6; i++ {
		queue2.Push(i)
	}

	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}
}
