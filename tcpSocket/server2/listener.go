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

		x := uint32(randomNum(10, 90))
		y := uint32(randomNum(10, 90))

		users = append(users, User{Conn: conn, userId: count, x: x, y: y, cx: x, cy: y})
		count++

		go ConnHandler(conn)
	}
}
