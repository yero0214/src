package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Make 3 random numbers
	numbers := MakeNumbers()

	cnt := 1
	for {
		cnt++
		// Get player input
		inputNumbers := InputNumbers()

		// compare result
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S check
		if IsGameEnd(result) {
			// end
			break
		}
	}

	// print how long it took
	fmt.Printf("You tried %d times,\n", cnt)
}

func MakeNumbers() [3]int {
	// make three different numbers between 0 ~ 9
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// get 3 numbers between 0 ~ 9 and return
	var rst [3]int

	for {
		fmt.Println("enter 3 different numbers between 0 ~  9")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("try again")
			continue
		}
		fmt.Println(no)
		success := true

		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("can't use multiple same numbers")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("only 3 numbers allowed")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if success && idx < 3 {
			fmt.Println("enter 3 numbers")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]

	fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// compare numbers and return result
	strikes := 0
	balls := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}
	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS %dB\n", result.strikes, result.balls)
}

func IsGameEnd(result Result) bool {
	//check is the result equals 3S
	return result.strikes == 3
}
