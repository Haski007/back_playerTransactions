package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
**	Structures
*/

type user struct {
	ID           uint64  `json:"id"`
	Balance      float64 `json:"balance"`
	DepositCount int
	DepositSum   float64
	BetCount     int
	BetSum       float64
	WinCount     int
	WinSum       float64
	Token        string  `json:"token"`
}

var usersMap = map[uint64]*user{}

/*
**	Code
*/

func getUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var dat user

	json.Unmarshal(bytes, &dat)

	res, err := json.Marshal(usersMap[dat.ID])
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}
	w.Write([]byte(res))
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var u user

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.Unmarshal(bytes, &u)
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}

	if u.Token != TokenFromConfig {
		fmt.Println("invalid token")
		w.Write([]byte(`{"error" : "invalid token"}`))
		return
	}

	if _, ok := usersMap[u.ID]; ok {
		fmt.Println("user already exist")
		w.Write([]byte(`{"error": "user already exist"}`))
		return
	} else {
		fmt.Println(string(bytes))
		usersMap[u.ID] = &u
		w.Write([]byte(`{"error" : ""}`))

	}
}
