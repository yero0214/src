package main

import (
	"fmt"
	"net"
)

var roomNo string
var champs []Champ
var conn net.Conn
var playerNo string
var user User
var game Game
var state bool

func main() {
	res := make(chan []byte)
	finished := make(chan bool)
	champs = append(champs, Champ{No: "00", Name: "Ashe", Hp: 100, Atk: 10})
	champs = append(champs, Champ{No: "01", Name: "MG", Hp: 1000, Atk: 1})

	user.Name = login()
	conn = connect()
	write("10" + user.Name)

	defer conn.Close()

	go reactor(res, &user)
	go read(conn, res)
	for {
		fmt.Println("Type 'start' to start match making.")
		if input() == "start" {
			start(conn, user)
			<-finished
		}
	}
}
