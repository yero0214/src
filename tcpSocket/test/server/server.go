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

		queue <- User{conn: conn, health: 100}
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
			if string(data[:2]) == "11" {
				roomNo, _ := strconv.Atoi(string(data[2:4]))
				for index, _ := range rooms.room[roomNo].users {
					if rooms.room[roomNo].users[index].conn == conn {
						rooms.room[roomNo].users[index].name = string(data[4:])
					}
				}

			} else if string(data[:2]) == "10" {
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
					gameStatus(&rooms.room[roomNo])

				} else if string(data[4:6]) == "02" {
					var str string
					for index, _ := range rooms.room[roomNo].users {
						if rooms.room[roomNo].users[index].conn == conn {
							rooms.room[roomNo].users[index].health += 10
						}
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))
					gameStatus(&rooms.room[roomNo])
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

func gameStatus(room *Room) {
	for _, v := range room.users {
		if v.health <= 0 {
			gameEnd(room)
			break
		}
	}
}

func gameEnd(room *Room) {
	for _, v := range room.users {
		if v.health <= 0 {
			v.conn.Write([]byte("Defeat"))
			v.conn.Close()
		} else {
			v.conn.Write([]byte("Victory"))
			v.conn.Close()
		}
	}
	room = nil
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
