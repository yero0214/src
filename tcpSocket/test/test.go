package main

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type User struct {
	Conn net.Conn
	no   int
	x    int
	y    int
}

var users []User
var count int

func main() {

	go inGame()

	l, err := net.Listen("tcp", ":9393")
	if nil != err {
		log.Println(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			continue
		}
		users = append(users, User{Conn: conn, no: count, x: randomNum(), y: randomNum()})
		count++
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			for _, v := range users {
				if conn == v.Conn {
					v.x, _ = strconv.Atoi(string(data[:4]))
					v.y, _ = strconv.Atoi(string(data[4:8]))
				}
			}
		}
	}
}

func inGame() {
	for {
		time.Sleep(time.Second)
		var result string
		for _, v := range users {
			result += makeNo(v.no)
			result += makeNo(v.x)
			result += makeNo(v.y)
		}

		broadCast(result)
	}
}

func write(conn net.Conn, content string) {
	conn.Write([]byte(content))
}

func broadCast(data string) {
	for _, v := range users {
		write(v.Conn, data)
	}
}

func randomNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 90
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
