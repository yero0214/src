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
	var name string
	state := false
	fmt.Println("Enter your name")
	for {
		text := read()
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
					break
				}
				res := data[:n]

				//match found
				if string(res[:2]) == "01" {
					roomNo = string(res[2:4])
					conn.Write([]byte("11" + roomNo + name))
					fmt.Println("Match found!")
					state = true
				}
				fmt.Println(string(res))
			}
		}()

		for {
			s := read()
			if !state {
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
