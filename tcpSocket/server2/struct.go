package main

import "net"

type User struct {
	Conn   net.Conn
	userId uint64
	x      uint64
	y      uint64
	cx     uint64
	cy     uint64
}
