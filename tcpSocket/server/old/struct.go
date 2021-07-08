package main

import "net"

type User struct {
	conn   net.Conn
	health int
	name   string
}

type Rooms struct {
	room []Room
}

type Room struct {
	users []User
}
