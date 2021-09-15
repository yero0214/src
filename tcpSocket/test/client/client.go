package main

import (
	"bufio"
	"encoding/binary"
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
			payload := input()
			typeOfService := make([]byte, 4)
			payloadLength := make([]byte, 4)
			binary.LittleEndian.PutUint32(typeOfService, 1)
			binary.LittleEndian.PutUint32(payloadLength, 4)

			packetHeader := append(typeOfService, payloadLength...)
			buffer := append(packetHeader, []byte(payload)...)
			log.Println(buffer)
			conn.Write(buffer)
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
