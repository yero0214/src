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
	var rooms *Rooms
	queue := make(chan User)
	go findRoom(queue, rooms)

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

func findRoom(queue chan User, rooms *Rooms) {
	for {
		user := <-queue
		lastRoom := rooms.room[len(rooms.room)-1]
		if len(rooms.room) == 0 {
			rooms.room = append(rooms.room)
		}
		if len(lastRoom.users) < 2 {
			lastRoom.users = append(lastRoom.users, user)
		} else {
			rooms.room = append(rooms.room)
			lastRoom.users = append(lastRoom.users, user)
		}
		log.Println(lastRoom)
	}
}
