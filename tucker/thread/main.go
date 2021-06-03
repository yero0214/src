package main

import (
	"fmt"
	"time"
)

func main() {
	go fun1()
	go fun1()
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main", i)
	}
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("func1:", i)
	}
}
