package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

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

		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			log.Println(err)
		}
		if 0 < n {
			data := recvBuf[:n]
			fmt.Println(string(data[:n]))
			conn.Write(data)
		}
	}
}
