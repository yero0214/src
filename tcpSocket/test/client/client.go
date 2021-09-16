package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"net"
	"os"
	"strconv"
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
			input := input()
			slice := strings.Split(input, " ")
			x := make([]byte, 4)
			y := make([]byte, 4)
			ux, _ := strconv.ParseInt(slice[0], 10, 64)
			uy, _ := strconv.ParseUint(slice[1], 10, 64)
			binary.LittleEndian.PutUint32(x, uint32(ux))
			binary.LittleEndian.PutUint32(y, uint32(uy))
			payload := append(x, y...)

			typeOfService := make([]byte, 4)
			userId := make([]byte, 4)
			payloadLength := make([]byte, 4)
			binary.LittleEndian.PutUint32(typeOfService, 1)
			binary.LittleEndian.PutUint32(userId, 0)
			binary.LittleEndian.PutUint32(payloadLength, uint32(len(payload)))

			packetHeader := append(typeOfService, userId...)
			packetHeader = append(packetHeader, payloadLength...)
			buffer := append(packetHeader, payload...)
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
		log.Println(res)
	}
}
