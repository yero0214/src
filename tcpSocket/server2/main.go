package main

var recvChan = make(chan Recv)
var users []User
var count uint32

func main() {

	go inGame()
	go positionUpdate()
	go controller()

	startListen("9393")
}
