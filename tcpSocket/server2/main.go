package main

var recvChan = make(chan []byte)
var users []User
var count uint64

func main() {

	go inGame()
	go yUpdate()
	go xUpdate()
	go controller()

	startListen("9393")
}
