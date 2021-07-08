package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func login() string {
	var id string
	var pwd string
	for {
		fmt.Println("id : ")
		id = input()
		if id != "" {
			break
		}
	}
	for {
		fmt.Println("password : ")
		pwd = input()
		if pwd != "" {
			break
		}
	}
	resp, err := http.PostForm("http://localhost:9494/api/login",
		url.Values{"ID": {id}, "Pwd": {pwd}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func register() {

}
