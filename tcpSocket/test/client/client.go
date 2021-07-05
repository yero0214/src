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
	fmt.Println("Type 'connect'")
	for {
		text := read()
		if text == "connect" {
			break
		}
	}

	conn, err := net.Dial("tcp", ":9393")
	if nil != err {
		log.Println(err)
	}

	fmt.Println("connected")

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
		}
	}()

	for {
		s := read()
		if s == "attack" {
			attack(conn)
		} else if s == "heal" {
			heal(conn)
		}
	}
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func attack(conn net.Conn) {
	conn.Write([]byte("attack"))
}

func heal(conn net.Conn) {
	conn.Write([]byte("heal"))
}
