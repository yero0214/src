package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Id   string
	Pwd  string
	Name string
	Cnt  string
}

func main() {
	// for {
	// 	fmt.Println("Enter")
	// 	fmt.Println("'login' to login")
	// 	fmt.Println("'register' to register")

	// 	text := input()
	// 	if text == "login" {
	// 		fmt.Print(login())
	// 		// if login() == 1 {

	// 		// 	break
	// 		// }
	// 		fmt.Println("user not found")
	// 	} else if text == "register" {
	// 		register()
	// 	}
	// }
	var name string
	var roomNo string
	var state string
	fmt.Println("Enter your name")
	for {
		text := input()
		if len(text) > 10 {
			fmt.Println("less than 10 characters")
		} else if len(text) < 1 {
			fmt.Println("at least one character")
		} else {
			name = text
			break
		}
	}
	for {
		fmt.Println("Type 'start' to start match making")
		for {
			text := input()
			if text == "start" {
				break
			}
		}

		conn, err := net.Dial("tcp", ":9393")
		if nil != err {
			log.Println(err)
		}
		state = "start"
		fmt.Println("Finding Opponent...")
		defer conn.Close()

		go read(conn, name, &roomNo, &state)

		for {
			s := input()
			if state == "end" {
				break
			}
			if roomNo == "" {
				continue
			}
			skills(s, conn, roomNo)
		}
	}
}

func read(conn net.Conn, name string, roomNo *string, state *string) {
	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			*roomNo = ""
			*state = "end"
			fmt.Println("Press enter to quit")
			break
		}
		res := data[:n]

		//match found
		if string(res[:2]) == "01" {
			*roomNo = string(res[2:4])
			conn.Write([]byte("11" + *roomNo + name))
			fmt.Println("Match found!")
		} else {
			fmt.Println(string(res))
		}
	}
}

func login() string {
	var id string
	var pwd string
	for {
		fmt.Print("id : ")
		id = input()
		if id != "" {
			break
		}
	}
	for {
		fmt.Print("password : ")
		pwd = input()
		if pwd != "" {
			break
		}
	}
	result := apiCall("POST", "http://localhost:9494/api/login", `{"ID":"`+id+`", "Pwd":"`+pwd+`"}`)
	fmt.Println(result)
	return result.Cnt
}

func register() {

}

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

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func apiCall(method string, url string, str string) Data {
	var data Data
	var jsonStr = []byte(str)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	data.Cnt = string(body)
	return data
}
