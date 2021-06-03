package main

import "fmt"

func main() {
	i := 0
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("Final i value: ", i)
}
