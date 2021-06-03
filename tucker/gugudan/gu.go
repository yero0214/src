package main

import "fmt"

func main() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("\n%d dan \n", dan)

		for i := 1; i <= 9; i++ {
			if dan == 3 && i == 2 {
				continue
			}
			fmt.Printf("%d * %d = %d \n", dan, i, dan*i)
		}
	}
}
