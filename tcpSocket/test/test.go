package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var slice []byte
	// var buffer [4096]byte
	// slice[1] = buffer[100]
	// fmt.Printf("%s", slice[0:100])
	var test string
	test = "01"
	num, _ := strconv.Atoi(test)
	fmt.Println(num)
}
