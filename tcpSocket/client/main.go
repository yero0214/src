package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Id     string `json:"ID"`
	Pwd    string `json:"Pwd"`
	Name   string `json:"Name"`
	Status string `json:"Status"`
	Count  int    `json:"Count"`
}

func main() {
	var name string
	var roomNo string
	var state string

	conn, err := net.Dial("tcp", "218.50.42.8:9393")
	if nil != err {
		log.Println(err)
	}

	for {
		fmt.Println("'login' to login")
		fmt.Println("'register' to register")

		text := input()
		if text == "login" {
			data := login()
			if data.Name != "" {
				if checkStatus(data.Id).Status == "offline" {
					conn.Write([]byte("00" + data.Id))

					name = data.Name
					fmt.Println("login success")
					break
				} else {
					fmt.Println("your account is currently logged onto another device")
				}
			}
			fmt.Println("user not found")
		} else if text == "register" {
			if register().Id != "" {
				fmt.Println("register success")
			} else {
				fmt.Println("register fail")
			}
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

func login() Data {
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
	return apiCall("POST", "http://localhost:9494/api/login", `{"ID":"`+id+`", "Pwd":"`+pwd+`"}`)
}

func checkStatus(id string) Data {
	return apiCall("GET", "http://localhost:9494/api/login/status/"+id, `{}`)
}

func register() Data {
	var id string
	var pwd string
	var name string

	for {
		fmt.Print("id : ")
		id = input()
		if id != "" {
			result := apiCall("GET", "http://localhost:9494/api/register/idCheck/"+id, `{}`)
			if result.Count > 0 {
				fmt.Println("id already exist")
			} else {
				break
			}
		}
	}
	for {
		fmt.Print("password : ")
		pwd = input()
		if pwd != "" {
			break
		}
	}
	for {
		fmt.Print("name : ")
		name = input()
		if len(name) > 10 {
			fmt.Println("less than 10 characters")
		} else if len(name) < 1 {
			fmt.Println("at least one character")
		} else {
			result := apiCall("GET", "http://localhost:9494/api/register/nmCheck/"+name, `{}`)
			if result.Count > 0 {
				fmt.Println("name already exist")
			} else {
				break
			}
		}
	}
	return apiCall("POST", "http://localhost:9494/api/login", `{"ID":"`+id+`", "Pwd":"`+pwd+`"}`)
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
	json.Unmarshal(body, &data)
	return data
}
