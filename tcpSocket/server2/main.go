package main

var recvChan = make(chan []byte)

func main() {

	go controller()

	startListen("9393")
}
