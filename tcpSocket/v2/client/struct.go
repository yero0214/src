package main

import "net"

type User struct {
	Conn net.Conn
	Name string
	No   string
}

type Champ struct {
	No   string
	Name string
	Hp   int
	Atk  int
}

type Game struct {
	Player1 Player
	Player2 Player
}

type Player struct {
	Name  string
	Champ Champ
	x     int
	y     int
}
