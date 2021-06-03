package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	// }

	// a, b := fun1(2, 3)

	// fmt.Println(a, b)

	s := sum(10, 0)
	fmt.Println(s)
}

func add(x int, y int) int {
	return x + y
}

func fun1(x, y int) (int, int) {
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println("func2", x, y)
}
func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}
