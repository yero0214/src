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
	for {
		conn, err := net.Dial("tcp", ":9393")
		if nil != err {
			log.Println(err)
		}
		defer conn.Close()

		go read(conn)

		for {
			conn.Write([]byte(input()))
		}
	}
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func read(conn net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			break
		}
		res := data[:n]
		fmt.Println(string(res))
	}
}
