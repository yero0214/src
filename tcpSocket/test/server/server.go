package main

import (
	"fmt"
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

		queue <- User{conn: conn, health: 100, name: "test"}
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

		if 0 < n {
			data := recvBuf[:n]
			fmt.Println(string(data[:n]))
			if string(data[0:2]) == "10" {
				roomNo, _ := strconv.Atoi(string(data[2:4]))
				if string(data[4:6]) == "01" {
					var str string
					for index, _ := range rooms.room[roomNo].users {
						if rooms.room[roomNo].users[index].conn != conn {
							rooms.room[roomNo].users[index].health -= 10
						}
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))

				} else if string(data[4:6]) == "02" {
					var str string
					for index, _ := range rooms.room[roomNo].users {
						if rooms.room[roomNo].users[index].conn == conn {
							rooms.room[roomNo].users[index].health += 10
						}
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))
				}
			}
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

func matchStart(room Room, intRoomNo int) {
	var roomNo string
	if intRoomNo < 10 {
		roomNo = "0" + strconv.Itoa(intRoomNo)
	} else {
		roomNo = strconv.Itoa(intRoomNo)
	}

	broadCast(&room, []byte("01"+roomNo))
}
