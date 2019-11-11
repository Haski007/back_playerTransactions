package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/globalsign/mgo/bson"
)

/*
**	Structures
*/

type user struct {
	ID           uint64  `json:"id" bson:"_id,omitempty"`
	Balance      float64 `json:"balance" bson:"balance"`
	DepositCount int	`bson:"depositcount"`
	DepositSum   float64	`bson:"depositsum"`
	BetCount     int	`bson:"betcount"`
	BetSum       float64	`bson:"betsum"`
	WinCount     int	`bson:"wincount"`
	WinSum       float64 `bson:"winsum"`
	Token        string  `json:"token" bson:"token"`
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

	var data user
	var zalupa	user

	json.Unmarshal(bytes, &data)

	usersCollection.FindId(bson.ObjectId(8)).One(&zalupa)
	fmt.Println(zalupa.Balance)
	res, err := json.Marshal(usersCollection.FindId(data.ID))
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

	// usr := user{ID, Balance, DepositCount, DepositSum, BetCount, BetSum, WinCount, WinSum, Token}
	// id := r.FormValue("id")
	// token := r.FormValue("Token")
	if _, ok := usersMap[u.ID]; ok {
		fmt.Println("user already exist")
		w.Write([]byte(`{"error": "user already exist"}`))
		usersCollection.UpdateId(u.ID, u)
		return
	} else {
		fmt.Println(string(bytes))
		usersCollection.UpdateId(u.ID, u)
		usersCollection.Insert(u)
		w.Write([]byte(`{"error" : ""}`))

	}
}
