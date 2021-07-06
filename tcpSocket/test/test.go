package main

import "fmt"

func main() {
	var slice []byte
	var buffer [4096]byte
	slice[1] = buffer[100]
	fmt.Printf("%s", slice[0:100])
}
