package main

import "net"

type User struct {
	Conn   net.Conn
	userId uint32
	x      uint32
	y      uint32
	cx     uint32
	cy     uint32
}
