package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	stack := []int{}
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}
}
