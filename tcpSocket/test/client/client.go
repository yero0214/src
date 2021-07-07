package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	var name string
	var roomNo string
	var state string
	fmt.Println("Enter your name")
	for {
		text := input()
		if len(text) > 10 {
			fmt.Println("less than 10 characters")
		} else if len(text) < 1 {
			fmt.Println("at least one character")
		} else {
			name = text
			break
		}
	}
	for {
		fmt.Println("Type 'start' to start match making")
		for {
			text := input()
			if text == "start" {
				break
			}
		}

		conn, err := net.Dial("tcp", ":9393")
		if nil != err {
			log.Println(err)
		}
		state = "start"
		fmt.Println("Finding Opponent...")
		defer conn.Close()

		go read(conn, name, &roomNo, &state)

		for {
			s := input()
			if state == "end" {
				break
			}
			if roomNo == "" {
				continue
			}
			if s == "attack" {
				attack(conn, roomNo)
			} else if s == "heal" {
				heal(conn, roomNo)
			}
		}
	}
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func attack(conn net.Conn, roomNo string) {
	conn.Write([]byte("10" + roomNo + "01"))
}

func heal(conn net.Conn, roomNo string) {
	conn.Write([]byte("10" + roomNo + "02"))
}

func read(conn net.Conn, name string, roomNo *string, state *string) {
	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			*roomNo = ""
			*state = "end"
			fmt.Println("Press enter to quit")
			break
		}
		res := data[:n]

		//match found
		if string(res[:2]) == "01" {
			*roomNo = string(res[2:4])
			conn.Write([]byte("11" + *roomNo + name))
			fmt.Println("Match found!")
		} else {
			fmt.Println(string(res))
		}
	}
}
