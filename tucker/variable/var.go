package main

import "fmt"

func main() {
	//First bit represents - +
	//Variable types
	// int8 == 8 bit 1 byte
	// can represent 256 numbers (-128 ~ 127)
	// uint8 0 ~ 255

	// int16 == 16 bit 2 byte
	// can represent 65535 numbers (-32768 ~ 32767)
	// uint16 0 ~ 65535

	// int32 == 32 bit 3 byte
	// can represent 256 numbers (-2147483648 ~ 2147483647)
	// uint32 0 ~ 4294967295

	// int64 == 64 bit 4 byte
	// can represent 256 numbers (-9223372036854775808 ~ 9223372036854775807)
	// uint64 0 ~ 18446744073709551615

	// int == 64 or 32 bit depends on op
	// uint 64 or 32 bit depends on op

	// float32 == 32 bit 4 byte
	// can represent 7 digits ex) 3.141516 (3141516 * 10^-6)

	// float64 == 64 bit 8 byte
	// can represent 15 digits

	// float == 64 or 32 bit depends on op

	// String == size change depends on value
	// golang uses utf-8 (each word uses 1 ~ 3 byte)

	// Range test

	// max := uint8(0)
	// for i := uint8(0); i >= 0; i++ {
	// 	max = i
	// }
	// fmt.Println("Max:", max)

	// min := uint8(0)
	// for i := uint8(0); i <= 0; i-- {
	// 	min = i
	// }
	// fmt.Println("Min:", min)

	//Ways to declarate variables
	a := 1
	b := int(2)
	var c = 3
	var d int = 4
	var e int
	e = 5

	fmt.Printf("a: %v b: %v c: %v d: %v e: %v \n", a, b, c, d, e)

}
