package main

import (
	"log"
	"net"
)

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			// remove disconnected user
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
			// data := recvBuf[:n]
			// log.Println(string(data))
			// log.Println(data)
			log.Println(n)
			recvChan <- recvBuf[:n]
		}
	}
}
