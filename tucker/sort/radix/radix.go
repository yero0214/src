package main

import "fmt"

func main() {
	arr := [12]int{0, 3, 7, 3, 7, 6, 2, 1, 9, 1, 4, 6}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}
	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}
