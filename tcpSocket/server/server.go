package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type List struct {
	Clients []net.Conn
}

func main() {
	var list List
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		fmt.Println("connected")
		if err != nil {
			fmt.Println(err)
			return
		}
		list.Clients = append(list.Clients, conn)
		go handleClient(&list, conn)
	}
}

func handleClient(list *List, conn net.Conn) {
	defer conn.Close()
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		broadCast(list, string(netData))
	}
}

func broadCast(list *List, text string) {
	for _, v := range list.Clients {
		v.Write([]byte(text))
	}
}
