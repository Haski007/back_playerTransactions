package main

import (
	"github.com/globalsign/mgo"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	// "github.com/globalsign/mgo/bson"
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

// var usersMap = map[uint64]*user{}

var usersCollection *mgo.Collection


/*
**	Code
*/

func getUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var inputData user
	var	resData user

	err = json.Unmarshal(bytes, &inputData)
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}

	err = usersCollection.FindId(inputData.ID).One(&resData)
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}
	
	res, err := json.Marshal(resData)
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

	var	test user
	if  err = usersCollection.FindId(u.ID).One(&test); err == nil {
		fmt.Println("user already exist")
		w.Write([]byte(`{"error": "user already exist"}`))
		usersCollection.UpdateId(u.ID, u)
		return
	} else {
		fmt.Println("Created user: ", string(bytes))

		err = usersCollection.Insert(u)
		if err != nil {
			w.Write([]byte(`{"error" :` + err.Error() + `}`))
			return
		}

		w.Write([]byte(`{"error" : ""}`))
	}
}
