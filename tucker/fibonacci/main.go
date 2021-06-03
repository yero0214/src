package main

import "fmt"

func main() {
	rst := f(10)
	fmt.Println(rst)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
