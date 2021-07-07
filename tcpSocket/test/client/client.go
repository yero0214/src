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
	var roomNo string
	fmt.Println("Type 'start' to start match making")
	for {
		text := read()
		if text == "start" {
			break
		}
	}

	conn, err := net.Dial("tcp", ":9393")
	if nil != err {
		log.Println(err)
	}

	fmt.Println("Finding Opponent...")
	defer conn.Close()

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}
			res := data[:n]

			//match found
			if string(res[:2]) == "01" {
				roomNo = string(res[2:4])
				log.Println(roomNo)
			}
			fmt.Println(string(res))
		}
	}()

	for {
		s := read()
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

func read() string {
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
