package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Result struct {
	ID     string
	Pwd    string
	Name   string
	Count  int
	Status string
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

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result Result
	var body Body
	_ = json.NewDecoder(r.Body).Decode(&body)

	conn().Raw("SELECT id, name FROM user_list WHERE id = ? AND pwd = ? AND delyn = 'N'", body.ID, body.Pwd).Scan(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var result Result

	conn().Raw("SELECT status FROM user_list WHERE id = ?", params["id"]).Scan(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result Result
	var body Body
	_ = json.NewDecoder(r.Body).Decode(&body)

	conn().Raw("INSERT INTO user_list (id, pwd, name, delyn, dttm) VALUES (?, ?, ?, 'N', now()) returning id", body.ID, body.Pwd, body.Name).Scan(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func idCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var result Result

	conn().Raw("SELECT count(id) FROM user_list WHERE id = ?", params["id"]).Scan(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func nmCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var result Result

	conn().Raw("SELECT count(name) FROM user_list WHERE name = ?", params["name"]).Scan(&result)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/login", login).Methods("POST")
	r.HandleFunc("/api/login/status/{id}", statusCheck).Methods("GET")
	r.HandleFunc("/api/register", register).Methods("POST")
	r.HandleFunc("/api/register/idCheck/{id}", idCheck).Methods("GET")
	r.HandleFunc("/api/register/nmCheck/{name}", nmCheck).Methods("GET")

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":9494", handler))
}

func conn() *gorm.DB {
	dsn := "host=localhost user=postgres password=365365 dbname=postgres port=5432 sslmode=disable"
	// dsn := "host=218.50.42.8 user=postgres password=365365 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db
}
