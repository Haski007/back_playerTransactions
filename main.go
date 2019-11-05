package main

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)

/*
**	Structures
*/

type account struct {
	DepositID		uint64 `json:"depositid"`
	transactionTime	string
	OldBalance		float64
	DpositAmount	float64 `json:"amount"`
	Token			string `json:"testtask"`
	

	// user
}

type user struct {
	ID				uint64 `json:"id"`
	Balance			float64 `json:"balance"`
	DepositCount	int
	DepositSum		float64
	BetCount		int
	betSum			float64
	WinCoun			int
	WinSum			float64
	Token			string `json:"testtask"`
	account
}

/*
**	Code
*/

var usersMap = map[uint64]*user{}

func	main() {
	http.HandleFunc("/user/create", adduser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/deposit", addDeposit)

	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func addDeposit(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var dat map[string]interface{}

	json.Unmarshal(bytes, &dat)
	tmp := fmt.Sprintf("%v", dat["id"])
	id, _ := strconv.ParseUint(tmp, 10, 64)
	usr := usersMap[id]
	usr.OldBalance = usr.Balance
	da := fmt.Sprintf("%v", dat["amount"])
	fmt.Println(da)
	amount, _ := strconv.ParseFloat(fmt.Sprintf("%v", dat["amount"]), 64)
	usr.Balance += amount
	did, _ := strconv.ParseUint(fmt.Sprintf("%v", dat["depositid"]), 10, 64)
	usr.DepositID = did
}

func getUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var dat map[string]interface{}

	json.Unmarshal(bytes, &dat)
	tmp := fmt.Sprintf("%v", dat["id"])
	id, _ := strconv.ParseUint(tmp, 10, 64)
	res, _ := json.Marshal(usersMap[id])
	fmt.Println(string(res))
}

func adduser(w http.ResponseWriter, r *http.Request) {
	var u user
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if usersMap[u.ID] != nil {
		fmt.Println("user with such id already exists!")
	} else {
		usersMap[u.ID] = &u
		for _ , v := range usersMap {
			fmt.Printf( "id = %2d\nbalance = %2.2f\ntoken = %2v\n", v.ID, v.Balance, v.Token)
		}
	}
}
