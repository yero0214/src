package main

import (
	"fmt"
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
	cx   int
	cy   int
}

var users []User
var count int

func main() {

	go inGame()
	go yUpdate()
	go xUpdate()

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
		x := randomNum()
		y := randomNum()

		users = append(users, User{Conn: conn, no: count, x: x, y: y, cx: x, cy: y})
		count++
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			for i, _ := range users {
				if users[i].Conn == conn {
					users[i] = users[len(users)-1]
					users[len(users)-1] = User{}
					users = users[:len(users)-1]
					break
				}
			}
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			fmt.Println(string(data))
			for i, v := range users {
				if conn == v.Conn {
					users[i].cx, _ = strconv.Atoi(string(data[:4]))
					users[i].cy, _ = strconv.Atoi(string(data[4:8]))
				}
			}
		}
	}
}

func inGame() {
	for {
		time.Sleep(time.Second * 5)
		var result string
		for _, v := range users {
			result += makeNo(v.no)
			result += makeNo(v.x)
			result += makeNo(v.y)
		}

		broadCast(result)
	}
}

func xUpdate() {
	for {
		time.Sleep(time.Second)
		for i, _ := range users {
			if users[i].cx > users[i].x {
				users[i].x++
			} else if users[i].cx < users[i].x {
				users[i].x--
			} else {
				continue
			}
		}
	}
}

func yUpdate() {
	for {
		time.Sleep(time.Second)
		for i, _ := range users {
			if users[i].cy > users[i].y {
				users[i].y++
			} else if users[i].cy < users[i].y {
				users[i].y--
			} else {
				continue
			}
		}
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
