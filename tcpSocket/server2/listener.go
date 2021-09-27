package main

import (
	"log"
	"net"
)

func startListen(port string) {
	l, err := net.Listen("tcp", ":"+port)
	if nil != err {
		log.Println(err)
	}
	defer l.Close()

	log.Println("listening...")

	for {
		conn, err := l.Accept()
		log.Println(conn)
		if nil != err {
			log.Println(err)
			continue
		}

		// x := float32(randomNum(10, 90))
		// y := float32(randomNum(10, 90))
		x := float32(50)
		y := float32(50)
		users = append(users, User{Conn: conn, userId: count, x: x, y: y, cx: x, cy: y})

		go ConnHandler(conn, count)
		count++
	}
}
