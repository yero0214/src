package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result Result
	var body Body
	var rtn string
	_ = json.NewDecoder(r.Body).Decode(&body)

	conn().Raw("SELECT count(id) FROM user_list WHERE id = ? AND pwd = ? AND delyn = 'N'", body.ID, body.Pwd).Scan(&result)
	fmt.Println(result)
	if result.Count != 0 {
		rtn = "suc"
	} else {
		rtn = "fail"
	}
	json.NewEncoder(w).Encode(rtn)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result Result
	var body Body
	var rtn string
	_ = json.NewDecoder(r.Body).Decode(&body)

	conn().Raw("INSERT INTO user_list (id, pwd, name, delyn, dttm) VALUES (?, ?, ?, 'N', now()) returning id", body.ID, body.Pwd, body.Name).Scan(&result)
	if result.ID != "" {
		rtn = "suc"
	} else {
		rtn = "fail"
	}
	json.NewEncoder(w).Encode(rtn)
}

func idCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var result Result
	var rtn string

	conn().Raw("SELECT count(id) FROM user_list WHERE id = ?", params["id"]).Scan(&result)
	fmt.Println(result)
	if result.Count == 0 {
		rtn = "suc"
	} else {
		rtn = "fail"
	}
	json.NewEncoder(w).Encode(rtn)
}
