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

		go ConnHandler(conn)
	}
}
