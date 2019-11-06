package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TransactionType string

const (
	Bet TransactionType = "Bet"
	Win TransactionType = "Win"
)

/*
**	Structures
 */

type deposit struct {
	DepositID uint64  `json:"depositid"`
	UserID    uint64  `json:"userid"`
	Amount    float64 `json:"amount"`
	Token     string  `json:"token"`
}

type transaction struct {
	TransactionID uint64          `json:"transactionId"`
	UserID        uint64          `json:"userid"`
	Type          TransactionType `json:"type"`
	Amount        float64         `json:"amount"`
	Token         string          `json:"token"`
}

var TokenFromConfig = "testtask"

type user struct {
	ID           uint64  `json:"id"`
	Balance      float64 `json:"balance"`
	DepositCount int
	Token        string `json:"token"`
}

/*
**	Code
 */

var usersMap = map[uint64]*user{}
var usersDepositMap = map[uint64][]deposit{}
var usersTransactionsMap = map[uint64][]transaction{}

func getOldBalance(userId uint64) float64 {
	var oldBalance float64

	deposits, ok := usersDepositMap[userId]
	if !ok {
		return oldBalance
	}

	for _, val := range deposits {
		oldBalance += val.Amount
	}
	return oldBalance
}

func addDeposit(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}
	var dat deposit

	json.Unmarshal(bytes, &dat)
	usr := usersMap[dat.UserID]
	usr.DepositCount++
	usr.Balance += getOldBalance(dat.UserID) + dat.Amount
	deposits := usersDepositMap[usr.ID]
	usersDepositMap[usr.ID] = append(deposits, dat)
	w.Write([]byte(fmt.Sprintf("{\"error\" : \"\", \"balance\" : %f}", usr.Balance)))
}

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
		w.Write([]byte(`{"error" : "invalid token"}`))
		return
	}

	if _, ok := usersMap[u.ID]; ok {
		w.Write([]byte(`{"error": "user already exist"}`))
		return
	} else {
		usersMap[u.ID] = &u
		w.Write([]byte(`{"error" : ""}`))

	}
}

func makeTransaction(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var data transaction
	var balance float64

	json.Unmarshal(bytes, &data)
	switch data.Type {
	case Bet:
		balance, err = MakeBet(data)
	case Win:
		balance, err = MakeWin(data)
	default:
		err = errors.New("undefined transaction type")
	}
	if err != nil {
		log.Println(err.Error)
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"error\" : \"\", \"balance\" :%f}", balance)))
}

func MakeBet(data transaction) (float64, error) {
	if data.Amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	usr := usersMap[data.UserID]
	if usr.Balance < data.Amount {
		return 0, errors.New("no money")
	}

	usr.Balance -= data.Amount
	usersTransactionsMap[usr.ID] = append(usersTransactionsMap[usr.ID], data)
	return usr.Balance, nil
}

func MakeWin(data transaction) (float64, error) {
	usr := usersMap[data.UserID]
	usr.Balance += data.Amount
	usersTransactionsMap[usr.ID] = append(usersTransactionsMap[usr.ID], data)
	return usr.Balance, nil
}

func main() {
	http.HandleFunc("/user/create", addUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/deposit", addDeposit)
	http.HandleFunc("/transaction", makeTransaction)

	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
