package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type Champ struct {
	No   string
	Name string
	Hp   int
	Atk  int
}

type User struct {
	Conn net.Conn
	Name string
	No   string
}

type Res struct {
	Conn net.Conn
	Data []byte
}

type Room struct {
	Users []User
	Ready int
	Game  Game
}

type Game struct {
	Player1 Player
	Player2 Player
	State   bool
}

type Player struct {
	Name  string
	Champ Champ
	x     int
	y     int
}

var champs []Champ
var rooms []Room
var users []User

func main() {
	champs = append(champs, Champ{No: "00", Name: "Ashe", Hp: 100, Atk: 10})
	champs = append(champs, Champ{No: "01", Name: "MG", Hp: 1000, Atk: 1})

	resChan := make(chan Res)
	userChan := make(chan User)
	queueChan := make(chan int)
	locChan := make(chan []byte)

	go reactor(resChan, userChan, queueChan, locChan)
	go addUser(&users, userChan)
	go queue(&users, &rooms, queueChan)

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
		go ConnHandler(conn, resChan)
	}
}

func ConnHandler(conn net.Conn, resChan chan Res) {

	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			log.Println(err)
			return
		}
		if 0 < n {
			resChan <- Res{Conn: conn, Data: recvBuf[:n]}
		}
	}
}

func reactor(resChan chan Res, userChan chan User, queueChan chan int, locChan chan []byte) {
	for {
		switch val := <-resChan; string(val.Data[:2]) {
		case "10":
			// add user to user array
			fmt.Println("10")
			userChan <- User{Conn: val.Conn, Name: string(val.Data[2:])}
		case "11":
			// add userNo to queue
			fmt.Println("11")
			no, _ := strconv.Atoi(string(val.Data[2:]))
			queueChan <- no
		case "12":
			//champ selected
			fmt.Println("12")
			champSelected(val.Data[2:])
		case "13":
			// get ready sign
			fmt.Println("13")
			readyChk(val.Data[2:])
		case "14":
			// update the game
			locChan <- val.Data[2:]
		case "15":
			fmt.Println("15")
		case "16":
			fmt.Println("16")
		case "17":
			fmt.Println("17")
		case "18":
			fmt.Println("18")
		case "19":
			fmt.Println("19")
		}
	}
}

func addUser(users *[]User, userChan chan User) {
	for {
		user := <-userChan
		fmt.Println("addUser: " + user.Name)
		*users = append(*users, user)
		write(user.Conn, "00"+makeNo(len(*users)-1))
	}
}

func queue(users *[]User, rooms *[]Room, queueChan chan int) {
	for {
		no := <-queueChan
		if len(*rooms) == 0 {
			room := Room{}
			*rooms = append(*rooms, room)
		}
		if len((*rooms)[len(*rooms)-1].Users) < 2 {
			(*rooms)[len(*rooms)-1].Users = append((*rooms)[len(*rooms)-1].Users, (*users)[no])
		} else {
			room := Room{}
			*rooms = append(*rooms, room)
			(*rooms)[len(*rooms)-1].Users = append((*rooms)[len(*rooms)-1].Users, (*users)[no])
		}
		if len((*rooms)[len(*rooms)-1].Users) == 2 {
			matchStart((*rooms)[len(*rooms)-1], len(*rooms)-1)
		}
	}
}

func matchStart(room Room, intRoomNo int) {
	roomNo := strconv.Itoa(intRoomNo)
	for i, v := range room.Users {
		write(v.Conn, "01"+strconv.Itoa(i)+roomNo)
	}
}

func champSelected(data []byte) {
	userNo, _ := strconv.Atoi(string(data[:4]))
	playerNo, _ := strconv.Atoi(string(data[4:5]))
	champNo, _ := strconv.Atoi(string(data[5:7]))
	roomNo, _ := strconv.Atoi(string(data[7:]))
	if playerNo == 0 {
		rooms[roomNo].Game.Player1.Champ = champs[champNo]
		rooms[roomNo].Game.Player1.Name = users[userNo].Name
		rooms[roomNo].Game.Player1.x = 0
		rooms[roomNo].Game.Player1.y = 0
	} else if playerNo == 1 {
		rooms[roomNo].Game.Player2.Champ = champs[champNo]
		rooms[roomNo].Game.Player2.Name = users[userNo].Name
		rooms[roomNo].Game.Player2.x = 100
		rooms[roomNo].Game.Player2.y = 100
	}

	var ready int

	if rooms[roomNo].Game.Player1.Champ != (Champ{}) {
		ready++
	}
	if rooms[roomNo].Game.Player2.Champ != (Champ{}) {
		ready++
	}

	if ready == 2 {
		broadCast(&rooms[roomNo], "02")
	}

}

func readyChk(data []byte) {
	roomNo, _ := strconv.Atoi(string(data))
	rooms[roomNo].Ready++
	if rooms[roomNo].Ready == 2 {
		p1Name := rooms[roomNo].Game.Player1.Name
		p2Name := rooms[roomNo].Game.Player2.Name
		p1ChampNo := rooms[roomNo].Game.Player1.Champ.No
		p2ChampNo := rooms[roomNo].Game.Player2.Champ.No
		broadCast(&rooms[roomNo], "03"+"1"+p1ChampNo+p1Name)
		broadCast(&rooms[roomNo], "03"+"2"+p2ChampNo+p2Name)
		time.Sleep(2 * time.Second)
		go inGame(&rooms[roomNo])
	}
}

func inGame(room *Room) {
	room.Game.State = true
	for room.Game.State {
		time.Sleep(1000 * time.Microsecond)
		p1Hp := makeNo(room.Game.Player1.Champ.Hp)
		p1x := makeNo(room.Game.Player1.x)
		p1y := makeNo(room.Game.Player1.x)
		p2Hp := makeNo(room.Game.Player2.Champ.Hp)
		p2x := makeNo(room.Game.Player2.x)
		p2y := makeNo(room.Game.Player2.x)
		broadCast(room, "05"+p1Hp+p1x+p1y+p2Hp+p2x+p2y)
	}
}

func write(conn net.Conn, content string) {
	conn.Write([]byte(content))
}

func broadCast(room *Room, data string) {
	for _, v := range room.Users {
		write(v.Conn, data)
	}
}

func makeNo(num int) string {
	if num < 10 {
		return "000" + strconv.Itoa(num)
	} else if num < 100 {
		return "00" + strconv.Itoa(num)
	} else if num < 1000 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}
}

func locationChange(locChan chan []byte) {
	// loc := <-locChan
	// for{

	// }
}
