package main

import "fmt"

func main() {
	// fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	// fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	// fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	// fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))

	// m := dataStruct.CreateMap()
	// m.Add("AAA", "01077777777")
	// m.Add("BBB", "01088888888")
	// m.Add("CDEFRGTEFVDF", "0111111111")
	// m.Add("CCC", "017575757575")

	// fmt.Println("AAA = ", m.Get("AAA"))
	// fmt.Println("CCC = ", m.Get("CCC"))
	// fmt.Println("DDD = ", m.Get("DDD"))
	// fmt.Println("CDEFRGTEFVDF = ", m.Get("CDEFRGTEFVDF"))

	var m map[string]string
	m = make(map[string]string)
	m["abc"] = "bbb"

	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[55])

	v, ok := m1[53]
	fmt.Println(v, ok)

	delete(m1, 53)
	v, ok = m1[53]
	fmt.Println(v, ok)
	m2 := make(map[int]int)
	m2[2] = 99
	m2[56] = 9090
	m2[33] = 34332

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}
}
