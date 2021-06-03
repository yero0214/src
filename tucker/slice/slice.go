package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[4:8]
	c := a[4:]
	d := a[:4]

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
