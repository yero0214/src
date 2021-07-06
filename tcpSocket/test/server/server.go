package main

import (
	"log"
	"net"
)

type User struct {
	conn net.Conn
}

type Rooms struct {
	room []Room
}

type Room struct {
	users []User
}

func main() {
	queue := make(chan User)
	go findRoom(queue)

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
		defer conn.Close()

		queue <- User{conn: conn}
		// go ConnHandler(&queue, conn)
	}
}

// func ConnHandler(list *Queue, conn net.Conn) {
// 	recvBuf := make([]byte, 4096)
// 	for {
// 		n, err := conn.Read(recvBuf)
// 		if nil != err {
// 			if io.EOF == err {
// 				log.Println(err)
// 				return
// 			}
// 			log.Println(err)
// 			return
// 		}
// 		if 0 < n {
// 			data := recvBuf[:n]
// 			log.Println(string(data))
// 			broadCast(list, data[:n])
// 		}
// 	}
// }

// func broadCast(list *Queue, data []byte) {
// 	for _, v := range list.Clients {
// 		v.Write([]byte(data))
// 	}
// }

func findRoom(queue chan User) {
	var rooms Rooms
	for {
		user := <-queue
		if len(rooms.room) == 0 {
			room := Room{}
			rooms.room = append(rooms.room, room)
		}
		if len(rooms.room[len(rooms.room)-1].users) < 2 {
			rooms.room[len(rooms.room)-1].users = append(rooms.room[len(rooms.room)-1].users, user)
		} else {
			room := Room{}
			rooms.room = append(rooms.room, room)
			rooms.room[len(rooms.room)-1].users = append(rooms.room[len(rooms.room)-1].users, user)
		}
	}
}
