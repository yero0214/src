package main

import (
	"bufio"
	"fmt"
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
	conn := connect()
	defer conn.Close()
	fmt.Println("connected")

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		// if text == "attack" {
		// attack(conn)
		// rec := receive(conn)
		// fmt.Println(rec)
		// } else if text == "heal" {
		// 	heal(conn)
		// 	rec := receive(conn)
		// 	fmt.Println(rec)
		// }

	}
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func connect() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:9393")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}

func attack(conn net.Conn) {
	conn.Write([]byte("attack"))
}

func heal(conn net.Conn) {
	conn.Write([]byte("heal"))
}

func receive(conn net.Conn) string {
	buf := make([]byte, 0, 4096)
	message, _ := conn.Read(buf)
	return string(message)
}
