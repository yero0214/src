package main

type Result struct {
	ID    string
	Pwd   string
	Name  string
	Count int
}

type Body struct {
	ID   string
	Pwd  string
	Name string
}

type message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
