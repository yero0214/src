package main

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

func write(conn net.Conn, content string) {
	conn.Write([]byte(content))
}

func broadCast(data string) {
	for _, v := range users {
		write(v.Conn, data)
	}
}

func randomNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func makeNo(num int) string {
	if num < 10 {
		return "000" + strconv.Itoa(num)
	} else if num < 100 {
		return "00" + strconv.Itoa(num)
	} else if num < 1000 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}
}
