package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 4; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i <= 4; i++ {
		for j := 0; j <= i; j++ {
			if i == 3 && j == 2 {
				break
			}
			if i == 4 && j == 1 {
				break
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 3-i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
