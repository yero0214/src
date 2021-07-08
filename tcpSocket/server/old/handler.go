package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func ConnHandler(conn net.Conn, rooms *Rooms) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			log.Println(err)
			for i, _ := range rooms.room {
				for j, _ := range rooms.room[i].users {
					if rooms.room[i].users[j].conn == conn {
						rooms.room[i].users[j].health = 0
						gameState(&rooms.room[i])
					}
				}
			}
			return
		}
		defer conn.Close()
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
				var str string
				//if have both name
				var users int
				for _, v := range rooms.room[roomNo].users {
					if v.name != "" {
						users++
					}
				}
				if users == 2 {
					for index, _ := range rooms.room[roomNo].users {
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))
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
					gameState(&rooms.room[roomNo])

				} else if string(data[4:6]) == "02" {
					var str string
					for index, _ := range rooms.room[roomNo].users {
						if rooms.room[roomNo].users[index].conn == conn {
							rooms.room[roomNo].users[index].health += 10
						}
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))
					gameState(&rooms.room[roomNo])

				} else if string(data[4:6]) == "00" {
					var str string
					for index, _ := range rooms.room[roomNo].users {
						if rooms.room[roomNo].users[index].conn != conn {
							rooms.room[roomNo].users[index].health = 0
						}
						str += " | " + rooms.room[roomNo].users[index].name + " " + strconv.Itoa(rooms.room[roomNo].users[index].health) + " | "
					}
					broadCast(&rooms.room[roomNo], []byte(str))
					gameState(&rooms.room[roomNo])
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

func gameState(room *Room) {
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

func matchStart(room Room, intRoomNo int) {
	var roomNo string
	if intRoomNo < 10 {
		roomNo = "0" + strconv.Itoa(intRoomNo)
	} else {
		roomNo = strconv.Itoa(intRoomNo)
	}

	broadCast(&room, []byte("01"+roomNo))
}
