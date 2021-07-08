package main

import "net"

func skills(text string, conn net.Conn, roomNo string) {
	if text == "mu" {
		mu(conn, roomNo)
	} else if text == "attack" {
		attack(conn, roomNo)
	} else if text == "heal" {
		heal(conn, roomNo)
	}
}

func attack(conn net.Conn, roomNo string) {
	conn.Write([]byte("10" + roomNo + "01"))
}

func mu(conn net.Conn, roomNo string) {
	conn.Write([]byte("10" + roomNo + "00"))
}

func heal(conn net.Conn, roomNo string) {
	conn.Write([]byte("10" + roomNo + "02"))
}
