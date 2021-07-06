package main

import (
	"io"
	"log"
	"net"
	"strconv"
)

type User struct {
	conn   net.Conn
	health int
	name   string
}

type Rooms struct {
	room []Room
}

type Room struct {
	users []User
}

func main() {
	var rooms Rooms
	queue := make(chan User)
	go findRoom(queue, &rooms)

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
		go ConnHandler(conn, &rooms)
	}
}

func ConnHandler(conn net.Conn, rooms *Rooms) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}

		room := Room{}
		log.Println(len(rooms.room))
		// for i := range rooms.room {
		// 	for j := range rooms.room[i].users {
		// 		if rooms.room[i].users[j].conn == conn {
		// 			room = rooms.room[i]
		// 		}
		// 	}
		// }

		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			broadCast(&room, data[:n])
		}
	}
}

func broadCast(room *Room, data []byte) {
	for _, v := range room.users {
		v.conn.Write(data)
	}
}

func findRoom(queue chan User, rooms *Rooms) {
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
		if len(rooms.room[len(rooms.room)-1].users) == 2 {
			matchStart(rooms.room[len(rooms.room)-1], len(rooms.room)-1)
		}
	}
}

func matchStart(room Room, roomNo int) {
	len := len([]byte(strconv.Itoa(roomNo)))
	var blen string
	if len < 10 {
		blen = "0" + strconv.Itoa(len)
	} else {
		blen = strconv.Itoa(len)
	}
	broadCast(&room, []byte("01"+blen+strconv.Itoa(roomNo)))
}
