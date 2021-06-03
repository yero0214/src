package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter numbers")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("You entered %d, %d \n", n1, n2)

	fmt.Println("Enter a operator")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	}

	// if line == "+" {
	// 	fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	// } else if line == "*" {
	// 	fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	// } else if line == "/" {
	// 	fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	// } else {
	// 	fmt.Println("Wrong input")
	// }
}
