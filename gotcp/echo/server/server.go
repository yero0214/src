package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// creates a tcp listener
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8989")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// accept connection
	conn, err := listener.AcceptTCP()
	if err != nil {
		continue
	}

	// run loop
	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
