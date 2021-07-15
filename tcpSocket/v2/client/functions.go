package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func login() string {
	fmt.Print("name : ")
	for {
		text := input()
		if len(text) > 10 {
			fmt.Println("less than 10 characters")
		} else if len(text) < 1 {
			fmt.Println("at least one character")
		} else {
			return text
		}
	}
}

func connect() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:9393")
	if nil != err {
		log.Println(err)
	}

	return conn
}

func read(conn net.Conn, res chan []byte) {
	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			log.Println(err)
			return
		}

		res <- data[:n]
	}
}

func reactor(res chan []byte, user *User) {
	for {
		switch data := <-res; string(data[:2]) {
		case "00":
			// get userNo
			user.No = string(data[2:])
		case "01":
			// match found
			found(data[2:])
		case "02":
			// game start
			write("13" + roomNo)
			cls()
		case "03":
			// set player names
			setPlayers(data[2:])
		case "04":
			fmt.Println("04")
		case "05":
			// tic
			update(data[2:])
		case "06":
			fmt.Println("06")
		case "07":
			fmt.Println("07")
		case "08":
			fmt.Println("08")
		case "09":
			fmt.Println("09")
		}
	}
}

func start(conn net.Conn, user User) {
	fmt.Println("Finding opponent")
	time.Sleep(time.Second / 2)
	write("11" + user.No)
}

func found(data []byte) {
	fmt.Println("Match found")
	playerNo = string(data[:1])
	roomNo = string(data[1:])
	time.Sleep(2 * time.Second)
	cls()
	selectChamp()
	inGame()
}

func selectChamp() {
	fmt.Println("Choose your champion")
	for _, v := range champs {
		fmt.Println(v.Name)
	}
	for {
		text := input()
		for _, v := range champs {
			if text == v.Name {
				write("12" + user.No + playerNo + v.No + roomNo)
				return
			}
		}
	}
}

func setPlayers(data []byte) {
	if string(data[:1]) == "1" {
		champNo, _ := strconv.Atoi(string(data[1:3]))
		game.Player1.Champ = champs[champNo]
		game.Player1.Name = string(data[3:])
	} else if string(data[:1]) == "2" {
		champNo, _ := strconv.Atoi(string(data[1:3]))
		game.Player1.Champ = champs[champNo]
		game.Player2.Name = string(data[3:])
	}
}

func update(data []byte) {
	game.Player1.Champ.Hp, _ = strconv.Atoi(string(data[:4]))
	game.Player1.x, _ = strconv.Atoi(string(data[4:8]))
	game.Player1.y, _ = strconv.Atoi(string(data[8:12]))

	game.Player2.Champ.Hp, _ = strconv.Atoi(string(data[12:16]))
	game.Player2.x, _ = strconv.Atoi(string(data[16:20]))
	game.Player2.y, _ = strconv.Atoi(string(data[20:24]))
	fmt.Println(game)
}

func inGame() {
	state = true
	for state {
		text := input()
		click(text)
	}
}

func click(loc string) {
	write("14" + loc)
}
