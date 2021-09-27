package main

import "net"

type User struct {
	Conn   net.Conn
	userId uint32
	x      float32
	y      float32
	cx     float32
	cy     float32
}

type Recv struct {
	userId uint32
	buffer []byte
}
