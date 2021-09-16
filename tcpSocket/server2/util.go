package main

import (
	"math/rand"
	"time"
)

func broadCast(buffer []byte) {
	for _, v := range users {
		v.Conn.Write(buffer)
	}
}

func randomNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
