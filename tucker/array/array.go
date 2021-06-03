package main

import "fmt"

func main() {
	// var A [10]int
	// for i := 0; i < len(A); i++ {
	// 	A[i] = i * i
	// }
	// fmt.Println(A)

	// s := "Hello 월드"
	// s2 := []rune(s)
	// fmt.Println("len(s2) = ", len(s2))
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Print(s2[i], ", ")
	// }

	// arr := [5]int{1, 2, 3, 4, 5}
	// temp := [5]int{}
	// for i := 0; i < 5; i++ {
	// 	temp[i] = arr[len(arr)-1-i]
	// }
	// for i := 0; i < 5; i++ {
	// 	arr[i] = temp[i]
	// }
	// fmt.Println("arr: ", arr)
	// fmt.Println("temp: ", temp)

	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(arr)

}
