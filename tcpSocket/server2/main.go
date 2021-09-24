package main

var recvChan = make(chan []byte)
var users []User
var count uint32

func main() {

	go inGame()
	go positionUpdate()
	go controller()

	startListen("9393")
}
